package comical

import (
	"net/http"
)

// HandlerFunc defines the request handler used by comical
type HandlerFunc func(c *Context)

// Engine implements the interface of http.Handler
type Engine struct {
	// router map: key is "METHOD PATTERN", value is HandlerFunc
	router *router
}

func New() *Engine {
	return &Engine{
		router: newRouter(),
	}
}

// ServeHTTP makes the Engine implement the interface of http.Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.handle(c)
}

// Implements RESTFUL API

// addRoute adds a new router to the Engine
func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.addRoute(method, pattern, handler)
}

// GET adds a new router with GET method to the Engine
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST adds a new router with POST method to the Engine
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// PUT adds a new router with PUT method to the Engine
func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.addRoute("PUT", pattern, handler)
}

// DELETE adds a new router with DELETE method to the Engine
func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.addRoute("DELETE", pattern, handler)
}

// OPTIONS adds a new router with OPTIONS method to the Engine
func (e *Engine) OPTIONS(pattern string, handler HandlerFunc) {
	e.addRoute("OPTIONS", pattern, handler)
}

// HEAD adds a new router with HEAD method to the Engine
func (e *Engine) HEAD(pattern string, handler HandlerFunc) {
	e.addRoute("HEAD", pattern, handler)
}

// TODO: implement more HTTP methods

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
