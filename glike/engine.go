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
	routesGet map[string]func(ctx *Context)
	routesPost map[string]func(ctx *Context)

	pool   sync.Pool
}

func NewEngine() *Engine {
	e := &Engine{
		RouteGroup: RouteGroup{
			Base:           "",
		},
		routesGet:make(map[string]func(ctx *Context)),
		routesPost:make(map[string]func(ctx *Context)),
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
	ctx := e.pool.Get().(*Context)
	ctx.request = r
	ctx.responseWriter = w

	e.handle(ctx)
	e.pool.Put(ctx)
}

func (e *Engine) handle(ctx *Context) {
	path := ctx.request.URL.Path
	method := ctx.request.Method

	fmt.Printf("method: %v, path: %v\n",method,path)

	// simple implement
	var handlers map[string]func(*Context)
	switch method{
	case "GET":
		handlers = e.routesGet
	case "POST":
		handlers = e.routesPost
	default:
		panic("unknown method")
	}
	handler ,ok := handlers[path]
	if !ok{
		ctx.responseWriter.WriteHeader(404)
		return
	}

	handler(ctx)
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
