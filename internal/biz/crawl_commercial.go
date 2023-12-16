package biz

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/google/go-querystring/query"
	"github.com/tidwall/gjson"
	"math"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CommercialHouseType string

const (
	CommercialHouseTypeBuyXzl  = "buy_xzl"
	CommercialHouseTypeRentXzl = "rent_xzl"
	CommercialHouseTypeBuySp   = "buy_sp"
	CommercialHouseTypeRentSp  = "rent_sp"
	CommercialBaseUrl          = "https://shang.lianjia.com/api/ke/"
	commercialDataRootPath     = "data"
	commercialDataTotalPath    = "total"
	commercialDataDocsPath     = "docs"
	commercialPageSize         = 50
)

var CommercialHouseTypes = []CommercialHouseType{
	CommercialHouseTypeBuyXzl,
	CommercialHouseTypeRentXzl,
	CommercialHouseTypeBuySp,
	CommercialHouseTypeRentSp,
}

type ListCommercialParams struct {
	CityID       int    `url:"city_id"`
	PageType     string `url:"pageType"`
	OriginType   string `url:"originType"`
	BusinessType int    `url:"business_type"`
	Platform     int    `url:"platform"`
	Device       int    `url:"device"`
	Page         int    `url:"page"`
	Size         int    `url:"size"`
	FromPage     string `url:"from_page"`
}

func NewCommercialUrl(cityId int, commType CommercialHouseType, page, size int) (string, error) {
	baseUrl := ""
	var params ListCommercialParams
	switch commType {
	case CommercialHouseTypeBuyXzl:
		baseUrl = CommercialBaseUrl + "xzl/list"
		params = ListCommercialParams{
			CityID:       cityId,
			PageType:     "buy",
			OriginType:   "xzl",
			BusinessType: 1,
			Platform:     2,
			Device:       2,
			Page:         page,
			Size:         size,
			FromPage:     "index",
		}
	case CommercialHouseTypeRentXzl:
		baseUrl = CommercialBaseUrl + "xzl/list"
		params = ListCommercialParams{
			CityID:       cityId,
			PageType:     "rent",
			OriginType:   "xzl",
			BusinessType: 2,
			Platform:     2,
			Device:       2,
			Page:         page,
			Size:         size,
			FromPage:     "index",
		}
	case CommercialHouseTypeBuySp:
		baseUrl = CommercialBaseUrl + "sp/list"
		params = ListCommercialParams{
			CityID:       cityId,
			PageType:     "buy",
			OriginType:   "sp",
			BusinessType: 3,
			Platform:     2,
			Device:       2,
			Page:         page,
			Size:         size,
			FromPage:     "index",
		}
	case CommercialHouseTypeRentSp:
		baseUrl = CommercialBaseUrl + "sp/list"
		params = ListCommercialParams{
			CityID:       cityId,
			PageType:     "rent",
			OriginType:   "sp",
			BusinessType: 4,
			Platform:     2,
			Device:       2,
			Page:         page,
			Size:         size,
			FromPage:     "index",
		}
	default:
		return "", errors.New("invalid commercial type")
	}
	queryString, _ := query.Values(params)
	fullUrl := baseUrl + "?" + queryString.Encode()
	return fullUrl, nil
}

