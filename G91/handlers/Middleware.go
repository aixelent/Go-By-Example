package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func CustomHeader(hf echo.HandlerFunc) echo.HandlerFunc {
	return func(e echo.Context) error {
		e.Response().Header().Set(echo.HeaderServer, "Bot")
		return hf(e)
	}
}

func MiddlewareMain(e echo.Context) error {
	return e.String(http.StatusOK, "Response from MiddlewareMain()")
}
