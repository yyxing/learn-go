package starter

import (
	"learn-go/web/core/context"
	"time"
)

type TaskStarter struct {
	AbstractStarter
	tickers []time.Ticker
}

func (starter *TaskStarter) Init(context context.ApplicationContext) {
	starter.tickers = make([]time.Ticker, 10)
}

func (starter *TaskStarter) Start(context context.ApplicationContext) {
}

func (starter *TaskStarter) Finalize(context context.ApplicationContext) {
}

func (starter *TaskStarter) GetOrder() int {
	return 0
}