type CommercialRaw struct {
	HouseCode               int         `json:"house_code"`
	BuildingID              int         `json:"building_id"`
	HousedelCode            int64       `json:"housedel_code"`
	Title                   string      `json:"title"`
	CityID                  int         `json:"city_id"`
	CityName                string      `json:"city_name"`
	DistrictName            string      `json:"district_name"`
	BizcircleName           string      `json:"bizcircle_name"`
	StreetName              string      `json:"street_name"`
	ResblockName            string      `json:"resblock_name"`
	UnitRentPrice           float64     `json:"unit_rent_price"`
	UnitMonthRentPrice      float64     `json:"unit_month_rent_price"`
	RentPrice               int         `json:"rent_price"`
	UnitSellPrice           int         `json:"unit_sell_price"`
	SellPrice               int         `json:"sell_price"`
	Area                    float64     `json:"area"`
	Score                   float64     `json:"score"`
	CmarkName               interface{} `json:"cmark_name"`
	Image                   string      `json:"image"`
	Fitment                 int         `json:"fitment"`
	FitmentName             string      `json:"fitment_name"`
	HasFurniture            bool        `json:"has_furniture"`
	HasFurnitureName        string      `json:"has_furniture_name"`
	HasMeetingRoom          bool        `json:"has_meeting_room"`
	HasMeetingRoomName      string      `json:"has_meeting_room_name"`
	ParkingType             int         `json:"parking_type"`
	IsNearSubway            bool        `json:"is_near_subway"`
	IsNearSubwayName        string      `json:"is_near_subway_name"`
	IsRegisteredCompany     bool        `json:"is_registered_company"`
	IsRegisteredCompanyName string      `json:"is_registered_company_name"`
	FloorPosition           int         `json:"floor_position"`
	FloorPositionName       string      `json:"floor_position_name"`
	ShowingTime             int         `json:"showing_time"`
	ShowingTimeName         string      `json:"showing_time_name"`
	SubwayDistance          int         `json:"subway_distance"`
	SubwayDistanceName      string      `json:"subway_distance_name"`
	Tags                    []string    `json:"tags"`
	IsReal                  bool        `json:"is_real"`
	HasVr                   bool        `json:"has_vr"`
	Ctime                   string      `json:"ctime"`
}

func (c *CommercialRaw) ConvertToCommercial() CommercialInfo {
	var hasFurniture int
	if c.HasFurniture {
		hasFurniture = 1
	}
	var isNearSubway int
	if c.IsNearSubway {
		isNearSubway = 1
	}
	var hasMettingRoom int
	if c.HasMeetingRoom {
		hasMettingRoom = 1
	}
	var isRegisteredCompany int
	if c.IsRegisteredCompany {
		isRegisteredCompany = 1
	}
	var tags string
	if len(c.Tags) > 0 {
		tags = strings.Join(c.Tags, ",")
	}
	var isReal int
	if c.IsReal {
		isReal = 1
	}
	var hasVr int
	if c.HasVr {
		hasVr = 1
	}
	var ctime time.Time
	if c.Ctime != "" {
		ctime, _ = time.Parse(time.DateTime, c.Ctime)
	}
	var cmakeName string
	if val, ok := c.CmarkName.(string); ok && cmakeName != "nil" {
		cmakeName = val
	}
	return CommercialInfo{
		HouseCode:               c.HouseCode,
		BuildingID:              c.BuildingID,
		HousedelCode:            c.HousedelCode,
		Title:                   c.Title,
		CityID:                  c.CityID,
		CityName:                c.CityName,
		DistrictName:            c.DistrictName,
		BizcircleName:           c.BizcircleName,
		StreetName:              c.StreetName,
		ResblockName:            c.ResblockName,
		UnitRentPrice:           c.UnitRentPrice,
		UnitMonthRentPrice:      c.UnitMonthRentPrice,
		RentPrice:               c.RentPrice,
		UnitSellPrice:           c.UnitSellPrice,
		SellPrice:               c.SellPrice,
		Area:                    c.Area,
		Score:                   c.Score,
		CmarkName:               cmakeName,
		Image:                   c.Image,
		Fitment:                 c.Fitment,
		FitmentName:             c.FitmentName,
		HasFurniture:            hasFurniture,
		HasFurnitureName:        c.HasFurnitureName,
		HasMeetingRoom:          hasMettingRoom,
		HasMeetingRoomName:      c.HasMeetingRoomName,
		ParkingType:             c.ParkingType,
		IsNearSubway:            isNearSubway,
		IsNearSubwayName:        c.IsNearSubwayName,
		IsRegisteredCompany:     isRegisteredCompany,
		IsRegisteredCompanyName: c.IsRegisteredCompanyName,
		FloorPosition:           c.FloorPosition,
		FloorPositionName:       c.FloorPositionName,
		ShowingTime:             c.ShowingTime,
		ShowingTimeName:         c.ShowingTimeName,
		SubwayDistance:          c.SubwayDistance,
		SubwayDistanceName:      c.SubwayDistanceName,
		Tags:                    tags,
		IsReal:                  isReal,
		HasVr:                   hasVr,
		Ctime:                   ctime,
	}
}

