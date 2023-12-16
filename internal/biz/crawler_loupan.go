package biz

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"net/http"
	"strconv"
	"strings"
	"time"
)

const (
	loupanRootPath           = "data"
	loupanTotalPageCountPath = "total"
	loupanListPath           = "body._resblock_list"
	rateLimitKeyword         = "您请求的次数过多"
)

var loupanSleepTime = time.Second

type LoupanJson struct {
	Pid                    string        `json:"pid"`
	ID                     string        `json:"id"`
	CityID                 string        `json:"city_id"`
	CityName               string        `json:"city_name"`
	CoverPic               string        `json:"cover_pic"`
	MinFrameArea           string        `json:"min_frame_area"`
	MaxFrameArea           string        `json:"max_frame_area"`
	DistrictName           string        `json:"district_name"`
	District               string        `json:"district"`
	DistrictID             string        `json:"district_id"`
	BizcircleID            string        `json:"bizcircle_id"`
	BizcircleName          string        `json:"bizcircle_name"`
	BuildID                string        `json:"build_id"`
	PermitAllReady         string        `json:"permit_all_ready"`
	ProcessStatus          string        `json:"process_status"`
	ResblockFrameArea      string        `json:"resblock_frame_area"`
	ResblockFrameAreaRange string        `json:"resblock_frame_area_range"`
	ResblockFrameAreaDesc  string        `json:"resblock_frame_area_desc"`
	Decoration             string        `json:"decoration"`
	Longitude              string        `json:"longitude"`
	Latitude               string        `json:"latitude"`
	FrameRoomsDesc         string        `json:"frame_rooms_desc"`
	Title                  string        `json:"title"`
	ResblockName           string        `json:"resblock_name"`
	ResblockAlias          string        `json:"resblock_alias"`
	Address                string        `json:"address"`
	StoreAddr              string        `json:"store_addr"`
	AvgUnitPrice           string        `json:"avg_unit_price"`
	AveragePrice           string        `json:"average_price"`
	AddressRemark          string        `json:"address_remark"`
	ProjectName            string        `json:"project_name"`
	SpecialTags            []interface{} `json:"special_tags"`
	Special                struct {
		LianjiaSpecial       []interface{} `json:"lianjia_special"`
		LianjiaSpecialComm   string        `json:"lianjia_special_comm"`
		DeveloperSpecial     []interface{} `json:"developer_special"`
		DeveloperSpecialType string        `json:"developer_special_type"`
		DeveloperSpecialComm string        `json:"developer_special_comm"`
	} `json:"special"`
	FrameRooms     []interface{} `json:"frame_rooms"`
	ConvergedRooms []struct {
		BedroomCount string `json:"bedroom_count"`
		AreaRange    string `json:"area_range"`
	} `json:"converged_rooms"`
	Tags        []string `json:"tags"`
	ProjectTags []struct {
		ID    string `json:"id"`
		Desc  string `json:"desc"`
		Color string `json:"color"`
	} `json:"project_tags"`
	HouseType            string `json:"house_type"`
	HouseTypeValue       string `json:"house_type_value"`
	SaleStatus           string `json:"sale_status"`
	HasEvaluate          string `json:"has_evaluate"`
	HasVrHouse           string `json:"has_vr_house"`
	HasShortVideo        string `json:"has_short_video"`
	OpenDate             string `json:"open_date"`
	HasVirtualView       string `json:"has_virtual_view"`
	LowestTotalPrice     string `json:"lowest_total_price"`
	PriceShowConfig      string `json:"price_show_config"`
	ShowPrice            string `json:"show_price"`
	ShowPriceUnit        string `json:"show_price_unit"`
	ShowPriceDesc        string `json:"show_price_desc"`
	ShowPriceConfirmTime string `json:"show_price_confirm_time"`
	PriceConfirmTime     string `json:"price_confirm_time"`
	Status               string `json:"status"`
	SubwayDistance       string `json:"subway_distance"`
	IsCooperation        string `json:"is_cooperation"`
	EvaluateStatus       string `json:"evaluate_status"`
	ShowPriceInfo        string `json:"show_price_info"`
	BrandID              string `json:"brand_id"`
	PreloadDetailImage   []struct {
		ID                 string      `json:"id"`
		ImageID            string      `json:"image_id"`
		TypeID             string      `json:"type_id"`
		TypeName           string      `json:"type_name"`
		ImageURL           string      `json:"image_url"`
		Title              string      `json:"title"`
		Desc               string      `json:"desc"`
		Extend             interface{} `json:"extend"`
		VrInfo             interface{} `json:"vr_info"`
		VideoInfo          interface{} `json:"video_info"`
		ImageListBlurryURL string      `json:"image_list_blurry_url"`
		ImageSizeURL       string      `json:"image_size_url"`
		ImageListSizeURL   string      `json:"image_list_size_url"`
	} `json:"preload_detail_image"`
	ReferenceAvgPrice        string        `json:"reference_avg_price"`
	ReferenceAvgPriceUnit    string        `json:"reference_avg_price_unit"`
	ReferenceAvgPriceDesc    string        `json:"reference_avg_price_desc"`
	ReferenceTotalPrice      string        `json:"reference_total_price"`
	ReferenceTotalPriceUnit  string        `json:"reference_total_price_unit"`
	ReferenceTotalPriceDesc  string        `json:"reference_total_price_desc"`
	NewSaleTags              []interface{} `json:"new_sale_tags"`
	IsSupportOnlineSale      string        `json:"is_support_online_sale"`
	IsUndertake              string        `json:"is_undertake"`
	Properright              string        `json:"properright"`
	DeveloperCompany         []string      `json:"developer_company"`
	PropertyCompany          []string      `json:"property_company"`
	LiveTag                  string        `json:"live_tag"`
	Prod                     string        `json:"prod"`
	ReferenceTotalPriceRange struct {
		Price     string `json:"price"`
		PriceUnit string `json:"price_unit"`
		PriceDesc string `json:"price_desc"`
	} `json:"reference_total_price_range"`
	SaleStatusColor         string        `json:"sale_status_color"`
	HouseTypeColor          string        `json:"house_type_color"`
	HouseNode               string        `json:"house_node"`
	TotalPriceStart         string        `json:"total_price_start"`
	TotalPriceStartUnit     string        `json:"total_price_start_unit"`
	AvgPriceStart           string        `json:"avg_price_start"`
	AvgPriceStartUnit       string        `json:"avg_price_start_unit"`
	OnTime                  string        `json:"on_time"`
	ProjectDesc             string        `json:"project_desc"`
	HasCarActivity          string        `json:"has_car_activity"`
	IsNewSale               string        `json:"is_new_sale"`
	FirstTags               []interface{} `json:"first_tags"`
	MFirstTags              []interface{} `json:"m_first_tags"`
	FbExpoID                string        `json:"fb_expo_id"`
	StrategyInfo            string        `json:"strategy_info"`
	RecommendLogInfo        string        `json:"recommend_log_info"`
	RecommendReason         interface{}   `json:"recommend_reason"`
	ReferenceTotalPriceTips string        `json:"reference_total_price_tips"`
	UserInfo                struct {
		IsFollow string `json:"is_follow"`
	} `json:"user_info"`
	AppDetailURL string `json:"app_detail_url"`
	FilterDesc   string `json:"filter_desc"`
	URL          string `json:"url"`
	DownloadInfo struct {
		Type          string `json:"type"`
		DownloadURL   string `json:"downloadUrl"`
		Schema        string `json:"schema"`
		UniversalLink string `json:"universalLink"`
	} `json:"download_info"`
}

