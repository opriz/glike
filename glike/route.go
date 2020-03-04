package glike

type HandleFunc func(*Context)

type HandlersChain []HandleFunc

func (hc HandlersChain) Handle(ctx *Context) {
	ctx.Next()
}

type RouteGroup struct {
	Base string
	engine *Engine

	// order maters
	// hold group-wide middlewares
	groupHandlers []HandleFunc
}

func (r *RouteGroup) GET(relativePath string, handlers ...HandleFunc) {
	// check validation of relativePath...

	mergedHandlers := r.combineHandlers(handlers...)
	r.engine.routesGet[r.Base+relativePath] = mergedHandlers
}

func (r *RouteGroup) POST(relativePath string, handlers ...HandleFunc) {

	mergedHandlers := r.combineHandlers(handlers...)
	r.engine.routesPost[r.Base+relativePath] = mergedHandlers
}

func (r *RouteGroup) combineHandlers(handlers ...HandleFunc) []HandleFunc{
	// could optimize allocation options
	mergedHandlers := append(r.groupHandlers,handlers...)
	copyOfMergedHandlers := make([]HandleFunc,len(mergedHandlers))
	copy(copyOfMergedHandlers,mergedHandlers)
	return copyOfMergedHandlers
}

func (r *RouteGroup) Use(funcs ...HandleFunc) {
	r.groupHandlers = append(r.groupHandlers,funcs...)
}
