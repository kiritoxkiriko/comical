package comical

import (
	"log"
	"net/http"
)

// router is a request multiplexer for comical
type router struct {
	// handlers stores the mapping of request path to handler
	handlers map[string]HandlerFunc
}

// newRouter creates a new router object
func newRouter() *router {
	return &router{
		handlers: map[string]HandlerFunc{},
	}
}

// addRoute adds a new route to the router
func (r *router) addRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle handles http requests by Context
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
