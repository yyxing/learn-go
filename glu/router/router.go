package router

import (
	"fmt"
	"github.com/yyxing/glu/context"
	"net/http"
)

var (
	separator = "-"
)

type Router struct {
	handlers map[string]context.Handler
}

func (router Router) AddRouter(method string, pattern string, handler context.Handler) {
	key := method + separator + pattern
	router.handlers[key] = handler
}

func (router *Router) Handle(method string, pattern string, handler context.Handler) {
	key := method + separator + pattern
	router.handlers[key] = handler
}

func (router *Router) Serve(ctx *context.Context) {
	key := ctx.Method + separator + ctx.Path
	if handler, ok := router.handlers[key]; ok {
		handler(ctx)
	} else {
		ctx.StatusCode(http.StatusNotFound)
		_, _ = ctx.WriteString(fmt.Sprintf("404 NOT FOUND: %s\n", ctx.Path))
	}
}
func NewRouter() *Router {
	return &Router{handlers: make(map[string]context.Handler)}
}
