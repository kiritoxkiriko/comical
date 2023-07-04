package main

import (
	"fmt"
	"github.com/kiritoxkiriko/comical"
	"net/http"
)

func main() {
	r := comical.New()

	r.GET("/", func(c *comical.Context) {
		c.HTML(http.StatusOK, "<h1>Welcome to Comical World!</h1>")
	})

	r.GET("/hello", func(c *comical.Context) {
		name := c.Query("name")
		if name == "" {
			name = "Guest"
		}
		c.String(http.StatusOK, fmt.Sprintf("Hello %s, you're at %s", name, c.Path))
	})

	r.POST("/login", func(c *comical.Context) {
		user, passwd := c.PostForm("user"), c.Query("passwd")
		c.JSON(http.StatusOK, comical.H{
			"user":   user,
			"passwd": passwd,
		})
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
