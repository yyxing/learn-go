package hello

import "github.com/kataras/iris/v12"

func TestHello(context iris.Context) {
	context.WriteString("Hello World")
}
