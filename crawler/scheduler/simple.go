package scheduler

import "awesomeProject/crawler/types"

type SimpleScheduler struct {
	workChan chan types.Request
}

func (s *SimpleScheduler) Submit(request types.Request) {
	// 将请求提交给通道
	go func() { s.workChan <- request }()
}

func (s *SimpleScheduler) WorkerReady(c chan types.Request) {

}

func (s *SimpleScheduler) GetWorkerChan() chan types.Request {
	return s.workChan
}
func (s *SimpleScheduler) Run() {
	s.workChan = make(chan types.Request)
}
