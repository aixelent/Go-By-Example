package main

import (
	"log"
	"os"
)

var oldName = "file1.txt"

func renameFile() {
	newName := "file1.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	renameFile()
}
