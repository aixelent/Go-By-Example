package main

import (
	"Go-By-Example/G89/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/album/:data", handlers.GetAlbum)
	e.POST("/album", handlers.AddAlbum)
	e.POST("/artist", handlers.AddArtist)
	e.POST("/event", handlers.AddEvent)

	e.Logger.Fatal(e.Start(":8080"))
}
