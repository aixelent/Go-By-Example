package handlers

import (
	"Go-By-Example/G88/models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func AddAlbum(e echo.Context) error {
	albums := models.Album{}

	defer e.Request().Body.Close()
	r, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		log.Fatalf("Cannot read request body %s:", err)
		return e.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(r, &albums)
	if err != nil {
		log.Fatalf("Failed unmarshal data: %s", err)
		return e.String(http.StatusInternalServerError, "")
	}
	log.Printf("Album: %s", albums)
	return e.String(http.StatusOK, "Album response - Done.")
}

func GetAlbum(e echo.Context) error {
	albumName := e.QueryParam("name")
	albumGenre := e.QueryParam("genre")

	data := e.Param("data")

	// http://localhost:8080/album/string?name=Bvdub&genre=Electronic
	if data == "string" {
		// http://localhost:8080/album?name=Bvdub&genre=Electronic
		return e.String(http.StatusOK, fmt.Sprintf("Album Name: %s.\nGenre: %s", albumName, albumGenre))
	}

	// http://localhost:8080/album/json?name=Bvdub&genre=Electronic
	if data == "json" {
		return e.JSON(http.StatusOK, map[string]string{
			"name":  albumName,
			"genre": albumGenre,
		})
	}

	return e.JSON(http.StatusBadRequest, map[string]string{
		"error:": "asdf",
	})
}
