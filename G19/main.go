package main

import (
	"bytes"
	"fmt"
)

func main() {
	data := []byte("Java Python Golang Scala Elixir Cobol")
	space := []byte{' '}
	sampleSplit := bytes.Split(data, space)
	fmt.Printf("To split: %q", data)
	for index, elem := range sampleSplit {
		fmt.Printf("\n%d - %q", index, elem)
	}
}
