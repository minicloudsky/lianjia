package biz

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// json path for lianjia ershoufang api response
const (
	ershoufangRootPath        = "data.data.getErShouFangList"
	ershoufangTotalCountPath  = "totalCount"
	ershoufangReturnCountPath = "returnCount"
	ershoufangListPath        = "list"
)

type ErShouFangRaw struct {
	CityID          string      `json:"cityId"`
	HouseCode       string      `json:"houseCode"`
	ResblockID      int64       `json:"resblockId"`
	Title           string      `json:"title"`
	Desc            string      `json:"desc"`
	BangdanTitle    string      `json:"bangdanTitle"`
	RecoDesc        string      `json:"recoDesc"`
	TotalPrice      string      `json:"totalPrice"`
	UnitPrice       string      `json:"unitPrice"`
	JumpURL         string      `json:"jumpUrl"`
	ListPictureURL  string      `json:"listPictureUrl"`
	RecommendReason interface{} `json:"recommendReason"`
	HouseStatus     int         `json:"houseStatus"`
	IsCtypeHouse    bool        `json:"isCtypeHouse"`
	FbExpoID        string      `json:"fbExpoId"`
	StatusSwitch    struct {
		IsYeZhuTuijian bool `json:"isYeZhuTuijian"`
		IsHaofang      bool `json:"isHaofang"`
		IsYezhuPay     bool `json:"isYezhuPay"`
		IsVr           bool `json:"isVr"`
		IsKey          bool `json:"isKey"`
		IsNew          bool `json:"isNew"`
	} `json:"statusSwitch"`
	BrandTags struct {
		Title  string `json:"title"`
		PicURL string `json:"picUrl"`
	} `json:"brandTags"`
	HotTop struct {
		DspAgentUcID string `json:"dspAgentUcId"`
		HotTopDigV   string `json:"hotTopDigV"`
		HotTop       int    `json:"hotTop"`
	} `json:"hotTop"`
	IsSellPrice       int `json:"isSellPrice"`
	PriceListingGovCn int `json:"priceListingGovCn"`
	UnitPriceGov      int `json:"unitPriceGov"`
	TotalPriceInfo    struct {
		Prefix string `json:"prefix"`
		Title  string `json:"title"`
		Suffix string `json:"suffix"`
	} `json:"totalPriceInfo"`
	UnitPriceInfo struct {
		Prefix string `json:"prefix"`
		Title  string `json:"title"`
		Suffix string `json:"suffix"`
	} `json:"unitPriceInfo"`
	PriceInfoList struct {
		TotalPricePrefix           string `json:"totalPricePrefix"`
		TotalPriceStr              string `json:"totalPriceStr"`
		TotalPriceSuffix           string `json:"totalPriceSuffix"`
		TotalPriceNoDataDesc       string `json:"totalPriceNoDataDesc"`
		UnitPricePrefix            string `json:"unitPricePrefix"`
		UnitPriceStr               string `json:"unitPriceStr"`
		UnitPriceSuffix            string `json:"unitPriceSuffix"`
		UnitPriceNoDataDesc        string `json:"unitPriceNoDataDesc"`
		StaticTotalPriceNoDataDesc string `json:"staticTotalPriceNoDataDesc"`
		StaticUnitPriceNoDataDesc  string `json:"staticUnitPriceNoDataDesc"`
	} `json:"priceInfoList"`
	ColorTags []struct {
		Key       string `json:"key"`
		Title     string `json:"title"`
		Color     string `json:"color"`
		TextColor string `json:"textColor"`
		BgColor   string `json:"bgColor"`
	} `json:"colorTags"`
	//MarketBooth []string `json:"marketBooth"`
	PoiDistance string `json:"poiDistance"`
	PoiIcon     string `json:"poiIcon"`
}

