package handlers

import (
	"github.com/labstack/echo"
	"net/http"
)

func MiddlewareMain(e echo.Context) error {
	return e.String(http.StatusOK, "Response from MiddlewareMain()")
}
