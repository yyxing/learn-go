package scheduler

import (
	"awesomeProject/crawler/types"
)

// 并发版Schedule使用goroutine去创建请求时已经可以达到要求，但是每一个request请求开一个协程导致所有的协程无法管理
// 这个调度器将请求与处理的协程统一管理 由调度器去将worker与请求匹配 可以在不开多个协程的情况下 实现并发的请求处理且不会阻塞
type QueueScheduler struct {
	requestChan chan types.Request
	workerChan  chan chan types.Request
}

func (s *QueueScheduler) Submit(c types.Request) {
	s.requestChan <- c
}

func (s *QueueScheduler) WorkerReady(c chan types.Request) {
	s.workerChan <- c
}
func (s *QueueScheduler) GetWorkerChan() chan types.Request {
	return make(chan types.Request)
}
func (s *QueueScheduler) Run() {
	// 初始化调度器的接收通道
	s.workerChan = make(chan chan types.Request)
	s.requestChan = make(chan types.Request)
	go func() {
		// 声明队列
		var requestQueue []types.Request
		var workerQueue []chan types.Request
		for {
			var activeRequest types.Request
			var activeWorker chan types.Request
			if len(requestQueue) > 0 &&
				len(workerQueue) > 0 {
				activeRequest = requestQueue[0]
				activeWorker = workerQueue[0]
			}
			// 使用select对通道进行多路复用，哪个通道有接收就对哪个进行处理
			select {
			case r := <-s.requestChan:
				// 收到worker解析完的新请求 将其加到队列中
				requestQueue = append(requestQueue, r)
			case w := <-s.workerChan:
				// 收到worker完成的请求 加到空闲worker队列中
				workerQueue = append(workerQueue, w)
			case activeWorker <- activeRequest:
				// 若activeWorker和activeRequest不为nil  则表示可以进行数据的处理 将请求发给worker处理
				requestQueue = requestQueue[1:]
				workerQueue = workerQueue[1:]
			}
		}
	}()
}
