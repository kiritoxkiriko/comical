package comical

import "log"

type RouteGroup struct {
	// router extend from router
	*router
	prefix      string
	parent      *RouteGroup
	middlewares []HandlerFunc
}

func newRootGroup() *RouteGroup {
	return &RouteGroup{
		prefix: "",
		router: newRouter(),
	}
}

func (g *RouteGroup) Group(prefix string) *RouteGroup {
	prefix = prefix + g.prefix
	return &RouteGroup{
		prefix: prefix,
		parent: g,
		router: g.router,
	}
}

// Use apply middleware into some group
func (g *RouteGroup) Use(middlewares ...HandlerFunc) {
	g.middlewares = append(g.middlewares, middlewares...)
}

// addRoute override method of router
func (g *RouteGroup) addRoute(method, pattern string, handler HandlerFunc) {
	pattern = g.prefix + pattern
	log.Printf("Route %4s - %s\n", method, pattern)
	// call router's addRouter
	g.router.addRoute(method, pattern, handler)
}

// GET adds a new defaultRouter with GET method to the Engine
func (g *RouteGroup) GET(pattern string, handler HandlerFunc) {
	g.addRoute("GET", pattern, handler)
}

// POST adds a new defaultRouter with POST method to the Engine
func (g *RouteGroup) POST(pattern string, handler HandlerFunc) {
	g.addRoute("POST", pattern, handler)
}

// PUT adds a new defaultRouter with PUT method to the Engine
func (g *RouteGroup) PUT(pattern string, handler HandlerFunc) {
	g.addRoute("PUT", pattern, handler)
}

// DELETE adds a new defaultRouter with DELETE method to the Engine
func (g *RouteGroup) DELETE(pattern string, handler HandlerFunc) {
	g.addRoute("DELETE", pattern, handler)
}

// OPTIONS adds a new defaultRouter with OPTIONS method to the Engine
func (g *RouteGroup) OPTIONS(pattern string, handler HandlerFunc) {
	g.addRoute("OPTIONS", pattern, handler)
}

// HEAD adds a new defaultRouter with HEAD method to the Engine
func (g *RouteGroup) HEAD(pattern string, handler HandlerFunc) {
	g.addRoute("HEAD", pattern, handler)
}

// TODO: implement more HTTP methods