func (loupanJson *LoupanJson) ConvertToLoupanInfo() (*LoupanInfo, error) {
	pid, _ := strconv.Atoi(loupanJson.Pid)
	id, _ := strconv.Atoi(loupanJson.ID)
	minFrameArea, _ := strconv.Atoi(loupanJson.MinFrameArea)
	maxFrameArea, _ := strconv.Atoi(loupanJson.MaxFrameArea)
	cityID, _ := strconv.Atoi(loupanJson.CityID)
	districtID, _ := strconv.Atoi(loupanJson.DistrictID)
	bizcircleID, _ := strconv.Atoi(loupanJson.BizcircleID)
	buildID, _ := strconv.ParseInt(loupanJson.BuildID, 10, 64)
	permitAllReady, _ := strconv.Atoi(loupanJson.PermitAllReady)
	processStatus, _ := strconv.Atoi(loupanJson.ProcessStatus)
	longitude, _ := strconv.ParseFloat(loupanJson.Longitude, 64)
	latitude, _ := strconv.ParseFloat(loupanJson.Latitude, 64)
	avgUnitPrice, _ := strconv.ParseFloat(loupanJson.AvgUnitPrice, 64)
	averagePrice, _ := strconv.ParseFloat(loupanJson.AveragePrice, 64)
	lowestTotalPrice, _ := strconv.Atoi(loupanJson.LowestTotalPrice)
	priceShowConfig, _ := strconv.Atoi(loupanJson.PriceShowConfig)
	showPrice, _ := strconv.Atoi(loupanJson.ShowPrice)
	status, _ := strconv.Atoi(loupanJson.Status)
	isCooperation, _ := strconv.Atoi(loupanJson.IsCooperation)
	evaluateStatus, _ := strconv.Atoi(loupanJson.EvaluateStatus)
	hasVrHouse, _ := strconv.Atoi(loupanJson.HasVrHouse)
	hasShortVideo, _ := strconv.Atoi(loupanJson.HasShortVideo)
	hasVirtualView, _ := strconv.Atoi(loupanJson.HasVirtualView)
	hasEvaluate, _ := strconv.Atoi(loupanJson.HasEvaluate)
	isSupportOnlineSale, _ := strconv.Atoi(loupanJson.IsSupportOnlineSale)
	isUndertake, _ := strconv.Atoi(loupanJson.IsUndertake)
	liveTag, _ := strconv.Atoi(loupanJson.LiveTag)
	prod, _ := strconv.Atoi(loupanJson.Prod)
	houseNode, _ := strconv.Atoi(loupanJson.HouseNode)
	totalPriceStart, _ := strconv.Atoi(loupanJson.TotalPriceStart)
	avgPriceStart, _ := strconv.Atoi(loupanJson.AvgPriceStart)
	onTime, err := time.Parse(time.DateTime, loupanJson.OnTime)
	if err != nil {
		onTime = time.Time{}
	}

	openDate, err := time.Parse(time.DateOnly, loupanJson.OpenDate)
	if err != nil {
		openDate = time.Time{}
	}

	priceConfirmTime, err := time.Parse(time.DateTime, loupanJson.PriceConfirmTime)
	if err != nil {
		priceConfirmTime = time.Time{}
	}
	showPriceConfirmTime, _ := strconv.ParseInt(strings.Replace(loupanJson.ShowPriceConfirmTime,
		"天", "", -1), 10, 64)
	referenceAvgPrice, _ := strconv.ParseFloat(loupanJson.ReferenceAvgPrice, 64)
	hasCarActivity, _ := strconv.Atoi(loupanJson.HasCarActivity)
	isNewSale, _ := strconv.Atoi(loupanJson.IsNewSale)
	var projectTags []string
	for _, v := range loupanJson.ProjectTags {
		projectTags = append(projectTags, v.Desc)
	}
	properRight, _ := strconv.Atoi(strings.Replace(loupanJson.Properright, "年", "", -1))
	var convergedRooms []byte
	if len(loupanJson.ConvergedRooms) > 0 {
		convergedRooms, _ = json.Marshal(loupanJson.ConvergedRooms)
	}
	var preloadDetailImage []byte
	if len(loupanJson.PreloadDetailImage) > 0 {
		preloadDetailImage, _ = json.Marshal(loupanJson.PreloadDetailImage)
	}
	var newSaleTags []byte
	if len(newSaleTags) > 0 {
		newSaleTags, _ = json.Marshal(loupanJson.NewSaleTags)
	}
	referenceTotalPriceRange, _ := json.Marshal(loupanJson.ReferenceTotalPriceRange)
	var firstTags []byte
	if len(loupanJson.FirstTags) > 0 {
		firstTags, _ = json.Marshal(loupanJson.FirstTags)
	}
	var mFirstTags []byte
	if len(loupanJson.MFirstTags) > 0 {
		mFirstTags, _ = json.Marshal(loupanJson.MFirstTags)
	}
	var recommendReasonStr string
	if recommendReason, ok := loupanJson.RecommendReason.(string); ok && recommendReason != "null" {
		recommendReasonStr = recommendReason
	}
	downloadInfo, _ := json.Marshal(loupanJson.DownloadInfo)
	var specialTags []byte
	if len(loupanJson.SpecialTags) > 0 {
		specialTags, _ = json.Marshal(loupanJson.SpecialTags)
	}
	var lianjiaSpecial []byte
	if len(loupanJson.Special.LianjiaSpecial) > 0 {
		lianjiaSpecial, _ = json.Marshal(loupanJson.Special.LianjiaSpecial)
	}
	var developerSpecial []byte
	if len(loupanJson.Special.DeveloperSpecial) > 0 {
		developerSpecial, _ = json.Marshal(loupanJson.Special.DeveloperSpecial)
	}
	var frameRooms []byte
	if len(loupanJson.FrameRooms) > 0 {
		frameRooms, _ = json.Marshal(loupanJson.FrameRooms)
	}
	loupanInfo := LoupanInfo{
		Pid:                      int32(pid),
		ID:                       int64(id),
		CityID:                   cityID,
		CityName:                 loupanJson.CityName,
		CoverPic:                 loupanJson.CoverPic,
		MinFrameArea:             int32(minFrameArea),
		MaxFrameArea:             int32(maxFrameArea),
		DistrictName:             loupanJson.DistrictName,
		District:                 loupanJson.District,
		DistrictID:               int32(districtID),
		BizcircleID:              int32(bizcircleID),
		BizcircleName:            loupanJson.BizcircleName,
		BuildID:                  buildID,
		PermitAllReady:           permitAllReady,
		ProcessStatus:            processStatus,
		ResblockFrameArea:        loupanJson.ResblockFrameArea,
		ResblockFrameAreaRange:   loupanJson.ResblockFrameAreaRange,
		ResblockFrameAreaDesc:    loupanJson.ResblockFrameAreaDesc,
		Decoration:               loupanJson.Decoration,
		Longitude:                longitude,
		Latitude:                 latitude,
		FrameRoomsDesc:           loupanJson.FrameRoomsDesc,
		Title:                    loupanJson.Title,
		ResblockName:             loupanJson.ResblockName,
		ResblockAlias:            loupanJson.ResblockAlias,
		Address:                  loupanJson.Address,
		StoreAddr:                loupanJson.StoreAddr,
		AvgUnitPrice:             avgUnitPrice,
		AveragePrice:             averagePrice,
		AddressRemark:            loupanJson.AddressRemark,
		ProjectName:              loupanJson.ProjectName,
		SpecialTags:              string(specialTags),
		LianjiaSpecial:           string(lianjiaSpecial),
		LianjiaSpecialComm:       loupanJson.Special.LianjiaSpecialComm,
		DeveloperSpecial:         string(developerSpecial),
		DeveloperSpecialType:     loupanJson.Special.DeveloperSpecialType,
		DeveloperSpecialComm:     loupanJson.Special.DeveloperSpecialComm,
		FrameRooms:               string(frameRooms),
		ConvergedRooms:           string(convergedRooms),
		Tags:                     strings.Join(loupanJson.Tags, ","),
		ProjectTags:              strings.Join(projectTags, ","),
		HouseType:                loupanJson.HouseType,
		HouseTypeValue:           loupanJson.HouseTypeValue,
		SaleStatus:               loupanJson.SaleStatus,
		HasEvaluate:              hasEvaluate,
		HasVrHouse:               hasVrHouse,
		HasShortVideo:            hasShortVideo,
		OpenDate:                 openDate,
		HasVirtualView:           hasVirtualView,
		LowestTotalPrice:         int32(lowestTotalPrice),
		PriceShowConfig:          int32(priceShowConfig),
		ShowPrice:                showPrice,
		ShowPriceUnit:            loupanJson.ShowPriceUnit,
		ShowPriceDesc:            loupanJson.ShowPriceDesc,
		ShowPriceConfirmTime:     showPriceConfirmTime,
		PriceConfirmTime:         priceConfirmTime,
		Status:                   status,
		SubwayDistance:           loupanJson.SubwayDistance,
		IsCooperation:            isCooperation,
		EvaluateStatus:           evaluateStatus,
		ShowPriceInfo:            loupanJson.ShowPriceInfo,
		BrandID:                  loupanJson.BrandID,
		PreloadDetailImage:       string(preloadDetailImage),
		ReferenceAvgPrice:        referenceAvgPrice,
		ReferenceAvgPriceUnit:    loupanJson.ReferenceAvgPriceUnit,
		ReferenceAvgPriceDesc:    loupanJson.ReferenceAvgPriceDesc,
		ReferenceTotalPrice:      loupanJson.ReferenceTotalPrice,
		ReferenceTotalPriceUnit:  loupanJson.ReferenceTotalPriceUnit,
		ReferenceTotalPriceDesc:  loupanJson.ReferenceTotalPriceDesc,
		NewSaleTags:              string(newSaleTags),
		IsSupportOnlineSale:      isSupportOnlineSale,
		IsUndertake:              isUndertake,
		Properright:              properRight,
		DeveloperCompany:         strings.Join(loupanJson.DeveloperCompany, ","),
		PropertyCompany:          strings.Join(loupanJson.PropertyCompany, ","),
		LiveTag:                  liveTag,
		Prod:                     prod,
		ReferenceTotalPriceRange: string(referenceTotalPriceRange),
		HouseNode:                houseNode,
		TotalPriceStart:          totalPriceStart,
		TotalPriceStartUnit:      loupanJson.TotalPriceStartUnit,
		AvgPriceStart:            avgPriceStart,
		AvgPriceStartUnit:        loupanJson.AvgPriceStartUnit,
		OnTime:                   onTime,
		ProjectDesc:              loupanJson.ProjectDesc,
		HasCarActivity:           hasCarActivity,
		IsNewSale:                isNewSale,
		FirstTags:                string(firstTags),
		MFirstTags:               string(mFirstTags),
		FbExpoID:                 loupanJson.FbExpoID,
		StrategyInfo:             loupanJson.StrategyInfo,
		RecommendLogInfo:         loupanJson.RecommendLogInfo,
		RecommendReason:          recommendReasonStr,
		ReferenceTotalPriceTips:  loupanJson.ReferenceTotalPriceTips,
		AppDetailURL:             loupanJson.AppDetailURL,
		FilterDesc:               loupanJson.FilterDesc,
		URL:                      loupanJson.URL,
		DownloadInfo:             string(downloadInfo),
	}

	return &loupanInfo, nil
}

