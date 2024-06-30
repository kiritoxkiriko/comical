package middleware

import (
	"fmt"
	"github.com/kiritoxkiriko/comical"
	"log"
	"time"
)

// Logger add logger for each request
func Logger(logger *log.Logger) comical.HandlerFunc {
	if logger == nil {
		logger = log.Default()
	}
	handler := func(c *comical.Context) {
		tStart := time.Now()
		c.Next()
		diff := time.Since(tStart)
		log.Printf("[%d] %s in %v", c.StatusCode, c.Req.RequestURI, diff)
	}
	return handler
}

// HeaderChecker check header for each request, if not have, return http 400
func HeaderChecker(key, value string) comical.HandlerFunc {
	handler := func(c *comical.Context) {
		val := c.Req.Header.Get(key)
		if val != value {
			c.Fail(400, fmt.Sprintf("header %s has wrong value", key))
		}
		c.Next()
	}
	return handler
}
