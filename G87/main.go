package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func EchoHello(e echo.Context) error {
	return e.String(http.StatusOK, "Hello Echo!")
}

func main() {
	e := echo.New()
	e.GET("/", EchoHello)

	e.Logger.Fatal(e.Start(":8080"))
}
