package main

import (
	"github.com/gin-gonic/gin"
	"github.com/julienschmidt/httprouter"
	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	routing "github.com/qiangxue/fasthttp-routing"
)

func RegisterHandlers() *httprouter.Router {
	router := httprouter.New()
	router.POST("/user", CreateUser)
	router.POST("/user/login", Login)
	router.GET("/user/:username", QueryUserInfo)
	//router.GET("/test/", TestQps)
	return router
}
func Route() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	//router.POST("/user", CreateUser)
	//router.POST("/user/login", Login)
	//router.GET("/user/:username", QueryUserInfo)
	router.GET("/test/", TestQps)
	return router
}

func IrisRoute() *iris.Application {
	app := iris.Default()
	app.Logger().Level = golog.WarnLevel
	app.Get("/test", TestQpsIris)
	return app
}

func FastRoute() *routing.Router {
	router := routing.New()
	router.Get("/test", DoJSONWrite)
	return router
}
func main() {
	//err := http.ListenAndServe(":8090", RegisterHandlers())
	//if err != nil {
	//	panic(err)
	//}
	//route := Route()
	//err := route.Run(":8090")
	//if err != nil {
	//	panic(err)
	//}

	//route := IrisRoute()
	//err := route.Run(iris.Addr(":8090"))
	//if err != nil {
	//	panic(err)
	//}
	//route := FastRoute()
	//err := fasthttp.ListenAndServe(":8090", route.HandleRequest)
	//if err != nil {
	//	panic(err)
	//}
}
