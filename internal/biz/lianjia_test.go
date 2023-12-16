package biz

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "{\"house_code\":2100181420,\"building_id\":10000043,\"housedel_code\":106103501924,\"title\":\"环球中心 1号线锦城广场地铁站  带租约房源 1号线地铁 锦城广场地铁口\",\"city_id\":510100,\"city_name\":\"成都市\",\"district_name\":\"高新\",\"bizcircle_name\":\"金融城\",\"street_name\":\"天府大道北段1700号\",\"resblock_name\":\"环球中心\",\"unit_rent_price\":0,\"unit_month_rent_price\":0,\"rent_price\":0,\"unit_sell_price\":12191,\"sell_price\":2100000,\"area\":172.26,\"score\":165.6,\"cmark_name\":null,\"image\":\"//image1.ljcdn.com/crep/product/image/1569571310261-201909271602008290.jpg.210x160.jpg\",\"fitment\":2,\"fitment_name\":\"精装修\",\"has_furniture\":true,\"has_furniture_name\":\"不带家具\",\"has_meeting_room\":true,\"has_meeting_room_name\":\"不带会议室\",\"parking_type\":3,\"is_near_subway\":true,\"is_near_subway_name\":\"地铁十分钟\",\"is_registered_company\":true,\"is_registered_company_name\":\"可注册\",\"floor_position\":2,\"floor_position_name\":\"中层\",\"showing_time\":1,\"showing_time_name\":\"提前预约随时可看\",\"subway_distance\":727,\"subway_distance_name\":\"距离1号线(科学城-韦家碾);1号线(五根松-韦家碾)锦城广场站727米\",\"tags\":[\"带家具\",\"会议室\",\"可注册\",\"近地铁\",\"精装修\"],\"is_real\":false,\"has_vr\":true,\"ctime\":\"2019-09-25 11:55:00\"}"
	hash := md5.New()
	hash.Write([]byte(str))
	h := hex.EncodeToString(hash.Sum(nil))
	fmt.Println(h)
}
