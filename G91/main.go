package main

import (
	"Go-By-Example/G91/handlers"
	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.Use(handlers.CustomHeader)
	g := e.Group("/middleware")

	g.GET("/main", handlers.MiddlewareMain)

	e.Logger.Fatal(e.Start(":8080"))
}