type LoupanInfo struct {
	Pid                      int32     `json:"pid"`
	ID                       int64     `json:"id"`
	CityID                   int       `json:"city_id"`
	CityName                 string    `json:"city_name"`
	CoverPic                 string    `json:"cover_pic"`
	MinFrameArea             int32     `json:"min_frame_area"`
	MaxFrameArea             int32     `json:"max_frame_area"`
	DistrictName             string    `json:"district_name"`
	District                 string    `json:"district"`
	DistrictID               int32     `json:"district_id"`
	BizcircleID              int32     `json:"bizcircle_id"`
	BizcircleName            string    `json:"bizcircle_name"`
	BuildID                  int64     `json:"build_id"`
	PermitAllReady           int       `json:"permit_all_ready"`
	ProcessStatus            int       `json:"process_status"`
	ResblockFrameArea        string    `json:"resblock_frame_area"`
	ResblockFrameAreaRange   string    `json:"resblock_frame_area_range"`
	ResblockFrameAreaDesc    string    `json:"resblock_frame_area_desc"`
	Decoration               string    `json:"decoration"`
	Longitude                float64   `json:"longitude"`
	Latitude                 float64   `json:"latitude"`
	FrameRoomsDesc           string    `json:"frame_rooms_desc"`
	Title                    string    `json:"title"`
	ResblockName             string    `json:"resblock_name"`
	ResblockAlias            string    `json:"resblock_alias"`
	Address                  string    `json:"address"`
	StoreAddr                string    `json:"store_addr"`
	AvgUnitPrice             float64   `json:"avg_unit_price"`
	AveragePrice             float64   `json:"average_price"`
	AddressRemark            string    `json:"address_remark"`
	ProjectName              string    `json:"project_name"`
	SpecialTags              string    `json:"special_tags"`
	LianjiaSpecial           string    `json:"lianjia_special"`
	LianjiaSpecialComm       string    `json:"lianjia_special_comm"`
	DeveloperSpecial         string    `json:"developer_special"`
	DeveloperSpecialType     string    `json:"developer_special_type"`
	DeveloperSpecialComm     string    `json:"developer_special_comm"`
	FrameRooms               string    `json:"frame_rooms"`
	ConvergedRooms           string    `json:"converged_rooms"`
	Tags                     string    `json:"tags"`
	ProjectTags              string    `json:"project_tags"`
	HouseType                string    `json:"house_type"`
	HouseTypeValue           string    `json:"house_type_value"`
	SaleStatus               string    `json:"sale_status"`
	HasEvaluate              int       `json:"has_evaluate"`
	HasVrHouse               int       `json:"has_vr_house"`
	HasShortVideo            int       `json:"has_short_video"`
	OpenDate                 time.Time `json:"open_date"`
	HasVirtualView           int       `json:"has_virtual_view"`
	LowestTotalPrice         int32     `json:"lowest_total_price"`
	PriceShowConfig          int32     `json:"price_show_config"`
	ShowPrice                int       `json:"show_price"`
	ShowPriceUnit            string    `json:"show_price_unit"`
	ShowPriceDesc            string    `json:"show_price_desc"`
	ShowPriceConfirmTime     int64     `json:"show_price_confirm_time"`
	PriceConfirmTime         time.Time `json:"price_confirm_time"`
	Status                   int       `json:"status"`
	SubwayDistance           string    `json:"subway_distance"`
	IsCooperation            int       `json:"is_cooperation"`
	EvaluateStatus           int       `json:"evaluate_status"`
	ShowPriceInfo            string    `json:"show_price_info"`
	BrandID                  string    `json:"brand_id"`
	PreloadDetailImage       string    `json:"preload_detail_image"`
	ReferenceAvgPrice        float64   `json:"reference_avg_price"`
	ReferenceAvgPriceUnit    string    `json:"reference_avg_price_unit"`
	ReferenceAvgPriceDesc    string    `json:"reference_avg_price_desc"`
	ReferenceTotalPrice      string    `json:"reference_total_price"`
	ReferenceTotalPriceUnit  string    `json:"reference_total_price_unit"`
	ReferenceTotalPriceDesc  string    `json:"reference_total_price_desc"`
	NewSaleTags              string    `json:"new_sale_tags"`
	IsFollowed               int       `json:"is_followed"`
	IsSupportOnlineSale      int       `json:"is_support_online_sale"`
	IsUndertake              int       `json:"is_undertake"`
	Properright              int       `json:"properright"`
	DeveloperCompany         string    `json:"developer_company"`
	PropertyCompany          string    `json:"property_company"`
	LiveTag                  int       `json:"live_tag"`
	Prod                     int       `json:"prod"`
	ReferenceTotalPriceRange string    `json:"reference_total_price_range"`
	HouseNode                int       `json:"house_node"`
	TotalPriceStart          int       `json:"total_price_start"`
	TotalPriceStartUnit      string    `json:"total_price_start_unit"`
	AvgPriceStart            int       `json:"avg_price_start"`
	AvgPriceStartUnit        string    `json:"avg_price_start_unit"`
	OnTime                   time.Time `json:"on_time"`
	ProjectDesc              string    `json:"project_desc"`
	HasCarActivity           int       `json:"has_car_activity"`
	IsNewSale                int       `json:"is_new_sale"`
	FirstTags                string    `json:"first_tags"`
	MFirstTags               string    `json:"m_first_tags"`
	FbExpoID                 string    `json:"fb_expo_id"`
	StrategyInfo             string    `json:"strategy_info"`
	RecommendLogInfo         string    `json:"recommend_log_info"`
	RecommendReason          string    `json:"recommend_reason"`
	ReferenceTotalPriceTips  string    `json:"reference_total_price_tips"`
	AppDetailURL             string    `json:"app_detail_url"`
	FilterDesc               string    `json:"filter_desc"`
	URL                      string    `json:"url"`
	DownloadInfo             string    `json:"download_info"`
}

