package main

import (
	"Go-By-Example/G88/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/album/:data", handlers.GetAlbum)

	e.Logger.Fatal(e.Start(":8080"))
}
