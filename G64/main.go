package main

import (
	"log"
	"os"
)

var fileName = "file1.txt"

func isFileExists() {
	f, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("File does not exists")
		}
	}
	log.Println("File info: ", &f)
}

func main() {
	isFileExists()
}
