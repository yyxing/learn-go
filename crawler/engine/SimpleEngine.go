package engine

import (
	"awesomeProject/crawler/fetcher"
	"awesomeProject/crawler/types"
	"fmt"
	"log"
)

// 具体控制流程的类 由他调取fetcher获取数据 然后再交给parser去解析信息 完成调度

// 根据根url开启爬虫引擎
func Run(seeds ...types.Request) {
	var tasks []types.Request
	for _, task := range seeds {
		tasks = append(tasks, task)
	}
	for len(tasks) > 0 {
		t := tasks[0]
		tasks = tasks[1:]
		// 执行fetcher 获取当前页面的数据
		bytes, err := fetcher.Fetcher(t)
		if err != nil {
			fmt.Println(err)
			continue
		}
		// 执行解析器获取解析结果以及从当前页面获取到的任务
		parseResult := t.ParseFunc(bytes)
		// 将请求结果放入任务队列中
		tasks = append(tasks, parseResult.Requests...)
		// 目前是打印获取信息
		for _, item := range parseResult.Items {
			log.Printf("Got item %s", item)
		}
	}
}
