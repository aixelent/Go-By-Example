package main

import (
	"fmt"
	"log"
	"os"
)

func readStructure() {
	file, err := os.Open(".")
	if err != nil {
		log.Fatalf("Failed: %s", err)
	}
	defer file.Close()

	listing, _ := file.Readdirnames(0)
	for _, name := range listing {
		fmt.Println(name)
	}
}

func main() {
	readStructure()
}
