package main

import (
	"log"
	"os"
)

func create(file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(f)
	f.Close()
}

func main() {
	create("file1.txt")
}
