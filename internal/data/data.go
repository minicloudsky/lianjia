package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/minicloudsky/lianjia/internal/biz"
	"github.com/minicloudsky/lianjia/internal/conf"
	"github.com/redis/go-redis/v9"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"
	rawLog "log"
	"os"
	"strings"
	// init mysql driver
	_ "github.com/go-sql-driver/mysql"
	"gorm.io/driver/mysql"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewLianjiaRepo, NewRedis, NewTransaction,
	NewDB, NewKafkaManager, NewChannelQueue, NewKafkaQueue, NewQueue)

// Data .
type Data struct {
	redisClient *redis.Client
	db          *gorm.DB
	confData    *conf.Data
	km          *biz.KafkaManager
	queue       Queue
}

type BaseModel struct {
	ID         int64      `gorm:"primary_key;type:bigint;" json:"id"`
	CreateTime *time.Time `gorm:"column:create_time;type:datetime not null;default:CURRENT_TIMESTAMP();comment:创建时间" json:"created_at"`
	UpdateTime *time.Time `gorm:"column:update_time;type:datetime on update current_timestamp;comment:更新时间" json:"updated_at"`
}

type contextTransactionKey struct{}

func NewKafkaSender(addr, topic string) (*biz.KafkaSender, error) {
	var brokers []string
	if strings.Contains(addr, ",") {
		brokers = strings.Split(addr, ",")
	} else {
		brokers = append(brokers, addr)
	}
	kafkaAddr := kafka.TCP(brokers...)
	writer := kafka.Writer{
		Addr:         kafkaAddr,
		Topic:        topic,
		MaxAttempts:  3,
		BatchTimeout: 10 * time.Second,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		RequiredAcks: kafka.RequireOne,
	}
	return &biz.KafkaSender{Writer: &writer}, nil
}

func NewKafkaReader(addr, topic, groupId string) *biz.KafkaReader {
	var brokers []string
	if strings.Contains(addr, ",") {
		brokers = strings.Split(addr, ",")
	} else {
		brokers = append(brokers, addr)
	}
	// Kafka Reader conf
	config := kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MaxBytes: 10e6, // 10MB
		GroupID:  groupId,
	}
	// create Kafka Reader
	reader := kafka.NewReader(config)

	return &biz.KafkaReader{Reader: reader}
}

func NewKafkaManager(d *conf.Data, logger log.Logger) (*biz.KafkaManager, error) {
	brokers := make(map[string][]string)
	senders := make(map[string]*biz.KafkaSender)
	readers := make(map[string]*biz.KafkaReader)
	for _, topic := range d.Kafka.Topics {
		topicName := topic.Name
		var brokers []string
		if strings.Contains(d.Kafka.Addr, ",") {
			brokers = strings.Split(d.Kafka.Addr, ",")
		} else {
			brokers = append(brokers, d.Kafka.Addr)
		}
		sender, err := NewKafkaSender(d.Kafka.Addr, topicName)
		if err != nil {
			return nil, err
		}
		senders[topicName] = sender
		reader := NewKafkaReader(d.Kafka.Addr, topicName, topicName)
		readers[topicName] = reader
	}

	return &biz.KafkaManager{
		Brokers: brokers,
		Senders: senders,
		Readers: readers,
		Logger:  logger,
	}, nil
}

func (d *Data) Tx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, contextTransactionKey{}, tx)
		return fn(ctx)
	})
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	tx, ok := ctx.Value(contextTransactionKey{}).(*gorm.DB)
	if ok {
		return tx
	}
	return d.db.WithContext(ctx)
}

func NewTransaction(d *Data) biz.Transaction {
	return d
}

