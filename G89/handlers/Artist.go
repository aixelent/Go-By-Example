package handlers

import (
	"Go-By-Example/G89/models"
	"encoding/json"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func AddArtist(e echo.Context) error {
	artist := models.Artist{}
	defer e.Request().Body.Close()

	err := json.NewDecoder(e.Request().Body).Decode(&artist)
	if err != nil {
		log.Printf("Failed artist request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	now := time.Now()
	artist.Born = now

	log.Printf("Alias: %v", artist.Alias)
	log.Printf("Name: %v", artist.Name)
	log.Printf("Surname: %v", artist.Surname)
	log.Printf("Born: %v", artist.Born)

	return e.String(http.StatusOK, "done!")
}
