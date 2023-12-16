package data

import (
	"context"
	"errors"
	"github.com/minicloudsky/lianjia/internal/biz"
	"github.com/minicloudsky/lianjia/pkg/pagination"
	"gorm.io/gorm"
	"strings"
)

const batchInsertSize = 1

type erShouFangInfoPo struct {
	BaseModel
	CityID                        int64   `gorm:"column:city_id;type:bigint;default:0;not null;comment:城市ID;index:idx_city_district_street_community" json:"city_id"`
	HouseCode                     string  `gorm:"column:house_code;type:varchar(15);default:'';not null;comment:房屋编码" json:"house_code"`
	ResblockID                    int64   `gorm:"column:resblock_id;type:bigint;default:0;not null;comment:小区ID" json:"resblock_id"`
	Title                         string  `gorm:"column:title;type:varchar(100);default:'';not null;comment:标题" json:"title"`
	HouseDesc                     string  `gorm:"column:house_desc;type:varchar(100);default:'';not null;comment:描述" json:"house_desc"`
	RecoDesc                      string  `gorm:"column:reco_desc;type:varchar(150);default:'';not null;comment:推荐描述" json:"reco_desc"`
	District                      string  `gorm:"column:district;type:varchar(50);default:'';not null;comment:区域;index:idx_city_district_street_community" json:"district"`
	Street                        string  `gorm:"column:street;type:varchar(50);default:'';not null;comment:街道;index:idx_city_district_street_community" json:"street"`
	Community                     string  `gorm:"column:community;type:varchar(50);default:'';not null;comment:小区;index:idx_city_district_street_community" json:"community"`
	TotalFloor                    int32   `gorm:"column:total_floor;type:int(10);default:0;not null;comment:总楼层" json:"total_floor"`
	Layout                        string  `gorm:"column:layout;type:varchar(50);default:'';not null;comment:户型" json:"layout"`
	Area                          float64 `gorm:"column:area;type:DECIMAL(10, 3);default:0;not null;comment:面积" json:"area"`
	Direction                     string  `gorm:"column:direction;type:varchar(50);default:'';not null;comment:朝向" json:"direction"`
	Floor                         string  `gorm:"column:floor;type:varchar(50);default:'';not null;comment:楼层" json:"floor"`
	TotalPrice                    float64 `gorm:"column:total_price;type:DECIMAL(10, 3);default:0;not null;comment:总价" json:"total_price"`
	UnitPrice                     float64 `gorm:"column:unit_price;type:DECIMAL(10, 3);default:0;not null;comment:单价" json:"unit_price"`
	TotalPriceTitle               string  `gorm:"column:total_price_title;type:varchar(20);default:'';not null;comment:总价标题" json:"total_price_title"`
	UnitPriceTitle                string  `gorm:"column:unit_price_title;type:varchar(25);default:'';not null;comment:单价标题" json:"unit_price_title"`
	TotalPricePrefix              string  `gorm:"column:total_price_prefix;type:varchar(30);default:'';not null;comment:总价前缀" json:"total_price_prefix"`
	TotalPriceSuffix              string  `gorm:"column:total_price_suffix;type:varchar(20);default:'';not null;comment:总价后缀" json:"total_price_suffix"`
	UnitPricePrefix               string  `gorm:"column:unit_price_prefix;type:varchar(20);default:'';not null;comment:单价前缀" json:"unit_price_prefix"`
	UnitPriceSuffix               string  `gorm:"column:unit_price_suffix;type:varchar(20);default:'';not null;comment:单价后缀" json:"unit_price_suffix"`
	JumpURL                       string  `gorm:"column:jump_url;type:varchar(100);default:'';not null;comment:跳转URL" json:"jump_url"`
	ListPictureURL                string  `gorm:"column:list_picture_url;type:varchar(500);default:'';not null;comment:列表图片URL" json:"list_picture_url"`
	HouseStatus                   int     `gorm:"column:house_status;type:int(2);default:0;not null;comment:房屋状态" json:"house_status"`
	CtypeHouseStatus              int     `gorm:"column:is_ctype_house;type:int(2);default:0;not null;comment:是否C型房屋" json:"is_ctype_house"`
	FbExpoID                      string  `gorm:"column:fb_expo_id;type:varchar(30);default:'';not null;comment:Fb博览会ID" json:"fb_expo_id"`
	YeZhuTuijianStatus            int     `gorm:"column:ye_zhu_tuijian_status;type:int(2);default:0;not null;comment:业主推荐状态" json:"ye_zhu_tuijian_status"`
	HaofangStatus                 int     `gorm:"column:haofang_status;type:int(2);default:0;not null;comment:好房状态" json:"haofang_status"`
	YezhuPayStatus                int     `gorm:"column:yezhu_pay_status;type:int(2);default:0;not null;comment:业主支付状态" json:"yezhu_pay_status"`
	VrStatus                      int     `gorm:"column:vr_status;type:int(2);default:0;not null;comment:VR状态" json:"vr_status"`
	KeyStatus                     int     `gorm:"column:key_status;type:int(2);default:0;not null;comment:钥匙状态" json:"key_status"`
	NewStatus                     int     `gorm:"column:new_status;type:int(2);default:0;not null;comment:新房状态" json:"new_status"`
	BrandTitle                    string  `gorm:"column:brand_title;type:varchar(20);default:'';not null;comment:品牌标题" json:"brand_title"`
	HotTopDspAgentUcID            string  `gorm:"column:hot_top_dsp_agent_uc_id;type:varchar(255);default:'';not null;comment:热门置顶DSP代理UC ID" json:"hot_top_dsp_agent_uc_id"`
	HotTopDigV                    string  `gorm:"column:hot_top_dig_v;type:varchar(30);default:'';not null;comment:热门置顶数字V" json:"hot_top_dig_v"`
	HotTop                        int     `gorm:"column:hot_top;type:bigint;default:0;not null;comment:热门置顶" json:"hot_top"`
	SellPriceStatus               int     `gorm:"column:sell_price_status;type:int(2);default:0;not null;comment:售价状态" json:"sell_price_status"`
	PriceListingGovCn             int     `gorm:"column:price_listing_gov_cn;type:int(2);default:0;not null;comment:政府挂牌价" json:"price_listing_gov_cn"`
	UnitPriceGov                  int     `gorm:"column:unit_price_gov;type:int(2);default:0;not null;comment:政府单价" json:"unit_price_gov"`
	Tags                          string  `gorm:"column:tags;type:varchar(100);default:'';not null;comment:标签" json:"tags"`
	PriceInfoListTotalPricePrefix string  `gorm:"column:price_info_list_total_price_prefix;type:varchar(30);default:'';not null;comment:价格信息列表总价前缀" json:"price_info_list_total_price_prefix"`
	PriceInfoListTotalPriceStr    string  `gorm:"column:price_info_list_total_price_str;type:varchar(40);default:'';not null;comment:价格信息列表总价字符串" json:"price_info_list_total_price_str"`
	PriceInfoListTotalPriceSuffix string  `gorm:"column:price_info_list_total_price_suffix;type:varchar(20);default:'';not null;comment:价格信息列表总价后缀" json:"price_info_list_total_price_suffix"`
}

