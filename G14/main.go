package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func NewFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatalf("File not created", err)
	}
	defer file.Close()

	bytes, err := file.WriteString("report")
	if err != nil {
		log.Fatalf("Failed writing to file.")
	}

	fmt.Printf("\nFile: %s, successfully created", file.Name())
	fmt.Printf("\nWrote %d bytes", bytes)
}

func ReadFile() {
	bytes, err := ioutil.ReadFile("file.txt")
	if err != nil {
		log.Fatalf("Cannot reading data from file %s", err)
	}
	fmt.Printf("\nBytes length: %d", len(bytes))
	fmt.Printf("\nData: %s", bytes)
	fmt.Printf("\nError: %v", err)
}

func main() {
	fmt.Printf("Creating file...\n")
	NewFile()
	fmt.Printf("\n\nReading file...")
	ReadFile()
}
