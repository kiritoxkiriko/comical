package comical

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H is use for json
type H map[string]interface{}

// Context is the encapsulation of http request and response
type Context struct {
	// http handlers objects
	// ResponseWriter is an interface used by an HTTP handler to construct an HTTP response.
	Writer http.ResponseWriter
	// Request is a data structure that represents the client HTTP request.
	Req *http.Request

	// request info
	// Path is the path component of the request URL.
	Path string
	// Method is the HTTP method of the request.
	Method string
	// Params is the parameters of the request.
	Params map[string]string
	// StatusCode is the HTTP status code of the request.
	StatusCode int
	// middleware
	handlers []HandlerFunc
	index    int
}

// newContext creates a new Context object
func newContext(w http.ResponseWriter, r *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    r,
		Path:   r.URL.Path,
		Method: r.Method,
		// omit status code for now
		StatusCode: 0,
		index:      -1,
	}
}

// Next call this function inside middleware, will jump pre request part, and jump intoc post request part
func (c *Context) Next() {
	c.index++
	l := len(c.handlers)
	for c.index < l {
		c.handlers[c.index](c)
		// if some function not call c.Next. then this will not stuck
		c.index++
	}
}

// Fail fail current and all next handler
func (c *Context) Fail(code int, err string) {
	// this will omit next handler and return
	c.Abort()
	c.JSON(code, H{
		"message": err,
	})
}

func (c *Context) Abort() {
	c.index = len(c.handlers)
}

// Err returns the error message
func (c *Context) Err(err error) {
	http.Error(c.Writer, err.Error(), 500)
}

// PostForm sets the request form value by key
func (c *Context) PostForm(key string) (value string) {
	value = c.Req.FormValue(key)
	return
}

// Query sets the request query value by key
func (c *Context) Query(key string) (value string) {
	value = c.Req.URL.Query().Get(key)
	return
}

func (c *Context) Param(key string) (value string) {
	value, _ = c.Params[key]
	return
}

// Status sets the response status code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets the response header
func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

// String sets the response body with string
func (c *Context) String(code int, format string, values ...any) {
	// set header to plain text first
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	_, err := c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
	if err != nil {
		c.Err(err)
	}
}

// JSON sets the response body with json
func (c *Context) JSON(code int, obj any) {
	// set header to json first
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		c.Err(err)
	}
}

// Data sets the response body with bytes
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	_, err := c.Writer.Write(data)
	if err != nil {
		c.Err(err)
	}
}

// HTML sets the response body with html string
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	_, err := c.Writer.Write([]byte(html))
	if err != nil {
		c.Err(err)
	}
}
