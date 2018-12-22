package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

var filePath = "file1.txt"

func createFile() {
	f, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("File created.")
}

func readFile() {
	file, err := os.OpenFile(filePath, os.O_RDWR, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	f, err := ioutil.ReadFile(filePath)
	s := string(f)
	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}

func updateFileContent() {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	toFile, err := file.WriteString("\n\nThe new IDE extends the IntelliJ platform with coding assistance and tool integrations specific for the Go language.")
	if err != nil {
		panic(err)
	}
	fmt.Printf("String was added to file correctly. Added %d bytes.", toFile)
}

func deleteFile() {
	err := os.Remove(filePath)
	if err != nil {
		panic(err)
	}
	fmt.Println("File was deleted.", filePath)
}

func main() {
	createFile()
	updateFileContent()
	readFile()
	deleteFile()
}
