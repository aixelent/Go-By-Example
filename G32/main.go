package main

import (
	"encoding/json"
	"fmt"
)

func JSON(str string) bool {
	var jsonString string
	err := json.Unmarshal([]byte(str), &jsonString)
	return err == nil
}

func jString(str string) bool {
	var jsonString string
	err := json.Unmarshal([]byte(str), &jsonString)
	return err == nil
}

func main() {
	fmt.Println(jString("Scala"))
	fmt.Println(jString(`"Cobol"`))
}
