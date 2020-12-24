package starter

import (
	log "github.com/sirupsen/logrus"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"learn-go/web/core"
	"learn-go/web/core/context"
	"strings"
)

type LogStarter struct {
	AbstractStarter
}

// 配置Log 后期增加日志相关的配置 和滚动更新
func (starter LogStarter) Init(context context.ApplicationContext) {
	config := GetConfig()
	logLevel := config.GetString("logger.level")
	logLevel = strings.ToLower(logLevel)
	switch logLevel {
	case "info":
		log.SetLevel(log.InfoLevel)
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	case "warn":
		log.SetLevel(log.WarnLevel)
	case "trace":
		log.SetLevel(log.TraceLevel)
	case "fatal":
		log.SetLevel(log.FatalLevel)
	case "panic":
		log.SetLevel(log.PanicLevel)
	}
	formatter := prefixed.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		ForceFormatting: true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:03.00000",
	}
	formatter.SetColorScheme(&prefixed.ColorScheme{
		TimestampStyle: "37",
	})
	log.SetFormatter(&formatter)
}

func (starter LogStarter) GetOrder() int {
	return core.Int32Min + 1
}
