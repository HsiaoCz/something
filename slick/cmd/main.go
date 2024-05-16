package main

import (
	"net/http"

	"github.com/HsiaoCz/something/slick"
)

func main() {
	app := slick.New()

	app.GET("/user", HandleGetUser)

	app.Run(":9001")
}

func HandleGetUser(c *slick.Context) {
	c.JSON(http.StatusOK, slick.H{
		"message": "all is well",
	})
}
