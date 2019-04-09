package handlers

import (
	"Go-By-Example/G89/models"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"io/ioutil"
	"log"
	"net/http"
)

func AddAlbum(e echo.Context) error {
	album := models.Album{}
	defer e.Request().Body.Close()

	body, err := ioutil.ReadAll(e.Request().Body)
	if err != nil {
		log.Printf("Failed reading request body: %s", err)
		return e.String(http.StatusInternalServerError, "")
	}

	err = json.Unmarshal(body, &album)
	if err != nil {
		log.Printf("Failed unmarshal: %s", err)
		return e.String(http.StatusInternalServerError, "")
	}

	log.Printf("Artist: %v", album.Artist)
	log.Printf("Album name: %v", album.Title)
	log.Printf("Released: %v", album.Released)
	log.Printf("Genre: %v", album.Genre)
	log.Printf("Label: %v", album.Label)
	log.Printf("Style: %v", album.Style)

	return e.String(http.StatusOK, "done!")
}

func GetAlbum(e echo.Context) error {
	albumName := e.QueryParam("name")
	albumGenre := e.QueryParam("genre")

	data := e.Param("data")

	if data == "string" {
		// http://localhost:8080/album?name=Bvdub&genre=Electronic
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
