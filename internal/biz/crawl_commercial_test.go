package biz

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"math"
	"net/http"
	"testing"
	"time"
)

func TestCrawlCommercial(t *testing.T) {
	var client *resty.Client
	client = resty.New()
	var headers http.Header
	cityId := 110000
	page := 1
	size := 50
	headers = http.Header{}
	commUrl, err := NewCommercialUrl(cityId, CommercialHouseTypeBuyXzl, page, size)
	client = resty.New()
	headers = http.Header{}
	headers.Add("User-Agent", RandomUA())
	client.Header = headers
	client.SetRetryCount(3)
	client.SetRetryMaxWaitTime(5 * time.Second)
	resp, err := client.R().Get(commUrl)
	if err != nil {
		fmt.Println(err)
	}
	parser := gjson.Get(resp.String(), commercialDataRootPath)
	if parser.String() == "null" || parser.String() == "" {
		fmt.Println("ListCityCommercial city  haven't ershoufang! url: ")
	}
	totalCount := parser.Get(commercialDataTotalPath).Num
	results := parser.Get(commercialDataDocsPath).Array()
	pages := int(math.Ceil(totalCount / 20))
	fmt.Println(pages)
	fmt.Println(results)
}
