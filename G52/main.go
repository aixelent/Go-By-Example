package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func mapData() {
	languages := make(map[string]interface{})
	languages["name"] = "Python"
	languages["year"] = "1991"
	languages["type"] = map[string]interface{}{
		"OOP":        "Yes",
		"Functional": "Yes",
		"Imperative": "Yes",
		"Procedural": "Tes",
	}

	data, err := json.Marshal(languages)
	if err != nil {
		log.Fatalln(err.Error())
		return
	}

	jsonStr := string(data)
	fmt.Println("JSON format:", jsonStr)
}

func main() {
	mapData()
}
