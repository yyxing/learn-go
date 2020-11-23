package main

import (
	"awesomeProject/crawler/engine"
	"awesomeProject/crawler/scheduler"
	"awesomeProject/crawler/types"
	"awesomeProject/crawler/za/parser"
	"net/http"
)

const rootUrl = "http://www.zhenai.com/zhenghun"

func main() {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueueScheduler{},
		WorkCount: 100,
	}
	e.Run(types.Request{
		Url:       rootUrl,
		Method:    http.MethodGet,
		Body:      nil,
		ParseFunc: parser.ParseCityList,
	})
}
