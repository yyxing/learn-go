package main

import (
	"learn-go/web/app"
	"learn-go/web/core/boot"
)

func main() {
	application := boot.Default()
	iris := app.InitIris()
	application.RunIrisServer(iris)
	defer application.Stop()
}