func (c *erShouFangInfoPo) TableName() string {
	return "t_ershoufang"
}

func (c *erShouFangInfoPo) Comment() string {
	return "链家二手房"
}

func (repo lianjiaRepo) InsertErshoufangInfo(ctx context.Context, infos []*biz.ErShouFangInfo) error {
	var pos []*erShouFangInfoPo
	for _, info := range infos {
		po := &erShouFangInfoPo{
			CityID:                        info.CityID,
			HouseCode:                     info.HouseCode,
			ResblockID:                    info.ResblockID,
			Title:                         info.Title,
			HouseDesc:                     info.Desc,
			RecoDesc:                      info.RecoDesc,
			TotalPrice:                    info.TotalPriceTitle,
			UnitPrice:                     info.UnitPriceTitle,
			JumpURL:                       info.JumpURL,
			ListPictureURL:                info.ListPictureURL,
			HouseStatus:                   info.HouseStatus,
			CtypeHouseStatus:              info.VrStatus,
			FbExpoID:                      info.FbExpoID,
			YeZhuTuijianStatus:            info.YeZhuTuijianStatus,
			HaofangStatus:                 info.HaofangStatus,
			YezhuPayStatus:                info.YezhuPayStatus,
			VrStatus:                      info.VrStatus,
			KeyStatus:                     info.KeyStatus,
			NewStatus:                     info.NewStatus,
			BrandTitle:                    info.BrandTitle,
			HotTopDspAgentUcID:            info.HotTopDspAgentUcID,
			HotTopDigV:                    info.HotTopDigV,
			HotTop:                        info.HotTop,
			SellPriceStatus:               info.SellPriceStatus,
			PriceListingGovCn:             info.PriceListingGovCn,
			UnitPriceGov:                  info.UnitPriceGov,
			TotalPricePrefix:              info.TotalPricePrefix,
			TotalPriceTitle:               info.TotalPrice,
			TotalPriceSuffix:              info.TotalPriceSuffix,
			UnitPricePrefix:               info.UnitPricePrefix,
			UnitPriceTitle:                info.UnitPrice,
			Tags:                          info.Tags,
			UnitPriceSuffix:               info.UnitPriceSuffix,
			PriceInfoListTotalPricePrefix: info.PriceInfoListTotalPricePrefix,
			PriceInfoListTotalPriceStr:    info.PriceInfoListTotalPriceStr,
			PriceInfoListTotalPriceSuffix: info.PriceInfoListTotalPriceSuffix,
			Layout:                        info.Layout,
			Area:                          info.Area,
			Direction:                     info.Direction,
			Community:                     info.Community,
			District:                      info.District,
			Street:                        info.Street,
			Floor:                         info.Floor,
			TotalFloor:                    info.TotalFloor,
		}
		pos = append(pos, po)
	}
	tx := repo.data.DB(ctx).Model(&erShouFangInfoPo{}).CreateInBatches(pos, batchInsertSize)
	if tx.Error != nil && !strings.Contains(tx.Error.Error(), "Duplicate entry") {
		return tx.Error
	}
	return nil
}