type CommercialInfo struct {
	HouseCode               int       `json:"house_code"`
	BuildingID              int       `json:"building_id"`
	HousedelCode            int64     `json:"housedel_code"`
	Title                   string    `json:"title"`
	CityID                  int       `json:"city_id"`
	CityName                string    `json:"city_name"`
	DistrictName            string    `json:"district_name"`
	BizcircleName           string    `json:"bizcircle_name"`
	StreetName              string    `json:"street_name"`
	ResblockName            string    `json:"resblock_name"`
	UnitRentPrice           float64   `json:"unit_rent_price"`
	UnitMonthRentPrice      float64   `json:"unit_month_rent_price"`
	RentPrice               int       `json:"rent_price"`
	UnitSellPrice           int       `json:"unit_sell_price"`
	SellPrice               int       `json:"sell_price"`
	Area                    float64   `json:"area"`
	Score                   float64   `json:"score"`
	CmarkName               string    `json:"cmark_name"`
	Image                   string    `json:"image"`
	Fitment                 int       `json:"fitment"`
	FitmentName             string    `json:"fitment_name"`
	HasFurniture            int       `json:"has_furniture"`
	HasFurnitureName        string    `json:"has_furniture_name"`
	HasMeetingRoom          int       `json:"has_meeting_room"`
	HasMeetingRoomName      string    `json:"has_meeting_room_name"`
	ParkingType             int       `json:"parking_type"`
	IsNearSubway            int       `json:"is_near_subway"`
	IsNearSubwayName        string    `json:"is_near_subway_name"`
	IsRegisteredCompany     int       `json:"is_registered_company"`
	IsRegisteredCompanyName string    `json:"is_registered_company_name"`
	FloorPosition           int       `json:"floor_position"`
	FloorPositionName       string    `json:"floor_position_name"`
	ShowingTime             int       `json:"showing_time"`
	ShowingTimeName         string    `json:"showing_time_name"`
	SubwayDistance          int       `json:"subway_distance"`
	SubwayDistanceName      string    `json:"subway_distance_name"`
	Tags                    string    `json:"tags"`
	IsReal                  int       `json:"is_real"`
	HasVr                   int       `json:"has_vr"`
	Ctime                   time.Time `json:"ctime"`
}

func (uc *LianjiaUsecase) HandleCommercialMessage(ctx context.Context) error {
	var slices []*CommercialInfo
	for {
		select {
		case msgs := <-CommercialProcessChan:
			for _, msg := range msgs {
				if len(msg.Content) > 0 {
					var raw *CommercialRaw
					err := json.Unmarshal(msg.Content, &raw)
					if err != nil {
						uc.log.Errorf("Unmarshal commercial info err! commercial: %s err: %v",
							string(msg.Content), err)
						continue
					}
					commercialInfo := raw.ConvertToCommercial()
					exist, err := uc.repo.CheckCommercialExists(ctx, int64(commercialInfo.CityID), commercialInfo.HousedelCode)
					if err != nil {
						uc.log.Errorf("check commercial exists err! %v", err)
						continue
					}
					if !exist {
						slices = append(slices, &commercialInfo)
					}
				}
			}
			if len(slices) > 0 && (len(slices) >= insertBatchSize || uc.IsCommercialTaskFinish(ctx)) {
				err := uc.repo.InsertCommercialInfo(ctx, slices)
				if err != nil {
					uc.log.Errorf("insert commercial info err! %v", err)
				}
				slices = make([]*CommercialInfo, 0)
			}
		}
	}
}

