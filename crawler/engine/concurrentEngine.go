package engine

import (
	"awesomeProject/crawler/fetcher"
	"awesomeProject/crawler/types"
	"fmt"
	"log"
)

type Scheduler interface {
	NotifyReady
	Submit(types.Request)
	GetWorkerChan() chan types.Request
	Run()
}

type NotifyReady interface {
	WorkerReady(chan types.Request)
}
type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
}

//存在卡死问题是因为由于只开了指定个通道去等待请求接受，
//而通道的特性是请求和返回接收是一个通道，也就是说在请求不断发起的时候，必须有空闲的worker不能参与请求的接受，而是再等待请求的返回
//否则由于通道等待阻塞的特性导致第一批的worker完成接收任务时，没有worker处理返回从而导致阻塞，而发起请求需要等返回处理后才有新的任务进来
func (engine *ConcurrentEngine) Run(seeds ...types.Request) {
	engine.Scheduler.Run()
	out := make(chan types.ParseResult)
	for i := 0; i < engine.WorkCount; i++ {
		createWorker(engine.Scheduler.GetWorkerChan(), out, engine.Scheduler)
	}
	for _, task := range seeds {
		engine.Scheduler.Submit(task)
	}
	result := types.ParseResult{}
	itemCount := 1
	for {
		result = <-out
		for _, item := range result.Items {
			log.Printf("Got #%d item %v", itemCount, item)
			itemCount++
		}
		for _, r := range result.Requests {
			engine.Scheduler.Submit(r)
		}
	}
}

// 创建协程执行器 将耗时的请求由协程单独完成不占用主执行引擎的时间，使主引擎无需等待
func createWorker(in chan types.Request,
	out chan types.ParseResult, ready NotifyReady) {
	go func() {
		for {
			ready.WorkerReady(in)
			// worker接受请求
			request := <-in
			// 执行fetcher 获取当前页面的数据
			bytes, err := fetcher.Fetcher(request)
			if err != nil {
				fmt.Println(err)
				continue
			}
			// 执行解析器获取解析结果以及从当前页面获取到的任务
			parseResult := request.ParseFunc(bytes)
			// 请求结果通知回channel
			out <- parseResult
		}
	}()
}
