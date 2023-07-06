package comical

import (
	"net/http"
)

// HandlerFunc defines the request handler used by comical
type HandlerFunc func(c *Context)

// Engine implements the interface of http.Handler
type Engine struct {
	// router map: key is "METHOD PATTERN", value is HandlerFunc
	router Router
}

// New is the constructor of comical.Engine
func New(router Router) *Engine {
	if router == nil {
		router = NewDefaultRouter()
	}
	return &Engine{
		router: router,
	}
}

// ServeHTTP makes the Engine implement the interface of http.Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c := newContext(w, r)
	e.router.Handle(c)
}

// Implements RESTFUL API

// addRoute adds a new DefaultRouter to the Engine
func (e *Engine) addRoute(method, pattern string, handler HandlerFunc) {
	e.router.AddRoute(method, pattern, handler)
}

// GET adds a new defaultRouter with GET method to the Engine
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST adds a new defaultRouter with POST method to the Engine
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRoute("POST", pattern, handler)
}

// PUT adds a new defaultRouter with PUT method to the Engine
func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.addRoute("PUT", pattern, handler)
}

// DELETE adds a new defaultRouter with DELETE method to the Engine
func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.addRoute("DELETE", pattern, handler)
}

// OPTIONS adds a new defaultRouter with OPTIONS method to the Engine
func (e *Engine) OPTIONS(pattern string, handler HandlerFunc) {
	e.addRoute("OPTIONS", pattern, handler)
}

// HEAD adds a new defaultRouter with HEAD method to the Engine
func (e *Engine) HEAD(pattern string, handler HandlerFunc) {
	e.addRoute("HEAD", pattern, handler)
}

// TODO: implement more HTTP methods

// Run defines the method to start a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
