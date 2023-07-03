package main

import (
	"fmt"
	"github.com/kiritoxkiriko/comical"
	"net/http"
)

func main() {
	c := comical.New()
	c.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	c.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})

	c.Run(":8080")
}
