package parser

import (
	"io/ioutil"
	"testing"
)

func TestParseCity(t *testing.T) {
	expectRequestsLen := 470
	expectCitiesLen := 470
	// 表格驱动测试
	expectRequestUrls := []string{
		"http://www.zhenai.com/zhenghun/aba",
		"http://www.zhenai.com/zhenghun/akesu",
		"http://www.zhenai.com/zhenghun/alashanmeng",
	}
	expectRequestCities := []string{
		"阿坝",
		"阿克苏",
		"阿拉善盟",
	}
	file, err := ioutil.ReadFile("city_list_test_data.html")
	if err != nil {
		panic(err)
	}
	result := ParseCityList(file)
	if len(result.Requests) != expectRequestsLen {
		t.Errorf("expect requestLen %d, but %d", expectRequestsLen, len(result.Requests))
	}
	if len(result.Items) != expectCitiesLen {
		t.Errorf("expect citiesLen %d, but %d", expectCitiesLen, len(result.Items))
	}

	for i, url := range expectRequestUrls {
		if url != result.Requests[i].Url {
			t.Errorf("expect url %s, but %s", url, result.Requests[i].Url)
		}
	}

	for i, city := range expectRequestCities {
		if city != result.Items[i] {
			t.Errorf("expect city %s, but %s", city, result.Items[i])
		}
	}
}
