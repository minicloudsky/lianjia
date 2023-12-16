package data

import (
	"context"
	"errors"
	"github.com/minicloudsky/lianjia/internal/biz"
	"gorm.io/gorm"
	"strings"
	"time"
)

type commercialPo struct {
	BaseModel
	CityID                  int       `gorm:"column:city_id;type:int;default:0;not null;comment:城市ID" json:"city_id"`
	HouseCode               int       `gorm:"column:house_code;type:bigint;default:0;not null;comment:房屋编码" json:"house_code"`
	BuildingID              int       `gorm:"column:building_id;type:bigint;default:0;not null;comment:楼栋ID" json:"building_id"`
	HousedelCode            int64     `gorm:"column:housedel_code;type:bigint;default:0;not null;comment:户型编码" json:"housedel_code"`
	Title                   string    `gorm:"column:title;type:varchar(300);default:'';not null;comment:标题" json:"title"`
	CityName                string    `gorm:"column:city_name;type:varchar(50);default:'';not null;comment:城市名称" json:"city_name"`
	DistrictName            string    `gorm:"column:district_name;type:varchar(80);default:'';not null;comment:区域名称" json:"district_name"`
	BizcircleName           string    `gorm:"column:bizcircle_name;type:varchar(80);default:'';not null;comment:商圈名称" json:"bizcircle_name"`
	StreetName              string    `gorm:"column:street_name;type:varchar(150);default:'';not null;comment:街道名称" json:"street_name"`
	ResblockName            string    `gorm:"column:resblock_name;type:varchar(150);default:'';not null;comment:小区名称" json:"resblock_name"`
	UnitRentPrice           float64   `gorm:"column:unit_rent_price;type:decimal(10,2);default:0.00;not null;comment:单元租金" json:"unit_rent_price"`
	UnitMonthRentPrice      float64   `gorm:"column:unit_month_rent_price;type:decimal(10,2);default:0.00;not null;comment:单元月租金" json:"unit_month_rent_price"`
	RentPrice               int       `gorm:"column:rent_price;type:int;default:0;not null;comment:租金" json:"rent_price"`
	UnitSellPrice           int       `gorm:"column:unit_sell_price;type:int;default:0;not null;comment:单元售价" json:"unit_sell_price"`
	SellPrice               int       `gorm:"column:sell_price;type:int;default:0;not null;comment:售价" json:"sell_price"`
	Area                    float64   `gorm:"column:area;type:decimal(10,2);default:0.00;not null;comment:面积" json:"area"`
	Score                   float64   `gorm:"column:score;type:decimal(5,2);default:0.00;not null;comment:评分" json:"score"`
	CmarkName               string    `gorm:"column:cmark_name;type:varchar(255);default:'';not null;comment:标记名称" json:"cmark_name"`
	Image                   string    `gorm:"column:image;type:varchar(255);default:'';not null;comment:图片" json:"image"`
	Fitment                 int       `gorm:"column:fitment;type:int;default:0;not null;comment:装修" json:"fitment"`
	FitmentName             string    `gorm:"column:fitment_name;type:varchar(255);default:'';not null;comment:装修名称" json:"fitment_name"`
	HasFurniture            int       `gorm:"column:has_furniture;type:int(3);default:0;not null;comment:是否有家具" json:"has_furniture"`
	HasFurnitureName        string    `gorm:"column:has_furniture_name;type:varchar(255);default:'';not null;comment:有无家具名称" json:"has_furniture_name"`
	HasMeetingRoom          int       `gorm:"column:has_meeting_room;type:int(3);default:0;not null;comment:是否有会议室" json:"has_meeting_room"`
	HasMeetingRoomName      string    `gorm:"column:has_meeting_room_name;type:varchar(255);default:'';not null;comment:有无会议室名称" json:"has_meeting_room_name"`
	ParkingType             int       `gorm:"column:parking_type;type:int;default:0;not null;comment:停车类型" json:"parking_type"`
	IsNearSubway            int       `gorm:"column:is_near_subway;type:int(2);default:0;not null;comment:是否靠近地铁" json:"is_near_subway"`
	IsNearSubwayName        string    `gorm:"column:is_near_subway_name;type:varchar(255);default:'';not null;comment:是否靠近地铁名称" json:"is_near_subway_name"`
	IsRegisteredCompany     int       `gorm:"column:is_registered_company;type:int(2);default:0;not null;comment:是否注册公司" json:"is_registered_company"`
	IsRegisteredCompanyName string    `gorm:"column:is_registered_company_name;type:varchar(255);default:'';not null;comment:是否注册公司名称" json:"is_registered_company_name"`
	FloorPosition           int       `gorm:"column:floor_position;type:int(2);default:0;not null;comment:楼层位置" json:"floor_position"`
	FloorPositionName       string    `gorm:"column:floor_position_name;type:varchar(255);default:'';not null;comment:楼层位置名称" json:"floor_position_name"`
	ShowingTime             int       `gorm:"column:showing_time;type:int;default:0;not null;comment:展示时间" json:"showing_time"`
	ShowingTimeName         string    `gorm:"column:showing_time_name;type:varchar(255);default:'';not null;comment:展示时间名称" json:"showing_time_name"`
	SubwayDistance          int       `gorm:"column:subway_distance;type:int(10);default:0;not null;comment:距离地铁距离" json:"subway_distance"`
	SubwayDistanceName      string    `gorm:"column:subway_distance_name;type:varchar(255);default:'';not null;comment:距离地铁距离名称" json:"subway_distance_name"`
	Tags                    string    `gorm:"column:tags;type:varchar(255);default:'';not null;comment:标签" json:"tags"`
	IsReal                  int       `gorm:"column:is_real;type:int(2);default:0;not null;comment:是否真实" json:"is_real"`
	HasVr                   int       `gorm:"column:has_vr;type:int(2);default:0;not null;comment:是否有VR" json:"has_vr"`
	Ctime                   time.Time `gorm:"column:ctime;type:datetime;comment:创建时间" json:"ctime"`
}

