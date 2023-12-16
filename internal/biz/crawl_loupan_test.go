package biz

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"github.com/tidwall/gjson"
	"net/http"
	"strings"
	"testing"
)

func TestTrimMobileHost(t *testing.T) {
	host := "m.lianjia.com/you/hnld/"
	fmt.Println(strings.TrimRight(host, "/"))
}

func TestGetLoupan(t *testing.T) {
	host := "https://m.lianjia.com/quanzhou/"
	client := resty.New()
	headers := http.Header{}
	headers.Add("User-Agent", RandomUA())
	url := fmt.Sprintf(loupanListUrl, host, 1)
	resp, err := client.R().
		EnableTrace().
		Get(url)
	if err != nil {
		fmt.Println(err)
	}
	parser := gjson.Get(resp.String(), loupanRootPath)
	if parser.String() == "null" {
		fmt.Printf("ListCityLoupan err! city %s haven't loupan!", "xx")
	}
	total := parser.Get(loupanTotalPageCountPath).Num
	list := parser.Get(loupanListPath).Array()
	fmt.Println(total)
	fmt.Println(list)
}
