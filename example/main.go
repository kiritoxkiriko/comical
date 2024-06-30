package main

import (
	"fmt"
	"html/template"
	"net/http"
	"time"

	"github.com/kiritoxkiriko/comical"
	"github.com/kiritoxkiriko/comical/middleware"
)

func main() {
	r := comical.New()

	r.Use(
		middleware.Recovery(nil),
		middleware.Logger(nil),
	)

	// load template
	formatAsDate := func(t time.Time) string {
		return t.Format("2006-01-02 15:04:05")
	}
	r.SetFuncMap(template.FuncMap{
		"FormatAsDate": formatAsDate,
	})
	r.LoadHTMLTemplates("templates/*")
	r.Static("/assets", "./static")

	r.GET("/index", func(c *comical.Context) {
		c.HTML(http.StatusOK, "index.tmpl", comical.H{
			"now":   time.Now(),
			"title": "comical",
		})
	})

	v1 := r.Group("/v1")
	{
		v1.GET("", func(c *comical.Context) {
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

	// add panic router for test
	r.GET("/panic", func(c *comical.Context) {
		arr := make([]int, 10)
		// overflow, should panic
		c.String(200, fmt.Sprintf("%d", arr[11]))
	})

	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
