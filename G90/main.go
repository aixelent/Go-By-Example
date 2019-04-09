package main

import (
	"Go-By-Example/G90/handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()
	g := e.Group("/middleware")

	g.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           `[${time_rfc3339}] ${method} ${status} ${host}${path} ${latency} ${latency_human}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05",
	}))
	g.GET("/main", handlers.MiddlewareMain)

	e.Logger.Fatal(e.Start(":8080"))
}
