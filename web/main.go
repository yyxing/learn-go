package main

import (
	"learn-go/web/infra"
)

func main() {
	application := infra.Default()
	application.Run()
	defer application.Stop()
}