func (raw *ErShouFangRaw) ConvertErShouFangRawToInfo() *ErShouFangInfo {
	cityId, err := strconv.Atoi(raw.CityID)
	if err != nil {
		cityId = 0
	}
	var ctypeHouseStatus int
	if raw.IsCtypeHouse {
		ctypeHouseStatus = 1
	}
	var yeZhuTuijianStatus int
	if raw.StatusSwitch.IsYeZhuTuijian {
		yeZhuTuijianStatus = 1
	}
	var haofangStatus int
	if raw.StatusSwitch.IsHaofang {
		haofangStatus = 1
	}
	var yezhuPayStatus int
	if raw.StatusSwitch.IsYezhuPay {
		yezhuPayStatus = 1
	}
	var vrStatus int
	if raw.StatusSwitch.IsVr {
		vrStatus = 1
	}
	var keyStatus int
	if raw.StatusSwitch.IsKey {
		keyStatus = 1
	}
	var newStatus int
	if raw.StatusSwitch.IsNew {
		newStatus = 1
	}
	var colorTags []string
	for _, c := range raw.ColorTags {
		colorTags = append(colorTags, c.Key)
	}
	totalPriceTitle, err := strconv.ParseFloat(raw.TotalPriceInfo.Title, 64)
	if err != nil {
		totalPriceTitle = 0
	}
	unitPriceTitle, err := strconv.ParseFloat(strings.Replace(raw.UnitPriceInfo.Title, ",", "", -1), 64)
	if err != nil {
		unitPriceTitle = 0
	}
	ershoufangHousedesc := ErshoufangHouseDesc(raw.Desc)
	descDetail := ershoufangHousedesc.Detail()
	ershoufangRecoDesc := RecoDesc(raw.RecoDesc)
	recoDesc := ershoufangRecoDesc.Detail()
	info := &ErShouFangInfo{
		CityID:                        int64(cityId),
		HouseCode:                     raw.HouseCode,
		ResblockID:                    raw.ResblockID,
		Title:                         raw.Title,
		Desc:                          raw.Desc,
		BangdanTitle:                  raw.BangdanTitle,
		RecoDesc:                      raw.RecoDesc,
		TotalPrice:                    raw.TotalPrice,
		UnitPrice:                     strings.Replace(raw.UnitPrice, ",", "", -1),
		JumpURL:                       raw.JumpURL,
		ListPictureURL:                raw.ListPictureURL,
		HouseStatus:                   raw.HouseStatus,
		CtypeHouseStatus:              ctypeHouseStatus,
		FbExpoID:                      raw.FbExpoID,
		YeZhuTuijianStatus:            yeZhuTuijianStatus,
		HaofangStatus:                 haofangStatus,
		YezhuPayStatus:                yezhuPayStatus,
		VrStatus:                      vrStatus,
		KeyStatus:                     keyStatus,
		NewStatus:                     newStatus,
		Tags:                          strings.Join(colorTags, ","),
		BrandTitle:                    raw.BrandTags.Title,
		HotTopDspAgentUcID:            raw.HotTop.DspAgentUcID,
		HotTopDigV:                    raw.HotTop.HotTopDigV,
		HotTop:                        raw.HotTop.HotTop,
		SellPriceStatus:               raw.IsSellPrice,
		PriceListingGovCn:             raw.PriceListingGovCn,
		UnitPriceGov:                  raw.UnitPriceGov,
		TotalPricePrefix:              raw.TotalPriceInfo.Prefix,
		TotalPriceTitle:               totalPriceTitle,
		TotalPriceSuffix:              raw.TotalPriceInfo.Suffix,
		UnitPricePrefix:               raw.UnitPriceInfo.Prefix,
		UnitPriceTitle:                unitPriceTitle,
		UnitPriceSuffix:               raw.UnitPriceInfo.Suffix,
		PriceInfoListTotalPricePrefix: raw.PriceInfoList.TotalPricePrefix,
		PriceInfoListTotalPriceStr:    raw.PriceInfoList.TotalPriceStr,
		PriceInfoListTotalPriceSuffix: raw.PriceInfoList.TotalPriceSuffix,
		Layout:                        descDetail.Layout,
		Area:                          descDetail.Area,
		Direction:                     descDetail.Direction,
		Community:                     descDetail.Community,
		District:                      recoDesc.District,
		Street:                        recoDesc.Street,
		Floor:                         recoDesc.Floor,
		TotalFloor:                    recoDesc.TotalFloor,
	}
	return info
}

