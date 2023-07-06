package comical

import (
	"log"
	"net/http"
)

// Router is the interface that wraps the Handle method
type Router interface {
	// Handle handles http requests by Context
	Handle(c *Context)
	// AddRoute adds a new route to the defaultRouter
	AddRoute(method, pattern string, handler HandlerFunc)
}

// DefaultRouter is a request multiplexer for comical
type DefaultRouter struct {
	// handlers stores the mapping of request path to handler
	handlers map[string]HandlerFunc
}

// NewDefaultRouter creates a new defaultRouter object
func NewDefaultRouter() *DefaultRouter {
	return &DefaultRouter{
		handlers: map[string]HandlerFunc{},
	}
}

// AddRoute adds a new route to the defaultRouter
func (r *DefaultRouter) AddRoute(method, pattern string, handler HandlerFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// Handle handles http requests by Context
func (r *DefaultRouter) Handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
