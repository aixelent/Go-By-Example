package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {
	fName := "file.txt"
	buffer, err := ioutil.ReadFile(fName)
	if err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}

	input := string(buffer)
	data := bufio.NewScanner(strings.NewReader(input))
	data.Split(bufio.ScanRunes)

	for data.Scan() {
		fmt.Print(data.Text())
	}
}
