package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/types"
	"awesomeProject/crawler/za/parser"
	"net/http"
)

const rootUrl = "http://www.zhenai.com/zhenghun"

func main() {
	engine.Run(types.Request{
		Url:       rootUrl,
		Method:    http.MethodGet,
		Body:      nil,
		ParseFunc: parser.ParseCityList,
	})
}
