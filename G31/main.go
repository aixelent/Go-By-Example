package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Genre struct {
	Techno string `json:"electechno"`
	Id     int    `json:"elecid"`
}

func main() {
	electronic := &Genre{Techno: "Sol Ortega", Id: 1}
	elec, err := json.Marshal(electronic)
	if err != nil {
		log.Fatalln("Error, because error", err)
		return
	}
	fmt.Println(string(elec))
}
