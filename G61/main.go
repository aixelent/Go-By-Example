package main

import (
	"fmt"
	"log"
	"os"
)

func getFileInfo() {
	info, err := os.Stat("file1.txt")
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