type ErShouFangInfo struct {
	CityID                        int64   `json:"city_id"`
	HouseCode                     string  `json:"house_code"`
	ResblockID                    int64   `json:"resblock_id"`
	Title                         string  `json:"title"`
	Desc                          string  `json:"desc"`
	BangdanTitle                  string  `json:"bangdan_title"`
	RecoDesc                      string  `json:"reco_desc"`
	TotalPrice                    string  `json:"total_price"`
	UnitPrice                     string  `json:"unit_price"`
	JumpURL                       string  `json:"jump_url"`
	ListPictureURL                string  `json:"list_picture_url"`
	HouseStatus                   int     `json:"house_status"`
	CtypeHouseStatus              int     `json:"is_ctype_house"`
	FbExpoID                      string  `json:"fb_expo_id"`
	YeZhuTuijianStatus            int     `json:"ye_zhu_tuijian_status"`
	HaofangStatus                 int     `json:"haofang_status"`
	YezhuPayStatus                int     `json:"yezhu_pay_status"`
	VrStatus                      int     `json:"vr_status"`
	KeyStatus                     int     `json:"key_status"`
	NewStatus                     int     `json:"new_status"`
	BrandTitle                    string  `json:"brand_title"`
	HotTopDspAgentUcID            string  `json:"hot_top_dsp_agent_uc_id"`
	HotTopDigV                    string  `json:"hot_top_dig_v"`
	HotTop                        int     `json:"hot_top"`
	SellPriceStatus               int     `json:"sell_price_status"`
	PriceListingGovCn             int     `json:"price_listing_gov_cn"`
	UnitPriceGov                  int     `json:"unit_price_gov"`
	TotalPricePrefix              string  `json:"total_price_prefix"`
	TotalPriceTitle               float64 `json:"total_price_title"`
	TotalPriceSuffix              string  `json:"total_price_suffix"`
	UnitPricePrefix               string  `json:"unit_price_prefix"`
	UnitPriceTitle                float64 `json:"unit_price_title"`
	UnitPriceSuffix               string  `json:"unit_price_suffix"`
	PriceInfoListTotalPricePrefix string  `json:"price_info_list_total_price_prefix"`
	PriceInfoListTotalPriceStr    string  `json:"price_info_list_total_price_str"`
	PriceInfoListTotalPriceSuffix string  `json:"price_info_list_total_price_suffix"`
	PriceInfoListUnitPricePrefix  string  `json:"price_info_list_unit_price_prefix"`
	PriceInfoListUnitPriceStr     string  `json:"price_info_list_unit_price_str"`
	PriceInfoListUnitPriceSuffix  string  `json:"price_info_list_unit_price_suffix"`
	Tags                          string  `json:"tags"`
	Layout                        string  `json:"layout"`
	Area                          float64 `json:"area"`
	Direction                     string  `json:"direction"`
	Community                     string  `json:"community"`
	District                      string  `json:"district"`
	Street                        string  `json:"street"`
	Floor                         string  `json:"floor"`
	TotalFloor                    int32   `json:"total_floor"`
}

type ErshoufangHouseDesc string

type ErshoufangHouseDescDetail struct {
	Layout    string
	Area      float64
	Direction string
	Community string
}

func (e ErshoufangHouseDesc) Detail() ErshoufangHouseDescDetail {
	desc := strings.Split(string(e), "/")
	if len(desc) == 4 {
		area, err := strconv.ParseFloat(strings.Replace(desc[1], "m²", "", -1), 64)
		if err != nil {
			area = 0
		}
		return ErshoufangHouseDescDetail{
			Layout:    desc[0],
			Area:      area,
			Direction: desc[2],
			Community: desc[3],
		}
	} else {
		return ErshoufangHouseDescDetail{}
	}
}

type RecoDesc string

type RecoDescDetail struct {
	District   string
	Street     string
	Floor      string
	TotalFloor int32
	Direction  string
}

func (r RecoDesc) Detail() RecoDescDetail {
	desc := strings.Split(string(r), "/")
	if len(desc) == 3 {
		recoDescDetail := RecoDescDetail{}
		res := strings.Split(desc[0], " ")
		if len(res) == 2 {
			recoDescDetail.District = res[0]
			recoDescDetail.Street = res[1]
		} else if len(res) == 1 {
			recoDescDetail.District = res[0]
		}
		res = strings.Split(desc[1], " ")
		if len(res) == 2 {
			recoDescDetail.Floor = res[0]
			if strings.Contains(res[1], "共") {
				res[1] = strings.Replace(res[1], "共", "", -1)
			}
			if strings.Contains(res[1], "层") {
				res[1] = strings.Replace(res[1], "层", "", -1)
			}
			totalFloor, err := strconv.ParseInt(res[1], 10, 32)
			if err != nil {
				totalFloor = 0
			}
			recoDescDetail.TotalFloor = int32(totalFloor)
		} else if len(res) == 1 {
			recoDescDetail.Floor = res[0]
		}
		recoDescDetail.Direction = desc[2]
		return recoDescDetail
	} else {
		return RecoDescDetail{}
	}
}

