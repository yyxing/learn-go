package main

import (
	"learn-go/web/apis"
	"learn-go/web/app"
	"learn-go/web/core/boot"
	"learn-go/web/tasks"
)

func main() {
	application := boot.Default()
	iris := app.InitIris()
	apis.Routes(iris)
	tasks.RegisterTasks()
	application.RunIrisServer(iris)
	defer application.Stop()
}
