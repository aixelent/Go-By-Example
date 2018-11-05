package main

import (
	"encoding/csv"
	"os"
)

var file = "file.csv"

func makeCSV() {
	f, err := os.Create(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
}

func makeData() {
	f, err := os.OpenFile(file, os.O_WRONLY|os.O_APPEND, 0755)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	x := []string{"Language", "Year"}
	y := []string{"Cobol", "1960"}
	y1 := []string{"Golang", "2009"}

	csvWrite := csv.NewWriter(f)
	strWrite := [][]string{x, y, y1}
	csvWrite.WriteAll(strWrite)
	csvWrite.Flush()
}

func main() {
	makeCSV()
	makeData()
}
