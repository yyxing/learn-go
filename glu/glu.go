package glu

import (
	"github.com/yyxing/glu/context"
	"github.com/yyxing/glu/router"
	"log"
	"net/http"
)

type Engine struct {
	router *router.Router
}

func New() *Engine {
	return &Engine{router: router.NewRouter()}
}

func (e *Engine) addRoute(method string, pattern string, handler context.Handler) {
	e.router.AddRouter(method, pattern, handler)
}

func (e *Engine) Post(pattern string, handler context.Handler) {
	e.addRoute(http.MethodPost, pattern, handler)
}

func (e *Engine) Get(pattern string, handler context.Handler) {
	e.addRoute(http.MethodGet, pattern, handler)
}

func (e *Engine) Run(addr string) {
	log.Printf("Now listening on: http://localhost%s\n", addr)
	log.Fatal(http.ListenAndServe(addr, e))
}
func (e *Engine) ServeHTTP(w http.ResponseWriter, request *http.Request) {
	e.router.Serve(context.NewContext(w, request))
}
