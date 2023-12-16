package biz

import (
	"context"
	"crypto/md5"
	"encoding/hex"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minicloudsky/lianjia/pkg/pagination"
	"strconv"
	"time"
)

type CityInfo struct {
	Id            int    `json:"id"`
	CityCode      int    `json:"city_code"`
	Name          string `json:"name"`
	Host          string `json:"host"`
	MobileHost    string `json:"mobile_host"`
	Short         string `json:"short"`
	Province      string `json:"province"`
	ProvinceShort string `json:"province_short"`
	Pinyin        string `json:"pinyin"`
	Longitude     string `json:"longitude"`
	Latitude      string `json:"latitude"`
	Url           string `json:"url"`
}

type ZufangInfo struct {
}

type CityMap map[string]CityInfo

type HoueseType string

const (
	HouseTypeErshoufang                  HoueseType = "ershoufang"
	HouseTypeLoupan                      HoueseType = "loupan"
	HouseTypeZufang                      HoueseType = "zufang"
	HouseTypeCommercial                  HoueseType = "commercial"
	LianjiaErshoufangRunningCrawlTaskKey            = "lianjia_ershoufang_running_task"
	LianjiaLoupanRunningCrawlTaskKey                = "lianjia_loupan_running_task"
	LianjiaCommercialRunningCrawlTaskKey            = "lianjia_commercial_running_task"
	LianjiaErshoufangIdempotentKey                  = "ershoufang_"
	LianjiaLoupanIdempotentKey                      = "loupan_"
	LianjiaCommercialIdempotentKey                  = "commercial_"
)

var LianjiaTaskKeys = []string{
	LianjiaErshoufangRunningCrawlTaskKey,
	LianjiaLoupanRunningCrawlTaskKey,
	LianjiaCommercialRunningCrawlTaskKey,
}

const insertBatchSize = 1000

func (h HoueseType) Topic() string {
	switch h {
	case HouseTypeErshoufang:
		return string(TopicErshoufang)
	case HouseTypeLoupan:
		return string(TopicLoupan)
	case HouseTypeCommercial:
		return string(TopicCommercial)
	case HouseTypeZufang:
		return string(TopicZufang)
	}
	return ""
}

type Topic string

const (
	TopicErshoufang Topic = "lianjiaershoufang"
	TopicLoupan     Topic = "lianjialoupan"
	TopicZufang     Topic = "lianjiazufang"
	TopicCommercial Topic = "lianjiacommercial"
)

var (
	ErshoufangProcessChan chan []Message
	LoupanProcessChan     chan []Message
	CommercialProcessChan chan []Message
	ZufangProcessChan     chan []Message
)

type Message struct {
	Content []byte
}

const (
	QueueModeKafka     = "kafka"
	QueueModeGoChannel = "channel"
	ChanQueueSize      = 100
)

type LianjiaRepo interface {
	UpsertCity(ctx context.Context, cities []*CityInfo) error
	GetAllCity(ctx context.Context) (cities []*CityInfo, err error)
	InsertErshoufangInfo(ctx context.Context, info []*ErShouFangInfo) error
	InsertLoupanInfo(ctx context.Context, lists []*LoupanInfo) error
	GetKafkaManager(ctx context.Context) (km *KafkaManager, err error)
	SetNXKey(ctx context.Context, key string, value interface{}, duration time.Duration) error
	IncrKey(ctx context.Context, key string, increment int) (int64, error)
	DelKey(ctx context.Context, key string) error
	GetKey(ctx context.Context, key string) (string, error)
	Send(ctx context.Context, msg []Message, houseType HoueseType) error
	Receive(ctx context.Context) error
	ListErshoufang(ctx context.Context, p pagination.Pagination, query string) ([]*ErShouFangInfo, int64, error)
	InsertCommercialInfo(ctx context.Context, infos []*CommercialInfo) error
	CheckErShouFangExists(ctx context.Context, cityID int64, houseCode string) (bool, error)
	CheckLoupanExists(ctx context.Context, cityID, loupanId int64) (bool, error)
	CheckCommercialExists(ctx context.Context, cityID, houseCode int64) (bool, error)
}

