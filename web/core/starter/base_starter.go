package starter

import (
	"learn-go/web/core/context"
)

// 类似Java的Abstract类
type AbstractStarter struct {
}

// 业务异常
type BusinessError interface {
	error
}

func (starter *AbstractStarter) Init(context context.ApplicationContext) {
}

func (starter *AbstractStarter) Start(context context.ApplicationContext) {
}

func (starter *AbstractStarter) Finalize(context context.ApplicationContext) {
}

func (starter *AbstractStarter) GetOrder() int {
	return 0
}
