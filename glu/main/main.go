package main

import (
	"github.com/yyxing/glu"
	"github.com/yyxing/glu/context"
)

func main() {
	engine := glu.New()
	engine.Get("/test", func(ctx *context.Context) {
		_, _ = ctx.WriteString("Hello World")
	})
	engine.Run(":8000")
}
