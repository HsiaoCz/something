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

func HandleGetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello my man"))
}
