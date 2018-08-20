package main

import "fmt"

func main() {
	mapp := map[int]string{1: "A", 2: "B", 3: "C", 4: "D", 5: "E"}
	toMapp := map[int]string{}

	// Copying from mapp -> toMapp
	for index, elem := range mapp {
		toMapp[index] = elem
	}

	// iterating over toMapp to list elements
	for index, elem := range toMapp {
		fmt.Println(index, " - ", elem)
	}
}