func NewDB(conf *conf.Data, logger log.Logger) *gorm.DB {
	dbLog := log.NewHelper(log.With(logger, "module", "data/gorm"))
	gormLogBaseDir := "/data/lianjia/logs"
	file, err := os.Create(fmt.Sprintf("%s/database.log", gormLogBaseDir))
	if err != nil {
		panic(err)
	}
	config := gormLogger.Config{
		SlowThreshold: time.Second,
		LogLevel:      gormLogger.Warn,
		Colorful:      true,
	}
	fileLogger := gormLogger.New(rawLog.New(file, "", rawLog.LstdFlags), config)
	db, err := gorm.Open(mysql.Open(conf.Database.Source), &gorm.Config{
		Logger: fileLogger,
		NowFunc: func() time.Time {
			return time.Now().Local()
		},
		PrepareStmt: true,
	})
	if err != nil {
		dbLog.Fatalf("failed opening connection to mysql: %v", err)
	}

	if !db.Migrator().HasTable(&cityPo{}) {
		if err := db.AutoMigrate(&cityPo{}); err != nil {
			dbLog.Fatal(err)
		}
	}
	if !db.Migrator().HasTable(&erShouFangInfoPo{}) {
		if err := db.AutoMigrate(&erShouFangInfoPo{}); err != nil {
			dbLog.Fatal(err)
		}
	}
	if !db.Migrator().HasTable(&loupanPo{}) {
		if err := db.AutoMigrate(&loupanPo{}); err != nil {
			dbLog.Fatal(err)
		}
	}

	if !db.Migrator().HasTable(&commercialPo{}) {
		if err := db.AutoMigrate(&commercialPo{}); err != nil {
			dbLog.Fatal(err)
		}
	}

	return db
}

// NewData .
func NewData(c *conf.Data, logger log.Logger, db *gorm.DB, km *biz.KafkaManager, q Queue) (*Data, func(), error) {
	dataLog := log.NewHelper(logger)
	redisClient, err := NewRedis(c, logger)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
		err = redisClient.Close()
		if err != nil {
			dataLog.Errorf("close redis conn err! %v", err)
			return
		}
		err = km.Close()
		if err != nil {
			dataLog.Errorf("close kafka writer err! %v", err)
			return
		}
	}
	return &Data{
		redisClient: redisClient,
		db:          db,
		confData:    c,
		km:          km,
		queue:       q,
	}, cleanup, nil
}

func NewRedis(c *conf.Data, logger log.Logger) (redisConn *redis.Client, err error) {
	redisLogger := log.NewHelper(logger)
	redisOptions := &redis.Options{
		Network:      c.Redis.Network,
		Addr:         c.Redis.Addr,
		Password:     c.Redis.Password,
		DB:           int(c.Redis.Db),
		MaxRetries:   3,
		ReadTimeout:  c.Redis.ReadTimeout.AsDuration(),
		WriteTimeout: c.Redis.WriteTimeout.AsDuration(),
	}
	redisClient := redis.NewClient(redisOptions)
	if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
		redisLogger.Errorf("fail to connect redis ! err: %v", err)
		panic(err)
		return nil, err
	}
	ctx := context.Background()

	// clean lianjia existing task
	for _, key := range biz.LianjiaTaskKeys {
		_, err = redisClient.Del(ctx, key).Result()
		if err != nil {
			redisLogger.Errorf("clean lianjia existing task err ! %v", err)
			return nil, err
		}
	}
	return redisClient, nil
}

func NewQueue(c *conf.Data, logger log.Logger, kafkaQueue *KafkaQueue, chanQueue *ChanQueue) (q Queue, err error) {
	queueHelper := log.NewHelper(logger)
	queueHelper.Infof("current queue mode: %s", c.Queue.Mode)
	switch c.Queue.Mode {
	case biz.QueueModeKafka:
		return kafkaQueue, nil
	case biz.QueueModeGoChannel:
		return chanQueue, nil
	default:
		return nil, errors.New(fmt.Sprintf("unknown mode %s", c.Queue.Mode))
	}
}

func NewKafkaQueue(km *biz.KafkaManager, logger log.Logger) *KafkaQueue {
	return &KafkaQueue{
		km:     km,
		logger: logger,
	}
}

func NewChannelQueue(logger log.Logger) *ChanQueue {
	ershoufangChan := make(chan []biz.Message, 100)
	loupanChan := make(chan []biz.Message, 100)
	commercialChan := make(chan []biz.Message, 100)
	zufangChan := make(chan []biz.Message, 100)
	return &ChanQueue{
		ErshoufangChan: ershoufangChan,
		LoupanChan:     loupanChan,
		CommercialChan: commercialChan,
		ZufangChan:     zufangChan,
		logger:         logger,
	}
}
