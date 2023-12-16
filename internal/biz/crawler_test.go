package biz

import (
	"encoding/json"
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func TestHttpGet(t *testing.T) {
	client := resty.New()
	headers := http.Header{}
	headers.Add("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36")
	resp, err := client.R().
		EnableTrace().
		Get(cityListUrl)
	if err != nil {
		fmt.Println(err)
	}
	cityRegex := "\"cityList\":(.*?)\\}}"
	regex := regexp.MustCompile(cityRegex)
	results := regex.FindAllString(resp.String(), -1)
	cityStr := results[0]
	cityStr = strings.TrimLeft(cityStr, "\"cityList\":")
	cityStr = strings.Replace(cityStr, "id", "city_code", -1)
	cityMap := make(CityMap, 0)
	err = json.Unmarshal([]byte(cityStr), &cityMap)
	if err != nil {
		return
	}
	fmt.Println(cityMap)
}