func (uc *LianjiaUsecase) ListCityCommercial(ctx context.Context, city *CityInfo) error {
	for _, commType := range CommercialHouseTypes {
		go func(commType CommercialHouseType) {
			defer func(uc *LianjiaUsecase, ctx context.Context, city *CityInfo) {
				err := uc.finishCurrentTask(ctx, HouseTypeCommercial, city)
				if err != nil {
					uc.log.Errorf("finish current commercial task err! %v", err)
				}
			}(uc, ctx, city)
			commUrl, err := NewCommercialUrl(city.CityCode, commType, 0, commercialPageSize)
			if err != nil {
				uc.log.Errorf("get commercial url err! %v", err)
				return
			}
			var client *resty.Client
			client = resty.New()
			var headers http.Header
			headers = http.Header{}
			headers.Add("User-Agent", RandomUA())
			client.Header = headers
			client.SetRetryCount(3)
			client.SetRetryMaxWaitTime(7 * time.Second)
			resp, err := client.R().Get(commUrl)
			if err != nil {
				uc.log.Errorf("ListErshoufang err! url: %s err: %v", commUrl, err)
			}
			parser := gjson.Get(resp.String(), commercialDataRootPath)
			if parser.String() == "null" || parser.String() == "" {
				uc.log.Warnf("ListCityCommercial city %s haven't ershoufang! url: %s", city.Name, commUrl)
				return
			}
			totalCount := parser.Get(commercialDataTotalPath).Num
			CommercialResults := parser.Get(commercialDataDocsPath).Array()
			pages := int(math.Ceil(totalCount / commercialPageSize))
			if pages > 0 || len(CommercialResults) > 0 {
				messages := make([]Message, 0)
				for _, comm := range CommercialResults {
					hashValue := uc.md5(comm.String())
					key := LianjiaCommercialIdempotentKey + hashValue
					if val, err := uc.repo.GetKey(ctx, key); err == nil && val != "" {
						continue
					} else {
						messages = append(messages, Message{
							Content: []byte(comm.String()),
						})
						if err := uc.SetNXKey(ctx, key, time.Now().Unix(), 6*time.Hour); err != nil {
							continue
						}
					}
				}
				err := uc.Send(ctx, messages, HouseTypeCommercial)
				if err != nil {
					uc.log.Errorf("fail to send commercial msgs %v", err)
					return
				}
				for page := 1; page <= pages; page++ {
					commUrl, err := NewCommercialUrl(city.CityCode, commType, page, commercialPageSize)
					if err != nil {
						uc.log.Errorf("get commercial url err! %v", err)
						continue
					}
					var client *resty.Client
					client = resty.New()
					var headers http.Header
					headers = http.Header{}
					headers.Add("User-Agent", RandomUA())
					client.Header = headers
					client.SetRetryCount(3)
					client.SetRetryMaxWaitTime(7 * time.Second)
					resp, err := client.R().Get(commUrl)
					if err != nil {
						uc.log.Errorf("ListErshoufang err! url: %s err: %v", commUrl, err)
						continue
					}
					parser := gjson.Get(resp.String(), commercialDataRootPath)
					if parser.String() == "null" || parser.String() == "" {
						uc.log.Warnf("ListCityCommercial city %s haven't ershoufang! url: %s", city.Name, commUrl)
						continue
					}
					CommercialResults := parser.Get(commercialDataDocsPath).Array()
					messages := make([]Message, 0)
					for _, comm := range CommercialResults {
						hashValue := uc.md5(comm.String())
						key := LianjiaCommercialIdempotentKey + hashValue
						if val, err := uc.repo.GetKey(ctx, key); err == nil && val != "" {
							continue
						} else {
							messages = append(messages, Message{
								Content: []byte(comm.String()),
							})
							if err := uc.SetNXKey(ctx, key, time.Now().Unix(), 6*time.Hour); err != nil {
								continue
							}
						}
					}
					err = uc.Send(ctx, messages, HouseTypeCommercial)
					if err != nil {
						uc.log.Errorf("fail to send commercial msgs %v", err)
						continue
					}
				}
			}
		}(commType)
	}
	return nil
}

func (uc *LianjiaUsecase) IsCommercialTaskFinish(ctx context.Context) bool {
	unFinishTaskStr, err := uc.repo.GetKey(ctx, LianjiaCommercialRunningCrawlTaskKey)
	if err != nil && unFinishTaskStr != "" {
		uc.log.Errorf("get commercial unfinish task key err! %v", err)
		return false
	}
	if unFinishTaskStr == "" {
		return true
	}
	unFinishTasks, err := strconv.ParseInt(unFinishTaskStr, 10, 64)
	if err != nil {
		uc.log.Errorf("parse commercial task num err! %v", err)
		return false
	}
	if unFinishTasks > 0 {
		return false
	}

	return true
}
