package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/minicloudsky/lianjia/internal/biz"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"strconv"
	"time"
)

type lianjiaRepo struct {
	data *Data
	log  *log.Helper
}

type cityPo struct {
	BaseModel
	CityCode      int     `gorm:"column:city_code;uniqueIndex:uniq_city_code;type:int not null default 0;comment:城市代码" json:"city_code"`
	Name          string  `gorm:"column:name;type:varchar(50) not null default '';comment:城市" json:"name"`
	Host          string  `gorm:"column:host;type:varchar(50) not null default '';comment:地址" json:"host"`
	MobileHost    string  `gorm:"column:mobile_host;type:varchar(50) not null default '';comment:移动端地址" json:"mobile_host"`
	Short         string  `gorm:"column:short;type:varchar(50) not null default '';comment:短名称" json:"short"`
	Province      string  `gorm:"column:province;type:varchar(50) not null default '';comment:省份" json:"province"`
	ProvinceShort string  `gorm:"column:province_host;type:varchar(30) not null default '';comment:省份短名称" json:"province_short"`
	Pinyin        string  `gorm:"column:pinyin;type:varchar(50) not null default '';comment:拼音" json:"pinyin"`
	Longitude     float64 `gorm:"column:longitude;type:decimal(20, 6) not null default 0;comment:经度" json:"longitude"`
	Latitude      float64 `gorm:"column:latitude;type:decimal(20, 6) not null default 0;comment:纬度" json:"latitude"`
	Url           string  `gorm:"column:url;type:varchar(100) not null default '';comment:URL" json:"url"`
}

func (c *cityPo) TableName() string {
	return "t_city"
}

func (c *cityPo) Comment() string {
	return "链家城市"
}

func (repo lianjiaRepo) UpsertCity(ctx context.Context, cities []*biz.CityInfo) error {
	for _, city := range cities {
		longitude, err := strconv.ParseFloat(city.Longitude, 64)
		if err != nil {
			return err
		}
		latitude, err := strconv.ParseFloat(city.Latitude, 64)
		if err != nil {
			return err
		}
		po := &cityPo{
			CityCode:      city.CityCode,
			Name:          city.Name,
			Host:          city.Host,
			MobileHost:    city.MobileHost,
			Short:         city.Short,
			Province:      city.Province,
			ProvinceShort: city.ProvinceShort,
			Pinyin:        city.Pinyin,
			Longitude:     longitude,
			Latitude:      latitude,
			Url:           city.Url,
		}
		tx := repo.data.DB(ctx).Model(&cityPo{}).Clauses(clause.OnConflict{
			Columns: []clause.Column{{Name: "city_code"}},
			DoUpdates: clause.AssignmentColumns([]string{"name", "host", "mobile_host", "short",
				"province", "province_host", "pinyin", "longitude", "latitude", "url",
			}),
		}).Create(&po)
		if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrDuplicatedKey) {
			return tx.Error
		}
	}
	return nil
}

func (repo lianjiaRepo) GetAllCity(ctx context.Context) (cities []*biz.CityInfo, err error) {
	var pos []*cityPo
	tx := repo.data.DB(ctx).Model(&cityPo{}).Select([]string{"city_code", "name", "url"}).
		Where("url like ?", "http%").Find(&pos)
	if tx.Error != nil {
		repo.log.Errorf("GetAllCity err! err: %v", err)
		return nil, tx.Error
	}
	for _, po := range pos {
		cities = append(cities, &biz.CityInfo{
			Id:            int(po.ID),
			CityCode:      po.CityCode,
			Name:          po.Name,
			Host:          po.Host,
			MobileHost:    po.MobileHost,
			Short:         po.Short,
			Province:      po.Province,
			ProvinceShort: po.ProvinceShort,
			Pinyin:        po.Pinyin,
			Longitude:     fmt.Sprintf("%f", po.Longitude),
			Latitude:      fmt.Sprintf("%f", po.Latitude),
			Url:           po.Url,
		})
	}

	return cities, nil
}

func (repo lianjiaRepo) SetNXKey(ctx context.Context, key string, value interface{}, duration time.Duration) error {
	res, err := repo.data.redisClient.SetNX(ctx, key, value, duration).Result()
	if err != nil {
		return err
	}
	if !res {
		return errors.New("key already exists")
	}
	return nil
}

func (repo lianjiaRepo) IncrKey(ctx context.Context, key string, increment int) (int64, error) {
	res, err := repo.data.redisClient.IncrBy(ctx, key, int64(increment)).Result()
	if err != nil {
		return -1, err
	}
	return res, nil
}

func (repo lianjiaRepo) DelKey(ctx context.Context, key string) error {
	_, err := repo.data.redisClient.Del(ctx, key).Result()
	if err != nil {
		return err
	}
	return nil
}

func (repo lianjiaRepo) GetKey(ctx context.Context, key string) (string, error) {
	res, err := repo.data.redisClient.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return res, nil
}

// NewLianjiaRepo .
func NewLianjiaRepo(data *Data, logger log.Logger) biz.LianjiaRepo {
	return &lianjiaRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}
