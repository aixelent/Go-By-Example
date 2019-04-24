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

	g.Use(middleware.BasicAuth(func(userLogin, password string, e echo.Context) (bool, error) {
		if userLogin == "admin" && password == "admin123" {
			return true, nil
		}
		return false, nil
	}))

	g.GET("/main", handlers.MiddlewareMain)

	e.Logger.Fatal(e.Start(":8080"))
}
