package main

import (
	"log"
	"os"
)

var fileName = "file1.txt"

func perm() {
	err := os.Chmod(fileName, 0755)
	if err != nil {
		log.Fatalln(err)
	}
}

func main() {
	perm()
}
