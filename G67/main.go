package main

import (
	"io"
	"log"
	"os"
)

func copyFile() {
	oFile, err := os.Open("file1.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer oFile.Close()

	nFile, err := os.Create("newFile.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer nFile.Close()

	bytes, err := io.Copy(nFile, oFile)
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("Copierd %d bytes", bytes)
}

func main() {
	copyFile()
}
