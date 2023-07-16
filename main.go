package main

import (
	"fmt"
	"github.com/kiritoxkiriko/comical"
	"net/http"
)

func main() {
	r := comical.New()

	r.GET("/index", func(c *comical.Context) {
		c.HTML(http.StatusOK, "<hi>Hi, this is a comical world</hi>")
	})

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *comical.Context) {
			c.String(http.StatusOK, "Hello v1")
		})
		v1.GET("/hello", func(c *comical.Context) {
			name := c.Query("name")
			c.String(http.StatusOK, fmt.Sprintf("hello %s, u r at %s\n", name, c.Path))
		})
	}

	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(c *comical.Context) {
			username, password := c.PostForm("username"), c.PostForm("password")
			c.JSON(http.StatusOK, comical.H{
				"username": username,
				"password": password,
			})
		})
	}

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
