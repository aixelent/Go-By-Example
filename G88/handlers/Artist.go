package handlers

import (
	"Go-By-Example/G88/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func AddArtist(e echo.Context) error {
	artists := models.Artist{}

	defer e.Request().Body.Close()

	err := json.NewDecoder(e.Request().Body).Decode(&artists)
	if err != nil {
		log.Fatalf("Failed unmarshal data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	log.Printf("Artist: %s", &artists)
	return e.String(http.StatusOK, "Artist response - Done.")
}
