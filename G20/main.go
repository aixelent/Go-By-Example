package main

import "fmt"

func main() {
	data := map[string]int{"Java": 1, "Python": 2, "Golang": 3, "Scala": 4, "Elixir": 5, "Cobol": 6}
	value, key := data["Scala"]
	if key {
		fmt.Println("Key founded. Value: ", value)
	} else {
		fmt.Println("Key not found.")
	}
	fmt.Println()

	// or...
	if value, exist := data["Python"]; exist {
		fmt.Println("Key founded. Value: ", value)
	} else {
		fmt.Println("Key not found.")
	}

	// or...
	nData := map[int]string{1: "Java", 2: "Python", 3: "Golang", 4: "Scala", 5: "Elixir", 6: "Cobol"}
	fmt.Printf("\nKey 0 exists: %t\n"+
		"Key 1 exists: %t\n"+
		"Key 2 exists: %t\n"+
		"Key 3 exists: %t\n"+
		"Key 4 exists: %t\n"+
		"Key 5 exists: %t\n"+
		"Key 6 exists: %t\n",
		nData[0] != "", nData[1] != "", nData[2] != "", nData[3] != "", nData[4] != "", nData[5] != "", nData[6] != "")
}
