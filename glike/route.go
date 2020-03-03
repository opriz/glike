package glike

type RouteGroup struct {
	Base string
	engine *Engine
}

func (r *RouteGroup) GET(relativePath string, handler func(ctx *Context)) {
	// check validation of relativePath...

	r.engine.routesGet[r.Base+relativePath] = handler
}

func (r *RouteGroup) POST(relativePath string, handler func(ctx *Context)) {

	r.engine.routesPost[r.Base+relativePath] = handler
}




