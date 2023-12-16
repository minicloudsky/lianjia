package data

import (
	"context"
	"errors"
	"github.com/minicloudsky/lianjia/internal/biz"
	"gorm.io/gorm"
	"strings"
	"time"
)

type loupanPo struct {
	BaseModel
	Pid                      int32     `json:"pid" gorm:"type:int(20);not null;default:0;comment:项目ID"`
	CityID                   int       `json:"city_id" gorm:"type:int(20);not null;default:0;comment:城市ID"`
	LoupanId                 int64     `json:"loupan_id" gorm:"type:bigint(20);not null;default:0;comment:楼盘ID"`
	CityName                 string    `json:"city_name" gorm:"type:varchar(30);not null;comment:城市名称"`
	DistrictName             string    `json:"district_name" gorm:"type:varchar(50);not null;comment:区域名称"`
	District                 string    `json:"district" gorm:"type:varchar(50);not null;comment:区域"`
	DistrictID               int32     `json:"district_id" gorm:"type:int(20);not null;default:0;comment:区域ID"`
	BizcircleID              int32     `json:"bizcircle_id" gorm:"type:int(20);not null;default:0;comment:商圈ID"`
	BizcircleName            string    `json:"bizcircle_name" gorm:"type:varchar(150);not null;comment:商圈名称"`
	CoverPic                 string    `json:"cover_pic" gorm:"type:varchar(300);not null;comment:封面图片"`
	MinFrameArea             int32     `json:"min_frame_area" gorm:"type:int(20);not null;default:0;comment:最小建筑面积"`
	MaxFrameArea             int32     `json:"max_frame_area" gorm:"type:int(20);not null;default:0;comment:最大建筑面积"`
	BuildID                  int64     `json:"build_id" gorm:"type:bigint(20);not null;default:0;comment:楼盘ID"`
	PermitAllReady           int       `json:"permit_all_ready" gorm:"type:int(2);not null;default:0;comment:许可证是否齐全"`
	ProcessStatus            int       `json:"process_status" gorm:"type:int(5);not null;default:0;comment:流程状态"`
	ResblockFrameArea        string    `json:"resblock_frame_area" gorm:"type:varchar(100);not null;comment:楼盘建筑面积"`
	ResblockFrameAreaRange   string    `json:"resblock_frame_area_range" gorm:"type:varchar(100);not null;comment:楼盘建筑面积范围"`
	ResblockFrameAreaDesc    string    `json:"resblock_frame_area_desc" gorm:"type:varchar(50);not null;comment:楼盘建筑面积描述"`
	Decoration               string    `json:"decoration" gorm:"type:varchar(50);not null;comment:装修情况"`
	Longitude                float64   `json:"longitude" gorm:"type:decimal(10,6);not null;default:0;comment:经度"`
	Latitude                 float64   `json:"latitude" gorm:"type:decimal(10,6);not null;default:0;comment:纬度"`
	FrameRoomsDesc           string    `json:"frame_rooms_desc" gorm:"type:varchar(100);not null;comment:户型描述"`
	Title                    string    `json:"title" gorm:"type:varchar(150);not null;comment:标题"`
	ResblockName             string    `json:"resblock_name" gorm:"type:varchar(150);not null;comment:楼盘名称"`
	ResblockAlias            string    `json:"resblock_alias" gorm:"type:varchar(200);not null;comment:楼盘别名"`
	Address                  string    `json:"address" gorm:"type:varchar(200);not null;comment:地址"`
	StoreAddr                string    `json:"store_addr" gorm:"type:varchar(300);not null;comment:门店地址"`
	AvgUnitPrice             float64   `json:"avg_unit_price" gorm:"type:decimal(10,2);not null;default:0;comment:均价"`
	AveragePrice             float64   `json:"average_price" gorm:"type:decimal(10,2);not null;default:0;comment:平均价"`
	AddressRemark            string    `json:"address_remark" gorm:"type:varchar(300);not null;comment:地址备注"`
	ProjectName              string    `json:"project_name" gorm:"type:varchar(100);not null;comment:项目名称"`
	SpecialTags              string    `json:"special_tags" gorm:"type:varchar(300);not null;comment:特色标签"`
	LianjiaSpecial           string    `json:"lianjia_special" gorm:"type:varchar(300);not null;comment:链家特色"`
	LianjiaSpecialComm       string    `json:"lianjia_special_comm" gorm:"type:varchar(300);not null;comment:链家特色说明"`
	DeveloperSpecial         string    `json:"developer_special" gorm:"type:varchar(300);not null;comment:开发商特色"`
	DeveloperSpecialType     string    `json:"developer_special_type" gorm:"type:varchar(300);not null;comment:开发商特色类型"`
	DeveloperSpecialComm     string    `json:"developer_special_comm" gorm:"type:varchar(300);not null;comment:开发商特色说明"`
	FrameRooms               string    `json:"frame_rooms" gorm:"type:varchar(300);not null;comment:户型"`
	ConvergedRooms           string    `json:"converged_rooms" gorm:"type:varchar(600);not null;comment:合并户型"`
	Tags                     string    `json:"tags" gorm:"type:varchar(200);not null;comment:标签"`
	ProjectTags              string    `json:"project_tags" gorm:"type:varchar(200);not null;comment:项目标签"`
	HouseType                string    `json:"house_type" gorm:"type:varchar(30);not null;comment:房屋类型"`
	HouseTypeValue           string    `json:"house_type_value" gorm:"type:varchar(30);not null;comment:房屋类型值"`
	SaleStatus               string    `json:"sale_status" gorm:"type:varchar(30);not null;comment:销售状态"`
	HasEvaluate              int       `json:"has_evaluate" gorm:"type:int(5);not null;default:0;comment:是否有评价"`
	HasVrHouse               int       `json:"has_vr_house" gorm:"type:int(5);not null;default:0;comment:是否有VR房源"`
	HasShortVideo            int       `json:"has_short_video" gorm:"type:int(5);not null;default:0;comment:是否有短视频"`
	OpenDate                 time.Time `json:"open_date" gorm:"type:datetime;not null;default:'1970-01-01 00:00:00';comment:开盘日期"`
	HasVirtualView           int       `json:"has_virtual_view" gorm:"type:int(5);not null;default:0;comment:是否有虚拟看房"`
	LowestTotalPrice         int32     `json:"lowest_total_price" gorm:"type:int(20);not null;default:0;comment:最低总价"`
	PriceShowConfig          int32     `json:"price_show_config" gorm:"type:int(20);not null;default:0;comment:价格展示配置"`
	ShowPrice                int       `json:"show_price" gorm:"type:int(20);not null;default:0;comment:是否显示价格"`
	ShowPriceUnit            string    `json:"show_price_unit" gorm:"type:varchar(30);not null;comment:价格单位"`
	ShowPriceDesc            string    `json:"show_price_desc" gorm:"type:varchar(30);not null;comment:价格描述"`
	ShowPriceConfirmTime     int       `json:"show_price_confirm_time" gorm:"type:int(10);not null;default:0;comment:价格确认时间(天)"`
	PriceConfirmTime         time.Time `json:"price_confirm_time" gorm:"type:datetime;not null;default:'1970-01-01 00:00:00';comment:价格确认时间"`
	Status                   int       `json:"status" gorm:"type:int(5);not null;default:0;comment:状态"`
	SubwayDistance           string    `json:"subway_distance" gorm:"type:varchar(100);not null;comment:地铁距离"`
	IsCooperation            int       `json:"is_cooperation" gorm:"type:int(3);not null;default:0;comment:是否合作"`
	EvaluateStatus           int       `json:"evaluate_status" gorm:"type:int(5);not null;default:0;comment:评价状态"`
	ShowPriceInfo            string    `json:"show_price_info" gorm:"type:varchar(100);not null;comment:价格信息"`
	BrandID                  string    `json:"brand_id" gorm:"type:varchar(100);not null;comment:品牌ID"`
	PreloadDetailImage       string    `json:"preload_detail_image" gorm:"type:text;not null;comment:预加载详情图片"`
	ReferenceAvgPrice        float64   `json:"reference_avg_price" gorm:"type:decimal(10,2);not null;default:0;comment:参考均价"`
	ReferenceAvgPriceUnit    string    `json:"reference_avg_price_unit" gorm:"type:varchar(50);not null;comment:参考均价单位"`
	ReferenceAvgPriceDesc    string    `json:"reference_avg_price_desc" gorm:"type:varchar(50);not null;comment:参考均价描述"`
	ReferenceTotalPrice      string    `json:"reference_total_price" gorm:"type:varchar(50);not null;comment:参考总价"`
	ReferenceTotalPriceUnit  string    `json:"reference_total_price_unit" gorm:"type:varchar(50);not null;comment:参考总价单位"`
	ReferenceTotalPriceDesc  string    `json:"reference_total_price_desc" gorm:"type:varchar(50);not null;comment:参考总价描述"`
	NewSaleTags              string    `json:"new_sale_tags" gorm:"type:text;not null;comment:新房标签"`
	IsSupportOnlineSale      int       `json:"is_support_online_sale" gorm:"type:tinyint(2);not null;default:0;comment:是否支持线上销售"`
	IsUndertake              int       `json:"is_undertake" gorm:"type:tinyint(2);not null;default:0;comment:是否承接"`
	Properright              int       `json:"properright" gorm:"type:int(10);not null;default:0;comment:产权"`
	DeveloperCompany         string    `json:"developer_company" gorm:"type:varchar(100);not null;comment:开发商公司"`
	PropertyCompany          string    `json:"property_company" gorm:"type:varchar(100);not null;comment:物业公司"`
	LiveTag                  int       `json:"live_tag" gorm:"type:int(2);not null;default:0;comment:直播标签"`
	Prod                     int       `json:"prod" gorm:"type:int(2);not null;default:0;comment:产品"`
	ReferenceTotalPriceRange string    `json:"reference_total_price_range" gorm:"type:varchar(100);not null;comment:参考总价范围"`
	HouseNode                int       `json:"house_node" gorm:"type:int(20);not null;default:0;comment:房屋节点"`
	TotalPriceStart          int       `json:"total_price_start" gorm:"type:int(20);not null;default:0;comment:总价起始值"`
	TotalPriceStartUnit      string    `json:"total_price_start_unit" gorm:"type:varchar(50);not null;comment:总价起始单位"`
	AvgPriceStart            int       `json:"avg_price_start" gorm:"type:int(20);not null;default:0;comment:均价起始值"`
	AvgPriceStartUnit        string    `json:"avg_price_start_unit" gorm:"type:varchar(50);not null;comment:均价起始单位"`
	OnTime                   time.Time `json:"on_time" gorm:"type:datetime;not null;default:'1970-01-01 00:00:00';comment:开盘时间"`
	ProjectDesc              string    `json:"project_desc" gorm:"type:varchar(100);not null;comment:项目描述"`
	HasCarActivity           int       `json:"has_car_activity" gorm:"type:int(2);not null;default:0;comment:是否有车活动"`
	IsNewSale                int       `json:"is_new_sale" gorm:"type:int(2);not null;default:0;comment:是否新售"`
	FirstTags                string    `json:"first_tags" gorm:"type:varchar(200);not null;comment:第一标签"`
	MFirstTags               string    `json:"m_first_tags" gorm:"type:varchar(200);not null;comment:M第一标签"`
	FbExpoID                 string    `json:"fb_expo_id" gorm:"type:varchar(50);not null;comment:Fb博览会ID"`
	StrategyInfo             string    `json:"strategy_info" gorm:"type:varchar(300);not null;comment:策略信息"`
	RecommendLogInfo         string    `json:"recommend_log_info" gorm:"type:varchar(300);not null;comment:推荐日志信息"`
	RecommendReason          string    `json:"recommend_reason" gorm:"type:text;not null;comment:推荐原因"`
	ReferenceTotalPriceTips  string    `json:"reference_total_price_tips" gorm:"type:varchar(100);not null;comment:参考总价提示"`
	AppDetailURL             string    `json:"app_detail_url" gorm:"type:varchar(200);not null;comment:应用详情URL"`
	FilterDesc               string    `json:"filter_desc" gorm:"type:varchar(200);not null;comment:过滤描述"`
	URL                      string    `json:"url" gorm:"type:varchar(100);not null;comment:URL"`
	DownloadInfo             string    `json:"download_info" gorm:"type:text;not null;comment:下载信息"`
}

