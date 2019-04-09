package handlers

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
)

func GetAlbum(e echo.Context) error {
	albumName := e.QueryParam("name")
	albumGenre := e.QueryParam("genre")

	data := e.Param("data")

	if data == "string" {
		return e.String(http.StatusOK, fmt.Sprintf("Album Name: %s\nGenre: %s", albumName, albumGenre))
	}

	if data == "json" {
		return e.JSON(http.StatusOK, map[string]string{
			"name":  albumName,
			"genre": albumGenre,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]string{
		"error:": "Bad data type - string or json only",
	})
}
