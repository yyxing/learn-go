package parser

import (
	"awesomeProject/crawler/types"
	"net/http"
	"regexp"
)

var cityUserCompile = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[0-9]+)"[^>]*>([^<]+)</a>`)

// 根据城市的url 获取当前城市目前可见的用户url
func ParseCityUser(content []byte) types.ParseResult {
	all := cityUserCompile.FindAllSubmatch(content, -1)

	cityUserResult := types.ParseResult{}
	for _, match := range all {
		name := string(match[2])
		cityUserResult.Items = append(cityUserResult.Items, name)
		cityUserResult.Requests = append(cityUserResult.Requests, types.Request{
			Url:    string(match[1]),
			Method: http.MethodGet,
			Body:   nil,
			// 这个解析器获取用户url的请求，解析函数使用用户的解析函数 解析用户信息
			ParseFunc: ParseUser,
		})
	}
	return cityUserResult
}
