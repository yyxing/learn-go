package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/logger"
	irisRecover "github.com/kataras/iris/v12/middleware/recover"
	"github.com/sirupsen/logrus"
	"time"
)

func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(204)
		return
	}
	ctx.Next()
}
func InitIris() *iris.Application {
	app := iris.New()
	app.Use(irisRecover.New())
	app.Use(Cors)
	cfg := logger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Query:   true,
		Columns: true,
		LogFunc: func(endTime time.Time, latency time.Duration, status, ip,
			method, path string, message interface{}, headerMessage interface{}) {
			logrus.Debugf("| %s | %s | %s | %s | %s |",
				latency.String(), status, ip, method,
				path)
		},
	}
	app.Use(logger.New(cfg))
	return app
}