type Transaction interface {
	Tx(context.Context, func(ctx context.Context) error) error
}

type LianjiaUsecase struct {
	repo LianjiaRepo
	log  *log.Helper
}

func NewLianjiaUsecase(repo LianjiaRepo, logger log.Logger) *LianjiaUsecase {
	return &LianjiaUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *LianjiaUsecase) InitLianjiaChan(ctx context.Context) {
	uc.log.Infof("---initializing lianjia channel...")
	ErshoufangProcessChan = make(chan []Message, ChanQueueSize)
	ZufangProcessChan = make(chan []Message, ChanQueueSize)
	LoupanProcessChan = make(chan []Message, ChanQueueSize)
	CommercialProcessChan = make(chan []Message, ChanQueueSize)
}

func (uc *LianjiaUsecase) SetNXKey(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	return uc.repo.SetNXKey(ctx, key, value, duration)
}

func (uc *LianjiaUsecase) Send(ctx context.Context, msg []Message, houseType HoueseType) error {
	return uc.repo.Send(ctx, msg, houseType)
}

func (uc *LianjiaUsecase) Receive(ctx context.Context) error {
	err := uc.repo.Receive(ctx)
	if err != nil {
		uc.log.Errorf("Receive message err! %v", err)
	}
	return nil
}

func (uc *LianjiaUsecase) md5(s string) string {
	hash := md5.New()
	hash.Write([]byte(s))
	return hex.EncodeToString(hash.Sum(nil))
}

func (uc *LianjiaUsecase) finishCurrentTask(ctx context.Context, houseType HoueseType, city *CityInfo) error {
	var taskKey string
	var currentChan chan []Message
	switch houseType {
	case HouseTypeErshoufang:
		taskKey = LianjiaErshoufangRunningCrawlTaskKey
		currentChan = ErshoufangProcessChan
	case HouseTypeLoupan:
		taskKey = LianjiaLoupanRunningCrawlTaskKey
		currentChan = ZufangProcessChan
	case HouseTypeCommercial:
		taskKey = LianjiaCommercialRunningCrawlTaskKey
		currentChan = CommercialProcessChan
	}
	unFinishTaskStr, err := uc.repo.GetKey(ctx, taskKey)
	if err != nil {
		uc.log.Errorf("get %s unfinished task key err! %v", string(houseType), err)
		return err
	}
	unFinishTasks, err := strconv.ParseInt(unFinishTaskStr, 10, 64)
	if err != nil {
		uc.log.Errorf("parse %s task num err! %v", string(houseType), err)
		return err
	}
	if unFinishTasks > 0 {
		_, err := uc.repo.IncrKey(ctx, taskKey, -1)
		if err != nil {
			uc.log.Errorf("decrease %s left task err! %v", string(houseType), err)
			return err
		}
		unFinishTaskStr, err := uc.repo.GetKey(ctx, taskKey)
		if err != nil {
			uc.log.Errorf("get %s unfinished task key err! %v", string(houseType), err)
			return err
		}
		unFinishTasks, err := strconv.ParseInt(unFinishTaskStr, 10, 64)
		if err != nil {
			uc.log.Errorf("parse %s task num err! %v", string(houseType), err)
			return err
		}
		if unFinishTasks == 0 {
			uc.log.Infof("crawl %s task finished, closing channel...", string(houseType))
			close(currentChan)
			uc.log.Infof("crawl %s task finished, deleting redis key...", string(houseType))
			err := uc.repo.DelKey(ctx, taskKey)
			if err != nil {
				uc.log.Errorf("delete %s running task key err! %v", string(houseType), err)
				return err
			}
		}
		uc.log.Infof("ListCity%s crawl city %s finished task left: %d", string(houseType), city.Name, unFinishTasks)
	}

	return nil
}
