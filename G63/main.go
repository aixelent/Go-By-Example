package main

import (
	"log"
	"os"
)

var fileName = "file1.txt"

func deleteFile() {
	err := os.Remove(fileName)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	deleteFile()
}
