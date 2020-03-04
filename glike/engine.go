package glike

import (
	"fmt"
	"net/http"
	"sync"
)

type Engine struct {
	// that's weird
	// engine-->RouteGroup-->engine
	RouteGroup

	// use string map not tree for simplity
	routesGet map[string]HandlersChain
	routesPost map[string]HandlersChain

	pool   sync.Pool
}

func NewEngine() *Engine {
	e := &Engine{
		RouteGroup: RouteGroup{
			Base:           "",
		},
		routesGet:make(map[string]HandlersChain),
		routesPost:make(map[string]HandlersChain),
	}
	// that's weird
	e.RouteGroup.engine = e

	e.pool.New = func() interface{} {
		return e.allocateContext()
	}
	return e
}

func (e *Engine) Group(base string) *RouteGroup{
	rg := &RouteGroup{
		Base:           base,
		engine:         e,
	}
	return rg
}

func (e *Engine) allocateContext() *Context {
	return &Context{}
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	//fmt.Println("new request")

	ctx := e.pool.Get().(*Context)
	ctx.reset()
	ctx.request = r
	ctx.responseWriter = w

	e.handle(ctx)
	e.pool.Put(ctx)
}

func (e *Engine) handle(ctx *Context) {
	path := ctx.request.URL.Path
	method := ctx.request.Method

	// simple implement
	var handlers map[string]HandlersChain
	switch method{
	case "GET":
		handlers = e.routesGet
	case "POST":
		handlers = e.routesPost
	default:
		panic("unknown method")
	}
	handlersChain ,ok := handlers[path]
	if !ok{
		ctx.responseWriter.WriteHeader(404)
		return
	}
	ctx.handlers = handlersChain

	ctx.Next()
}

func (e *Engine) Run(addr string) {
	// http listen and serve
	if addr == "" {
		addr = ":8080"
	}

	fmt.Println("running on",addr)
	// recover...
	http.ListenAndServe(addr,e)
}
