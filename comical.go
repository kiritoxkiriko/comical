package comical

import (
	"net/http"
	"strings"
)

// HandlerFunc defines the request handler used by comical
type HandlerFunc func(c *Context)

// Engine implements the interface of http.Handler
type Engine struct {
	// router map: key is "METHOD PATTERN", value is HandlerFunc
	*RouteGroup
	groups []*RouteGroup
}

// New is the constructor of comical.Engine
func New() *Engine {
	g := newRootGroup()
	return &Engine{
		RouteGroup: g,
		groups:     []*RouteGroup{g},
	}
}

// Group override RouteGroup method
func (e *Engine) Group(prefix string) *RouteGroup {
	g := e.RouteGroup.Group(prefix)
	e.groups = append(e.groups, g)
	return g
}

// ServeHTTP makes the Engine implement the interface of http.Handler
func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var middlewares []HandlerFunc
	for _, group := range e.groups {
		// only apply group middleware to matched group
		if strings.HasPrefix(r.URL.Path, group.prefix) {
			middlewares = append(middlewares, group.middlewares...)
		}
	}
	c := newContext(w, r)
	// add middleware to context
	c.handlers = middlewares
	// add templates to context
	c.templates = e.htmlTemplates
	e.handle(c)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