func (uc *LianjiaUsecase) ListCityLoupan(ctx context.Context, city *CityInfo) error {
	defer func(uc *LianjiaUsecase, ctx context.Context, city *CityInfo) {
		err := uc.finishCurrentTask(ctx, HouseTypeLoupan, city)
		if err != nil {
			uc.log.Errorf("finishCurrentLoupanTask err! city: %s err: %v", city.Name, err)
		}
	}(uc, ctx, city)
	var client *resty.Client
	client = resty.New()
	var headers http.Header
	headers = http.Header{}
	headers.Add("User-Agent", RandomUA())
	client.Header = headers
	client.SetRetryCount(3)
	client.SetRetryMaxWaitTime(5 * time.Second)
	url := fmt.Sprintf(loupanListUrl, city.Url, 1)
	resp, err := client.R().Get(url)
	if err != nil {
		uc.log.Errorf("ListCityLoupan err! url: %s err: %v", url, err)
	}
	if resp.StatusCode() == http.StatusNotFound {
		uc.log.Warnf("ListCityLoupan err! city %s haven't loupan! url: %s", city.Name, url)
		return nil
	}
	parser := gjson.Get(resp.String(), loupanRootPath)
	if parser.String() == "null" || parser.String() == "" {
		uc.log.Warnf("ListCityLoupan err! city %s haven't loupan! url: %s", city.Name, url)
		return nil
	}
	total := parser.Get(loupanTotalPageCountPath).String()
	pages, err := strconv.ParseFloat(total, 64)
	totalPage := int(pages)
	if err != nil {
		uc.log.Errorf("parse page err! %v", err)
		return err
	}
	if pages == 0 {
		return errors.New(fmt.Sprintf("city %s empty page", city.Name))
	}
	for page := 1; page <= totalPage; page++ {
		headers = http.Header{}
		headers.Add("User-Agent", RandomUA())
		client.Header = headers
		url := fmt.Sprintf(loupanListUrl, city.Url, page)
		resp, err := client.R().Get(url)
		if err != nil {
			uc.log.Errorf("ListLoupan err! url: %s err: %v", url, err)
			continue
		}
		if strings.Contains(resp.String(), rateLimitKeyword) {
			loupanSleepTime = 5 * time.Second
			uc.log.Errorf("rate limiting, start sleep %s", loupanSleepTime)
			time.Sleep(loupanSleepTime)
			continue
		}
		parser := gjson.Get(resp.String(), loupanRootPath)
		if parser.String() == "null" || parser.String() == "" {
			uc.log.Warnf("ListCityLoupan err! city %s haven't loupan! url: %s", city.Name, url)
			continue
		}
		houseList := parser.Get(loupanListPath).Array()
		if len(houseList) > 0 {
			var messages []Message
			for _, house := range houseList {
				hashValue := uc.md5(house.String())
				key := LianjiaLoupanIdempotentKey + hashValue
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
			err := uc.Send(ctx, messages, HouseTypeLoupan)
			if err != nil {
				uc.log.Errorf("send message err! err: %v", err)
			}
		}
		time.Sleep(loupanSleepTime)
	}

	return nil
}

func (uc *LianjiaUsecase) HandleLoupanMessage(ctx context.Context) error {
	var slices []*LoupanInfo
	for {
		select {
		case msgs := <-LoupanProcessChan:
			if len(msgs) > 0 {
				for _, msg := range msgs {
					var raw LoupanJson
					err := json.Unmarshal(msg.Content, &raw)
					if err != nil {
						uc.log.Errorf("parse loupan json err! %v", err)
						continue
					}
					loupanInfo, _ := raw.ConvertToLoupanInfo()
					exist, err := uc.repo.CheckLoupanExists(ctx, int64(loupanInfo.CityID), loupanInfo.ID)
					if !exist {
						slices = append(slices, loupanInfo)
					}
				}
				if len(slices) > 0 || (len(slices) >= insertBatchSize || uc.IsLoupanTaskFinish(ctx)) {
					err := uc.repo.InsertLoupanInfo(ctx, slices)
					if err != nil {
						uc.log.Errorf("batch insert loupan err! %v", err)
					}
					slices = make([]*LoupanInfo, 0)
				}
			}
		}
	}
}

func (uc *LianjiaUsecase) IsLoupanTaskFinish(ctx context.Context) bool {
	unFinishTaskStr, err := uc.repo.GetKey(ctx, LianjiaLoupanRunningCrawlTaskKey)
	if err != nil && unFinishTaskStr != "" {
		uc.log.Errorf("get loupan unfinish task key err! %v", err)
		return false
	}
	if unFinishTaskStr == "" {
		return true
	}
	unFinishTasks, err := strconv.ParseInt(unFinishTaskStr, 10, 64)
	if err != nil {
		uc.log.Errorf("parse loupan task num err! %v", err)
		return false
	}
	if unFinishTasks > 0 {
		return false
	}

	return true
}
