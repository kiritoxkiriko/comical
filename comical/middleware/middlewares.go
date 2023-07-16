package middleware

import (
	"github.com/kiritoxkiriko/comical"
	"log"
	"time"
)

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
