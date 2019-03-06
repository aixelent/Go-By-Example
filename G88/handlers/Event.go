package handlers

import (
	"Go-By-Example/G88/models"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func AddEvent(e echo.Context) error {
	events := models.Event{}

	err := e.Bind(&events)
	if err != nil {
		log.Fatalf("Failed unmarshal data: %s", err)
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	log.Printf("Event: %s", events)
	return e.String(http.StatusOK, "Event response - Done.")
}

// slowest
// func AddEvent(e echo.Context) (err error) {
// 	events := new(Event)
// 	if err := e.Bind(events); err != nil {
// 		return err
// 	}
// 	return e.JSON(http.StatusOK, events)
// }
