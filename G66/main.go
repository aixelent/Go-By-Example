package main

import (
	"log"
	"os"
	"time"
)

var fileName = "file1.txt"

func timestamp() {
	modTime := time.Now().Add(24 * time.Hour)
	accessTime := time.Now().Add(48 * time.Hour)
	lastModDate := modTime
	lastAccessDate := accessTime

	err := os.Chtimes(fileName, lastModDate, lastAccessDate)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	timestamp()
}
