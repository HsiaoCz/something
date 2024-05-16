package main

import (
	"net/http"

	"github.com/HsiaoCz/something/slick"
)

func main() {
	app := slick.New()

	app.GET("/user", HandleGetUser)
	app.GET("/hello", HandleHello)
	app.GET("/hello/:name", HandleHelloName)
	app.Run(":9001")
}

func HandleGetUser(c *slick.Context) {
	c.JSON(http.StatusOK, slick.H{
		"message": "all is well",
	})
}

func HandleHello(c *slick.Context) {
	c.JSON(http.StatusOK, slick.H{
		"message": "all is well",
	})
}

func HandleHelloName(c *slick.Context) {
	c.JSON(http.StatusOK, slick.H{
		"message": c.Param("name"),
	})
}
