package starter

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"learn-go/web/core"
	"learn-go/web/core/context"
)

type LogStarter struct {
	AbstractStarter
}

// 配置Log 后期增加日志相关的配置 和滚动更新
func (starter LogStarter) Init(context context.ApplicationContext) {
	formatter := prefixed.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03.00000",
	}
	log.SetFormatter(&formatter)
}

func (starter LogStarter) GetOrder() int {
	return core.Int32Min + 2
}
