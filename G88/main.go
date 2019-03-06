package main

import (
	"Go-By-Example/G88/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// http://localhost:8080/album/json?name=Bvdub&genre=Electronic
	e.GET("/album/:data", handlers.GetAlbum)
	e.POST("/album", handlers.AddAlbum)
	e.POST("/artist", handlers.AddArtist)
	e.POST("/event", handlers.AddEvent)
	e.Logger.Fatal(e.Start(":8080"))
}
