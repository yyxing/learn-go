package parser

import (
	"awesomeProject/crawler/types"
	"net/http"
	"regexp"
)

var cityListCompile = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)"[^>]*>([^<]+)</a>`)

func ParseCityList(content []byte) types.ParseResult {
	// 正则获取所有城市信息
	all := cityListCompile.FindAllSubmatch(content, -1)
	cityListResult := types.ParseResult{}
	for _, match := range all {
		// 填充城市名称
		cityListResult.Items = append(cityListResult.Items, string(match[2]))
		// 拼接下一个对于城市的请求以及解析函数
		cityListResult.Requests = append(cityListResult.Requests, types.Request{
			Url:    string(match[1]),
			Method: http.MethodGet,
			Body:   nil,
			// 单个城市解析用户的url
			ParseFunc: ParseCityUser,
		})
	}
	return cityListResult
}
