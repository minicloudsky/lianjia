package service

import (
	"context"
	v1 "github.com/minicloudsky/lianjia/api/lianjia/v1"
	"github.com/minicloudsky/lianjia/pkg/pagination"
)

func (s *Service) ListErshoufang(ctx context.Context, in *v1.ListErshoufangRequest) (*v1.ListErshoufangReply, error) {
	if err := in.Validate(); err != nil {
		return nil, err
	}
	erShouFangInfos, total, err := s.repo.ListErshoufang(ctx, pagination.Pagination{
		Page:     int(in.Page),
		PageSize: int(in.Size),
	}, in.Query)
	if err != nil {
		return nil, err
	}
	var erhousfangInfoPb []*v1.ErShouFangInfo
	for _, erShouFangInfo := range erShouFangInfos {
		erhousfangInfoPb = append(erhousfangInfoPb, &v1.ErShouFangInfo{
			CityId:                        erShouFangInfo.CityID,
			HouseCode:                     erShouFangInfo.HouseCode,
			ResblockId:                    erShouFangInfo.ResblockID,
			Title:                         erShouFangInfo.Title,
			Desc:                          erShouFangInfo.Desc,
			RecoDesc:                      erShouFangInfo.RecoDesc,
			TotalPriceTitle:               erShouFangInfo.TotalPrice,
			UnitPriceTitle:                erShouFangInfo.UnitPrice,
			JumpUrl:                       erShouFangInfo.JumpURL,
			ListPictureUrl:                erShouFangInfo.ListPictureURL,
			HouseStatus:                   int32(erShouFangInfo.HouseStatus),
			IsCtypeHouse:                  int32(erShouFangInfo.CtypeHouseStatus),
			FbExpoId:                      erShouFangInfo.FbExpoID,
			YeZhuTuijianStatus:            int32(erShouFangInfo.YeZhuTuijianStatus),
			HaofangStatus:                 int32(erShouFangInfo.HaofangStatus),
			YezhuPayStatus:                int32(erShouFangInfo.YezhuPayStatus),
			VrStatus:                      int32(erShouFangInfo.VrStatus),
			KeyStatus:                     int32(erShouFangInfo.KeyStatus),
			NewStatus:                     int32(erShouFangInfo.NewStatus),
			BrandTitle:                    erShouFangInfo.BrandTitle,
			HotTopDspAgentUcId:            erShouFangInfo.HotTopDspAgentUcID,
			HotTopDigV:                    erShouFangInfo.HotTopDigV,
			HotTop:                        int32(erShouFangInfo.HotTop),
			SellPriceStatus:               int32(erShouFangInfo.SellPriceStatus),
			PriceListingGovCn:             int32(erShouFangInfo.PriceListingGovCn),
			UnitPriceGov:                  int32(erShouFangInfo.UnitPriceGov),
			TotalPricePrefix:              erShouFangInfo.TotalPricePrefix,
			TotalPrice:                    float32(erShouFangInfo.TotalPriceTitle),
			TotalPriceSuffix:              erShouFangInfo.TotalPriceSuffix,
			UnitPricePrefix:               erShouFangInfo.UnitPricePrefix,
			UnitPrice:                     float32(erShouFangInfo.UnitPriceTitle),
			UnitPriceSuffix:               erShouFangInfo.UnitPriceSuffix,
			PriceInfoListTotalPricePrefix: erShouFangInfo.PriceInfoListTotalPricePrefix,
			PriceInfoListTotalPriceStr:    erShouFangInfo.PriceInfoListTotalPriceStr,
			PriceInfoListTotalPriceSuffix: erShouFangInfo.PriceInfoListTotalPriceSuffix,
			PriceInfoListUnitPricePrefix:  erShouFangInfo.UnitPricePrefix,
			PriceInfoListUnitPriceStr:     erShouFangInfo.PriceInfoListTotalPriceStr,
			PriceInfoListUnitPriceSuffix:  erShouFangInfo.UnitPriceSuffix,
			Tags:                          erShouFangInfo.Tags,
			Layout:                        erShouFangInfo.Layout,
			Area:                          float32(erShouFangInfo.Area),
			Direction:                     erShouFangInfo.Direction,
			Community:                     erShouFangInfo.Community,
			District:                      erShouFangInfo.District,
			Street:                        erShouFangInfo.Street,
			Floor:                         erShouFangInfo.Floor,
			TotalFloor:                    erShouFangInfo.TotalFloor,
		})
	}
	return &v1.ListErshoufangReply{
		Data:  erhousfangInfoPb,
		Total: int32(total),
	}, nil
}
