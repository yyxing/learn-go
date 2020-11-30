package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisRecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"time"
)

func initIris() *iris.Application {
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
			logrus.Infof("| %s | %s | %s | %s | %s | %s |",
				endTime.Format("2006-01-02 15-04-05.000000"), latency.String(), status, ip, method,
				path)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
