package handlers

import (
	"Go-By-Example/G89/models"
	"github.com/labstack/echo"
	"log"
	"net/http"
	"time"
)

func AddEvent(e echo.Context) error {
	event := models.Event{}

	err := e.Bind(&event)
	if err != nil {
		log.Printf("Failed event request: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError)
	}

	now := time.Time{}
	event.Date = now

	log.Printf("Event: %v", event.Name)
	log.Printf("Location: %v", event.Location)
	log.Printf("Date: %v", event.Date)

	return e.String(http.StatusOK, "done!")
}
