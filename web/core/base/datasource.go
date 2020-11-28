package base

import (
	"github.com/spf13/viper"
	"learn-go/web/core"
)

type DatasourceStarter struct {
	AbstractStarter
}

func (starter *DatasourceStarter) Init(context core.ApplicationContext) {

}

// mysql连接配置
// 正式建立连接
func (starter *DatasourceStarter) Start(context core.ApplicationContext) {
	config, ok := context.Get(GlobalConfigKey)
	if !ok {
		panic("config load failed")
	}
	config = config.(viper.Viper)

}

// 关闭db
func (starter *DatasourceStarter) Finalize(context core.ApplicationContext) {

}

func (starter *DatasourceStarter) GetOrder() int {
	return -1
}
