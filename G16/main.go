package main

import (
	"fmt"
	"log"
	"os"
)

func readStructure() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("Cannot open directory.", err)
	}
	defer file.Close()

	listing, _ := file.Readdir(0)
	fmt.Printf("\nName\t\tSize\tDirectory\t\t\t\tModification")

	for _, files := range listing {
		fmt.Printf("\n%-15s %-7v %-12v %v", files.Name(), files.Size(), files.IsDir(), files.ModTime())
	}
}

func main() {
	readStructure()
}
