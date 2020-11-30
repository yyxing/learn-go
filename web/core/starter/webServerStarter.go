package starter

import (
	"github.com/spf13/viper"
	"learn-go/web/core"
	"learn-go/web/core/context"
)

type WebServerStarter struct {
	AbstractStarter
}

func (starter *WebServerStarter) Init(context context.ApplicationContext) {
	config, ok := context.Get(GlobalConfigKey).(viper.Viper)
	if !ok {
		panic("config load failed")
	}
	var port string
	port = ":" + config.GetString(core.ServerPortKey)
	if len(port) < 1 {
		port = ":8080"
	}
	context.Set(core.ServerPortKey, port)
	//routes := core.IrisApplication.GetRoutes()
	//for _, route := range routes {
	//	logrus.Info(route)
	//}
}

func (starter *WebServerStarter) GetOrder() int {
	return core.Int32Max
}
