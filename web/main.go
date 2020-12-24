package main

import (
	"learn-go/web/apis"
	"learn-go/web/app"
	"learn-go/web/core/boot"
)

func main() {
	application := boot.Default()
	iris := app.InitIris()
	apis.Routes(iris)
	application.RunIrisServer(iris)

	defer application.Stop()
}
