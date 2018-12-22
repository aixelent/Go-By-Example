package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ReadFileBefore (...)
func ReadFileBefore() {
	file, err := ioutil.ReadFile("file1.txt")
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
	fmt.Printf("Length before appending new data to file: %d bytes.\n", len(file))
}

// AppendToFile (...)
func AppendToFile() {
	file, err := os.OpenFile("file1.txt", os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panicln(err)
	}
	defer file.Close()

	toFile, err := file.WriteString("\n\nThe new IDE extends the IntelliJ platform with coding assistance and tool integrations specific for the Go language.")
	if err != nil {
		log.Panicln(err)
	}
	fmt.Printf("Appended %d bytes.", toFile)
}

// ReadFileAfter (...)
func ReadFileAfter() {
	file, err := ioutil.ReadFile("file1.txt")
	if err != nil {
		log.Panicln(err)
		os.Exit(1)
	}
	fmt.Printf("\nLength after appending data to file: %d bytes.\n", len(file))
}

func main() {
	ReadFileBefore()
	AppendToFile()
	ReadFileAfter()
}
