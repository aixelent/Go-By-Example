package models

import "time"

type Artist struct {
	Name    string    `json:"name"`
	Surname string    `json:"surname"`
	Alias   string    `json:"alias"`
	Born    time.Time `json:"born"`
}
