package comical

import (
	"fmt"
	"net/http"
)

// HandlerFunc alias http.HandlerFunc, defines the request handler used by comical
type HandlerFunc http.HandlerFunc

// Engine implements the interface of http.Handler
type Engine struct {
	router map[string]HandlerFunc
}

func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

// ServeHTTP makes the Engine implement the interface of http.Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + " " + r.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, r)
	} else {
		// 404 NOT FOUND
		_, err := fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
		if err != nil {
			panic(fmt.Errorf("error writing response: %v", err))
		}
	}
}

// Implements RESTFUL API

// addRouter adds a new router to the Engine
func (e *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + " " + pattern
	// this will overwrite the old handler if the key already exists
	e.router[key] = handler
}

// GET adds a new router with GET method to the Engine
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRouter("GET", pattern, handler)
}

// POST adds a new router with POST method to the Engine
func (e *Engine) POST(pattern string, handler HandlerFunc) {
	e.addRouter("POST", pattern, handler)
}

// PUT adds a new router with PUT method to the Engine
func (e *Engine) PUT(pattern string, handler HandlerFunc) {
	e.addRouter("PUT", pattern, handler)
}

// DELETE adds a new router with DELETE method to the Engine
func (e *Engine) DELETE(pattern string, handler HandlerFunc) {
	e.addRouter("DELETE", pattern, handler)
}

// OPTIONS adds a new router with OPTIONS method to the Engine
func (e *Engine) OPTIONS(pattern string, handler HandlerFunc) {
	e.addRouter("OPTIONS", pattern, handler)
}

// HEAD adds a new router with HEAD method to the Engine
func (e *Engine) HEAD(pattern string, handler HandlerFunc) {
	e.addRouter("HEAD", pattern, handler)
}

// TODO: implement more HTTP methods

func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