func (c *loupanPo) TableName() string {
	return "t_loupan"
}

func (c *loupanPo) Comment() string {
	return "链家新房"
}

func (repo lianjiaRepo) InsertLoupanInfo(ctx context.Context, lists []*biz.LoupanInfo) error {
	var pos []*loupanPo
	for _, info := range lists {
		pos = append(pos, &loupanPo{
			Pid:                      info.Pid,
			LoupanId:                 info.ID,
			CityID:                   info.CityID,
			CityName:                 info.CityName,
			CoverPic:                 info.CoverPic,
			MinFrameArea:             info.MinFrameArea,
			MaxFrameArea:             info.MaxFrameArea,
			DistrictName:             info.DistrictName,
			District:                 info.District,
			DistrictID:               info.DistrictID,
			BizcircleID:              info.BizcircleID,
			BizcircleName:            info.BizcircleName,
			BuildID:                  info.BuildID,
			PermitAllReady:           info.PermitAllReady,
			ProcessStatus:            info.ProcessStatus,
			ResblockFrameArea:        info.ResblockFrameArea,
			ResblockFrameAreaRange:   info.ResblockFrameAreaRange,
			ResblockFrameAreaDesc:    info.ResblockFrameAreaDesc,
			Decoration:               info.Decoration,
			Longitude:                info.Longitude,
			Latitude:                 info.Latitude,
			FrameRoomsDesc:           info.FrameRoomsDesc,
			Title:                    info.Title,
			ResblockName:             info.ResblockName,
			ResblockAlias:            info.ResblockAlias,
			Address:                  info.Address,
			StoreAddr:                info.StoreAddr,
			AvgUnitPrice:             info.AvgUnitPrice,
			AveragePrice:             info.AveragePrice,
			AddressRemark:            info.AddressRemark,
			ProjectName:              info.ProjectName,
			SpecialTags:              info.SpecialTags,
			LianjiaSpecial:           info.LianjiaSpecial,
			LianjiaSpecialComm:       info.LianjiaSpecialComm,
			DeveloperSpecial:         info.DeveloperSpecial,
			DeveloperSpecialType:     info.DeveloperSpecialType,
			DeveloperSpecialComm:     info.DeveloperSpecialComm,
			FrameRooms:               info.FrameRooms,
			ConvergedRooms:           info.ConvergedRooms,
			Tags:                     info.Tags,
			ProjectTags:              info.ProjectTags,
			HouseType:                info.HouseType,
			HouseTypeValue:           info.HouseTypeValue,
			SaleStatus:               info.SaleStatus,
			HasEvaluate:              info.HasEvaluate,
			HasVrHouse:               info.HasVrHouse,
			HasShortVideo:            info.HasShortVideo,
			OpenDate:                 info.OpenDate,
			HasVirtualView:           info.HasVirtualView,
			LowestTotalPrice:         info.LowestTotalPrice,
			PriceShowConfig:          info.PriceShowConfig,
			ShowPrice:                info.ShowPrice,
			ShowPriceUnit:            info.ShowPriceUnit,
			ShowPriceDesc:            info.ShowPriceDesc,
			ShowPriceConfirmTime:     int(info.ShowPriceConfirmTime),
			PriceConfirmTime:         info.PriceConfirmTime,
			Status:                   info.Status,
			SubwayDistance:           info.SubwayDistance,
			IsCooperation:            info.IsCooperation,
			EvaluateStatus:           info.EvaluateStatus,
			ShowPriceInfo:            info.ShowPriceInfo,
			BrandID:                  info.BrandID,
			PreloadDetailImage:       info.PreloadDetailImage,
			ReferenceAvgPrice:        info.ReferenceAvgPrice,
			ReferenceAvgPriceUnit:    info.ReferenceAvgPriceUnit,
			ReferenceAvgPriceDesc:    info.ReferenceAvgPriceDesc,
			ReferenceTotalPrice:      info.ReferenceTotalPrice,
			ReferenceTotalPriceUnit:  info.ReferenceTotalPriceUnit,
			ReferenceTotalPriceDesc:  info.ReferenceTotalPriceDesc,
			NewSaleTags:              info.NewSaleTags,
			IsSupportOnlineSale:      info.IsSupportOnlineSale,
			IsUndertake:              info.IsUndertake,
			Properright:              info.Properright,
			DeveloperCompany:         info.DeveloperCompany,
			PropertyCompany:          info.PropertyCompany,
			LiveTag:                  info.LiveTag,
			Prod:                     info.Prod,
			ReferenceTotalPriceRange: info.ReferenceTotalPriceRange,
			HouseNode:                info.HouseNode,
			TotalPriceStart:          info.TotalPriceStart,
			TotalPriceStartUnit:      info.TotalPriceStartUnit,
			AvgPriceStart:            info.AvgPriceStart,
			AvgPriceStartUnit:        info.AvgPriceStartUnit,
			OnTime:                   info.OnTime,
			ProjectDesc:              info.ProjectDesc,
			HasCarActivity:           info.HasCarActivity,
			IsNewSale:                info.IsNewSale,
			FirstTags:                info.FirstTags,
			MFirstTags:               info.MFirstTags,
			FbExpoID:                 info.FbExpoID,
			StrategyInfo:             info.StrategyInfo,
			RecommendLogInfo:         info.RecommendLogInfo,
			RecommendReason:          info.RecommendReason,
			ReferenceTotalPriceTips:  info.ReferenceTotalPriceTips,
			AppDetailURL:             info.AppDetailURL,
			FilterDesc:               info.FilterDesc,
			URL:                      info.URL,
			DownloadInfo:             info.DownloadInfo,
		})
	}
	tx := repo.data.DB(ctx).Model(&loupanPo{}).CreateInBatches(pos, batchInsertSize)
	if tx.Error != nil && !strings.Contains(tx.Error.Error(), "Duplicate entry") {
		repo.log.Errorf("insert loupan info err! %v", tx.Error)
		return tx.Error
	}
	return nil
}

func (repo lianjiaRepo) CheckLoupanExists(ctx context.Context, cityID, loupanId int64) (bool, error) {
	var loupan loupanPo
	result := repo.data.DB(ctx).Where("city_id = ? AND loupan_id = ?", cityID, loupanId).First(&loupan)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	if loupan.ID > 0 {
		return true, nil
	}
	return false, nil
}
