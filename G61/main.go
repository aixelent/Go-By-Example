package main

import (
	"fmt"
	"log"
	"os"
)

var fileName = "file1.txt"

func getFileInfo() {
	info, err := os.Stat(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Is directory: ", info.IsDir())
	fmt.Println("Filename: ", info.Name())
	fmt.Println("Size: ", info.Size())
	fmt.Println("Modification time: ", info.ModTime())
	fmt.Println("Permission: ", info.Mode())
}

func main() {
	getFileInfo()
}
