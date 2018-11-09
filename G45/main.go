package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

const filePath = "file.csv"

func readFile() {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	reader := csv.NewReader(f)
	reader.FieldsPerRecord = 2

	for {
		record, err := reader.Read()
		if err != nil {
			panic(err)
			break
		}
		fmt.Println(record)
	}
}

func main() {
	readFile()
}
