package middleware

import (
	"fmt"
	"github.com/kiritoxkiriko/comical"
	"log"
	"net/http"
	"runtime"
	"strings"
)

// print stack trace for debug
func trace(message string) string {
	var pcs [32]uintptr
	n := runtime.Callers(3, pcs[:]) // skip first 3 caller

	var str strings.Builder
	str.WriteString(message + "\nTraceback:")
	for _, pc := range pcs[:n] {
		fn := runtime.FuncForPC(pc)
		file, line := fn.FileLine(pc)
		str.WriteString(fmt.Sprintf("\n\t%s:%d", file, line))
	}
	return str.String()
}

// Recovery a middleware that support simple recover from panic and return HTTP 500
func Recovery(logger *log.Logger) comical.HandlerFunc {
	if logger == nil {
		logger = log.Default()
	}
	handler := func(c *comical.Context) {
		defer func() {
			if err := recover(); err != nil {
				errMsg := fmt.Sprintf("%s", err)
				logger.Printf("%s\n\n", trace(errMsg))
				c.Fail(500, http.StatusText(500))
			}
		}()
		c.Next()
	}
	return handler
}
