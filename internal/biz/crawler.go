package biz

import (
	"context"
	"encoding/json"
	"github.com/go-resty/resty/v2"
	"github.com/minicloudsky/lianjia/pkg/time_track"
	"github.com/panjf2000/ants"
	"net/http"
	"regexp"
	"strings"
	"sync"
	"time"
)

const (
	cityListUrl       = "https://m.lianjia.com/city/"
	ershoufangListUrl = "https://m.lianjia.com/liverpool/api/ershoufang/getList?cityId=%d&curPage=%d"
	loupanListUrl     = "%sloupan/pg%d/?_t=1&source=index"
	zufangUrl         = "https://m.lianjia.com/chuzu/%s/zufang/"
	cityListRegex     = "\"cityList\":(.*?)\\}}"
	goroutineSize     = 40
)

type poolParam struct {
	ctx       context.Context
	city      *CityInfo
	houseType HoueseType
}

func (uc *LianjiaUsecase) FetchCityList(ctx context.Context) {
	defer time_track.Track(uc.log, "FetchCityList", time.Now())
	client := resty.New()
	headers := http.Header{}
	headers.Add("User-Agent", RandomUA())
	resp, err := client.R().
		EnableTrace().
		Get(cityListUrl)
	if err != nil {
		uc.log.Errorf("get city err! %v", err)
	}
	regex := regexp.MustCompile(cityListRegex)
	results := regex.FindAllString(resp.String(), -1)
	if len(results) == 0 {
		return
	}
	cityStr := results[0]
	cityStr = strings.TrimLeft(cityStr, "\"cityList\":")
	cityStr = strings.Replace(cityStr, "id", "city_code", -1)
	cityMap := make(CityMap)
	err = json.Unmarshal([]byte(cityStr), &cityMap)
	if err != nil {
		return
	}
	var cityList []*CityInfo
	for _, city := range cityMap {
		cityList = append(cityList, &CityInfo{
			CityCode:      city.CityCode,
			Name:          city.Name,
			Host:          city.Host,
			MobileHost:    city.MobileHost,
			Short:         city.Short,
			Province:      city.Province,
			ProvinceShort: city.ProvinceShort,
			Pinyin:        city.Pinyin,
			Longitude:     city.Longitude,
			Latitude:      city.Latitude,
			Url:           city.Url,
		})
	}
	uc.log.Infof("total City: %d", len(cityList))
	err = uc.repo.UpsertCity(ctx, cityList)
	if err != nil {
		uc.log.Errorf("FetchCityList err! err: %v", err)
		return
	}
}

func (uc *LianjiaUsecase) FetchErShouFang(ctx context.Context) {
	defer time_track.Track(uc.log, "FetchErShouFang", time.Now())
	go func() {
		err := uc.HandleErshoufangMessage(ctx)
		if err != nil {
			uc.log.Errorf("HandleErshoufangMessage err! %v", err)
		}
	}()
	cities, err := uc.repo.GetAllCity(ctx)
	err = uc.repo.SetNXKey(ctx, LianjiaErshoufangRunningCrawlTaskKey, len(cities), 24*time.Hour)
	if err != nil {
		uc.log.Errorf("ershoufang crawl task running, exit.")
		return
	}
	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(goroutineSize, func(i interface{}) {
		uc.poolFunc(i)
		wg.Done()
	})
	if err != nil {
		uc.log.Errorf("ants.NewPoolWithFunc err! err: %v", err)
	}
	for _, city := range cities {
		wg.Add(1)
		param := poolParam{
			ctx:       ctx,
			city:      city,
			houseType: HouseTypeErshoufang,
		}
		err = p.Invoke(param)
		if err != nil {
			uc.log.Errorf("invoke ershoufang task err! err: %v", err)
		}
	}
	wg.Wait()
}

func (uc *LianjiaUsecase) FetchLoupan(ctx context.Context) {
	defer time_track.Track(uc.log, "FetchLoupan", time.Now())
	go func() {
		err := uc.HandleLoupanMessage(ctx)
		if err != nil {
			uc.log.Errorf("HandleLoupanMessage err! %v", err)
		}
	}()
	cities, err := uc.repo.GetAllCity(ctx)
	err = uc.repo.SetNXKey(ctx, LianjiaLoupanRunningCrawlTaskKey, len(cities), 24*time.Hour)
	if err != nil {
		uc.log.Errorf("loupan crawl task running, exit.")
		return
	}
	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(10, func(i interface{}) {
		uc.poolFunc(i)
		wg.Done()
	})
	if err != nil {
		uc.log.Errorf("ants.NewPoolWithFunc err! err: %v", err)
	}
	for _, city := range cities {
		wg.Add(1)
		param := poolParam{
			ctx:       ctx,
			city:      city,
			houseType: HouseTypeLoupan,
		}
		err = p.Invoke(param)
		if err != nil {
			uc.log.Errorf("invoke loupan task err! err: %v", err)
		}
	}
	wg.Wait()
}

func (uc *LianjiaUsecase) FetchZufang(ctx context.Context) {
	defer time_track.Track(uc.log, "FetchZufang", time.Now())
}

func (uc *LianjiaUsecase) FetchCommercial(ctx context.Context) {
	defer time_track.Track(uc.log, "FetchCommercial", time.Now())
	go func() {
		err := uc.HandleCommercialMessage(ctx)
		if err != nil {
			uc.log.Errorf("HandleCommercialMessage err! %v", err)
		}
	}()
	cities, err := uc.repo.GetAllCity(ctx)
	err = uc.repo.SetNXKey(ctx, LianjiaCommercialRunningCrawlTaskKey, len(cities)*4, 24*time.Hour)
	if err != nil {
		uc.log.Errorf("commercial crawl task running, exit.")
		return
	}
	var wg sync.WaitGroup
	p, err := ants.NewPoolWithFunc(goroutineSize, func(i interface{}) {
		uc.poolFunc(i)
		wg.Done()
	})
	if err != nil {
		uc.log.Errorf("ants.NewPoolWithFunc err! err: %v", err)
	}
	for _, city := range cities {
		wg.Add(1)
		param := poolParam{
			ctx:       ctx,
			city:      city,
			houseType: HouseTypeCommercial,
		}
		err = p.Invoke(param)
		if err != nil {
			uc.log.Errorf("invoke commercial task err! err: %v", err)
		}
	}
	wg.Wait()
}

func (uc *LianjiaUsecase) poolFunc(i interface{}) {
	param := i.(poolParam)
	ctx := param.ctx
	city := param.city
	houseType := param.houseType
	switch houseType {
	case HouseTypeErshoufang:
		err := uc.ListCityErshouFang(ctx, city)
		if err != nil {
			return
		}
	case HouseTypeLoupan:
		err := uc.ListCityLoupan(ctx, city)
		if err != nil {
			return
		}
	case HouseTypeCommercial:
		err := uc.ListCityCommercial(ctx, city)
		if err != nil {
			return
		}
	case HouseTypeZufang:
		err := uc.ListCityZufang(ctx, city)
		if err != nil {
			return
		}
	}
}
