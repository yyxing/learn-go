package main

import (
	"learn-go/web/core/boot"
)

func main() {
	application := boot.Default()
	application.Run()
	defer application.Stop()
}