func (c *commercialPo) TableName() string {
	return "t_commercial"
}

func (c *commercialPo) Comment() string {
	return "链家商业办公"
}

func (repo lianjiaRepo) InsertCommercialInfo(ctx context.Context, infos []*biz.CommercialInfo) error {
	var pos []*commercialPo
	for _, info := range infos {
		po := &commercialPo{
			HouseCode:               info.HouseCode,
			BuildingID:              info.BuildingID,
			HousedelCode:            info.HousedelCode,
			Title:                   info.Title,
			CityID:                  info.CityID,
			CityName:                info.CityName,
			DistrictName:            info.DistrictName,
			BizcircleName:           info.BizcircleName,
			StreetName:              info.StreetName,
			ResblockName:            info.ResblockName,
			UnitRentPrice:           info.UnitRentPrice,
			UnitMonthRentPrice:      info.UnitMonthRentPrice,
			RentPrice:               info.RentPrice,
			UnitSellPrice:           info.UnitSellPrice,
			SellPrice:               info.SellPrice,
			Area:                    info.Area,
			Score:                   info.Score,
			CmarkName:               info.CmarkName,
			Image:                   info.Image,
			Fitment:                 info.Fitment,
			FitmentName:             info.FitmentName,
			HasFurniture:            info.HasFurniture,
			HasFurnitureName:        info.HasFurnitureName,
			HasMeetingRoom:          info.HasMeetingRoom,
			HasMeetingRoomName:      info.HasMeetingRoomName,
			ParkingType:             info.ParkingType,
			IsNearSubway:            info.IsNearSubway,
			IsNearSubwayName:        info.IsNearSubwayName,
			IsRegisteredCompany:     info.IsRegisteredCompany,
			IsRegisteredCompanyName: info.IsRegisteredCompanyName,
			FloorPosition:           info.FloorPosition,
			FloorPositionName:       info.FloorPositionName,
			ShowingTime:             info.ShowingTime,
			ShowingTimeName:         info.ShowingTimeName,
			SubwayDistance:          info.SubwayDistance,
			SubwayDistanceName:      info.SubwayDistanceName,
			Tags:                    info.Tags,
			IsReal:                  info.IsReal,
			HasVr:                   info.HasVr,
			Ctime:                   info.Ctime,
		}
		pos = append(pos, po)
	}
	tx := repo.data.DB(ctx).Model(&commercialPo{}).CreateInBatches(pos, batchInsertSize)
	if tx.Error != nil && !strings.Contains(tx.Error.Error(), "Duplicate entry") {
		return tx.Error
	}
	return nil
}

func (repo lianjiaRepo) CheckCommercialExists(ctx context.Context, cityID, houseCode int64) (bool, error) {
	var commercial commercialPo
	result := repo.data.DB(ctx).Where("city_id = ? AND house_code = ?", cityID, houseCode).First(&commercial)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	if commercial.ID > 0 {
		return true, nil
	}
	return false, nil
}
