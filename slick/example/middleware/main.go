package main

import (
	"log"
	"net/http"
	"time"

	"github.com/HsiaoCz/something/slick"
)

func onlyForV2() slick.Handlerfunc {
	return func(c *slick.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.R.RequestURI, time.Since(t))
	}
}

func main() {
	r := slick.New()
	r.Use(slick.Logger()) // global midlleware
	r.GET("/", func(c *slick.Context) {
		c.HTML(http.StatusOK, "<h1>Hello some</h1>", nil)
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *slick.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
