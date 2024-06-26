package main

import (
	"net/http"

	"github.com/HsiaoCz/something/slick"
)

func main() {
	r := slick.New()
	r.GET("/index", func(c *slick.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>", nil)
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *slick.Context) {
			c.HTML(http.StatusOK, "<h1>Hello Gee</h1>", nil)
		})

		v1.GET("/hello", func(c *slick.Context) {
			// expect /hello?name=geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *slick.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *slick.Context) {
			c.JSON(http.StatusOK, slick.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