func (repo lianjiaRepo) ListErshoufang(ctx context.Context, p pagination.Pagination, query string) (info []*biz.ErShouFangInfo, total int64, err error) {
	var pos []erShouFangInfoPo
	tx := repo.data.DB(ctx).Model(&erShouFangInfoPo{})
	txTotal := repo.data.DB(ctx).Model(&erShouFangInfoPo{})
	if query != "" {
		tx = tx.Where("title like ?", "%"+query+"%")
		txTotal = txTotal.Where("title like ?", "%"+query+"%")
	}
	txTotal = txTotal.Count(&total)
	if txTotal.Error != nil {
		repo.log.Errorf("list ershoufang err! %v", txTotal.Error)
		return nil, 0, txTotal.Error
	}
	tx = tx.Offset(p.GetOffset()).Limit(p.PageSize).Find(&pos)
	if tx.Error != nil {
		repo.log.Errorf("list ershoufang err! %v", tx.Error)
		return nil, 0, tx.Error
	}
	var infos []*biz.ErShouFangInfo
	for _, po := range pos {
		infos = append(infos, &biz.ErShouFangInfo{
			CityID:                        po.CityID,
			HouseCode:                     po.HouseCode,
			ResblockID:                    po.ResblockID,
			Title:                         po.Title,
			Desc:                          po.HouseDesc,
			BangdanTitle:                  po.BrandTitle,
			RecoDesc:                      po.RecoDesc,
			TotalPrice:                    po.TotalPriceTitle,
			UnitPrice:                     po.UnitPriceTitle,
			JumpURL:                       po.JumpURL,
			ListPictureURL:                po.ListPictureURL,
			HouseStatus:                   po.HouseStatus,
			CtypeHouseStatus:              po.CtypeHouseStatus,
			FbExpoID:                      po.FbExpoID,
			YeZhuTuijianStatus:            po.YeZhuTuijianStatus,
			HaofangStatus:                 po.HaofangStatus,
			YezhuPayStatus:                po.YezhuPayStatus,
			VrStatus:                      po.VrStatus,
			KeyStatus:                     po.KeyStatus,
			NewStatus:                     po.NewStatus,
			BrandTitle:                    po.BrandTitle,
			HotTopDspAgentUcID:            po.HotTopDspAgentUcID,
			HotTopDigV:                    po.HotTopDigV,
			HotTop:                        po.HotTop,
			SellPriceStatus:               po.SellPriceStatus,
			PriceListingGovCn:             po.PriceListingGovCn,
			UnitPriceGov:                  po.UnitPriceGov,
			TotalPricePrefix:              po.TotalPricePrefix,
			TotalPriceTitle:               po.TotalPrice,
			TotalPriceSuffix:              po.TotalPriceSuffix,
			UnitPricePrefix:               po.UnitPricePrefix,
			UnitPriceTitle:                po.UnitPrice,
			UnitPriceSuffix:               po.UnitPriceSuffix,
			PriceInfoListTotalPricePrefix: po.PriceInfoListTotalPricePrefix,
			PriceInfoListTotalPriceStr:    po.PriceInfoListTotalPriceStr,
			PriceInfoListTotalPriceSuffix: po.PriceInfoListTotalPriceSuffix,
			PriceInfoListUnitPricePrefix:  po.UnitPricePrefix,
			PriceInfoListUnitPriceStr:     po.PriceInfoListTotalPriceStr,
			PriceInfoListUnitPriceSuffix:  po.UnitPriceSuffix,
			Tags:                          po.Tags,
			Layout:                        po.Layout,
			Area:                          po.Area,
			Direction:                     po.Direction,
			Community:                     po.Community,
			District:                      po.District,
			Street:                        po.Street,
			Floor:                         po.Floor,
			TotalFloor:                    po.TotalFloor,
		})
	}
	return infos, total, nil
}

func (repo lianjiaRepo) CheckErShouFangExists(ctx context.Context, cityID int64, houseCode string) (bool, error) {
	var ershoufang erShouFangInfoPo
	result := repo.data.DB(ctx).Where("city_id = ? AND house_code = ?", cityID, houseCode).First(&ershoufang)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, result.Error
	}
	if ershoufang.ID > 0 {
		return true, nil
	}
	return false, nil
}
