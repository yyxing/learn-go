package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisRecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"time"
)

func InitIris() *iris.Application {
	app := iris.New()
	app.Use(irisRecover.New())
	cfg := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Query:   true,
		Columns: true,
		LogFunc: func(endTime time.Time, latency time.Duration, status, ip,
			method, path string, message interface{}, headerMessage interface{}) {
			logrus.Infof("| %s | %s | %s | %s | %s |",
				latency.String(), status, ip, method,
				path)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