func (uc *LianjiaUsecase) ListCityErshouFang(ctx context.Context, city *CityInfo) error {
	defer func(uc *LianjiaUsecase, ctx context.Context, city *CityInfo) {
		err := uc.finishCurrentTask(ctx, HouseTypeErshoufang, city)
		if err != nil {
			uc.log.Errorf("finish ershoufang task err! %v", err)
			return
		}
	}(uc, ctx, city)
	var client *resty.Client
	client = resty.New()
	var headers http.Header
	var url string
	headers = http.Header{}
	headers.Add("User-Agent", RandomUA())
	client.Header = headers
	client.SetRetryCount(3)
	client.SetRetryMaxWaitTime(5 * time.Second)
	url = fmt.Sprintf(ershoufangListUrl, city.CityCode, 1)
	resp, err := client.R().Get(url)
	if err != nil {
		uc.log.Errorf("ListErshoufang err! url: %s err: %v", url, err)
	}
	parser := gjson.Get(resp.String(), ershoufangRootPath)
	if parser.String() == "null" || parser.String() == "" {
		uc.log.Warnf("ListCityErshouFang city %s haven't ershoufang! url: %s", city.Name, url)
		return nil
	}
	totalCount := parser.Get(ershoufangTotalCountPath).Num
	returnCount := parser.Get(ershoufangReturnCountPath).Num
	pages := int(math.Ceil(totalCount / returnCount))
	for page := 1; page <= pages; page++ {
		headers = http.Header{}
		headers.Add("User-Agent", RandomUA())
		client.Header = headers
		url = fmt.Sprintf(ershoufangListUrl, city.CityCode, page)
		resp, err := client.R().Get(url)
		if err != nil {
			uc.log.Errorf("ListErshoufang err! url: %s err: %v", url, err)
		}
		parser := gjson.Get(resp.String(), ershoufangRootPath)
		houseList := parser.Get(ershoufangListPath).Array()
		if len(houseList) > 0 {
			messages := make([]Message, 0)
			for _, house := range houseList {
				hashValue := uc.md5(house.String())
				key := LianjiaErshoufangIdempotentKey + hashValue
				if val, err := uc.repo.GetKey(ctx, key); err == nil && val != "" {
					continue
				} else {
					messages = append(messages, Message{
						Content: []byte(house.String()),
					})
					if err := uc.SetNXKey(ctx, key, time.Now().Unix(),
						6*time.Hour); err != nil {
						continue
					}
				}
			}
			err := uc.Send(ctx, messages, HouseTypeErshoufang)
			if err != nil {
				uc.log.Errorf("send message err! err: %v", err)
			}
		}
	}
	return nil
}

func (uc *LianjiaUsecase) HandleErshoufangMessage(ctx context.Context) error {
	var slices []*ErShouFangInfo
	for {
		select {
		case msgs := <-ErshoufangProcessChan:
			for _, msg := range msgs {
				if len(msg.Content) > 0 {
					var raw *ErShouFangRaw
					err := json.Unmarshal(msg.Content, &raw)
					if err != nil {
						uc.log.Errorf("Unmarshal ershoufang info err! ershoufangInfo: %s err: %v",
							string(msg.Content), err)
						continue
					}
					ershoufangInfo := raw.ConvertErShouFangRawToInfo()
					exist, err := uc.repo.CheckErShouFangExists(ctx, ershoufangInfo.CityID, ershoufangInfo.HouseCode)
					if err != nil {
						uc.log.Errorf("check ershoufang exists err! ershoufangInfo: %s err: %v",
							string(msg.Content), err)
						continue
					}
					if !exist {
						slices = append(slices, ershoufangInfo)
					}
				}
			}
			if len(slices) > 0 || (len(slices) >= insertBatchSize || uc.IsErshouFangTaskFinish(ctx)) {
				err := uc.repo.InsertErshoufangInfo(ctx, slices)
				if err != nil {
					uc.log.Errorf("insert ershoufang info err! %v", err)
				}
				slices = make([]*ErShouFangInfo, 0)
			}
		}
	}
}

func (uc *LianjiaUsecase) IsErshouFangTaskFinish(ctx context.Context) bool {
	unFinishTaskStr, err := uc.repo.GetKey(ctx, LianjiaErshoufangRunningCrawlTaskKey)
	if err != nil && unFinishTaskStr != "" {
		uc.log.Errorf("get ershoufang unfinish task key err! %v", err)
		return false
	}
	if unFinishTaskStr == "" {
		return true
	}
	unFinishTasks, err := strconv.ParseInt(unFinishTaskStr, 10, 64)
	if err != nil {
		uc.log.Errorf("parse ershoufang task num err! %v", err)
		return false
	}
	if unFinishTasks > 0 {
		return false
	}

	return true
}
