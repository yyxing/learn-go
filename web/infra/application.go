package infra

import (
	"learn-go/web/core"
	"learn-go/web/core/base"
)

type Application struct {
	context core.ApplicationContext
}

// TODO 由用户指定starter启动 暂未想好如何实现
func New() {

}

// 默认配置启动 config log sql等
func Default() Application {
	application := Application{context: core.ApplicationContext{}}
	application.context.Register(&base.ConfigStarter{})
	application.context.SortStarter()
	return application
}

func (application *Application) Run() {
	application.init()
	application.start()
}

// 初始化starter
func (application *Application) init() {
	for _, starter := range application.context.GetAllStarters() {
		// 调用每个starter的Init方法
		starter.Init(application.context)
	}
}

// 启动所有starter
func (application *Application) start() {
	for _, starter := range application.context.GetAllStarters() {
		// 调用每个starter的start方法
		starter.Start(application.context)
	}
}
func (application *Application) Stop() {
	// 停止所有starter
	for _, starter := range application.context.GetAllStarters() {
		starter.Finalize(application.context)
	}
}
