package base

import "learn-go/web/core"

// 类似Java的Abstract类
type AbstractStarter struct{}

func (starter *AbstractStarter) Init(context core.ApplicationContext) {
}

func (starter *AbstractStarter) Start(context core.ApplicationContext) {
}

func (starter *AbstractStarter) Finalize(context core.ApplicationContext) {
}

func (starter *AbstractStarter) GetOrder() int {
	return 0
}
