package comical

import (
	"html/template"
	"log"
	"net/http"
	"path"
)

const (
	filePathPattern = "*filepath"
)

type RouteGroup struct {
	// router extend from router
	*router
	prefix string
	// not currently in use
	parent        *RouteGroup
	middlewares   []HandlerFunc
	htmlTemplates *template.Template // for html render
	funcMap       template.FuncMap   // for html render
}

func newRootGroup() *RouteGroup {
	return &RouteGroup{
		prefix: "",
		router: newRouter(),
	}
}

// Group create a group that match prefix
func (g *RouteGroup) Group(prefix string) *RouteGroup {
	prefix = prefix + g.prefix
	return &RouteGroup{
		prefix: prefix,
		parent: g,
		router: g.router,
		// set a empty funcMap in case nil panic
		funcMap: make(template.FuncMap),
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

// for static file

// createStaticHandler create static handler
func (g *RouteGroup) createStaticHandler(relPath string, fs http.FileSystem) HandlerFunc {
	// absPath for root router
	absPath := path.Join(g.prefix, relPath)
	// remove router path prefix to get real file path
	fileServer := http.StripPrefix(absPath, http.FileServer(fs))

	handler := func(c *Context) {
		filePath := c.Param(filePathPattern)
		// check file permission first
		if _, err := fs.Open(filePath); err != nil {
			// if error return 404
			c.Status(http.StatusNotFound)
			return
		}
		// use builtin file server
		fileServer.ServeHTTP(c.Writer, c.Req)
	}
	return handler
}

// Static map a file path root to relPath at router
func (g *RouteGroup) Static(relPath, root string) {
	handler := g.createStaticHandler(relPath, http.Dir(root))
	pattern := path.Join(relPath, filePathPattern)
	// register a static handler
	g.GET(pattern, handler)
}

// for html template

// SetFuncMap set custom render function for template, only support for root group
func (g *RouteGroup) SetFuncMap(funcMap template.FuncMap) {
	g.funcMap = funcMap
}

// LoadHTMLTemplates load html template, only support for root group
func (g *RouteGroup) LoadHTMLTemplates(pattern string) {
	g.htmlTemplates = template.Must(template.New("").Funcs(g.funcMap).ParseGlob(pattern))
}
