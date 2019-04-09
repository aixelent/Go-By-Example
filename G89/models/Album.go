package models

import "time"

type Album struct {
	Artist   string    `json:"artist"`
	Title    string    `json:"title"`
	Genre    string    `json:"genre"`
	Label    string    `json:"label"`
	Style    string    `json:"style"`
	Released time.Time `json:"released"`
}
