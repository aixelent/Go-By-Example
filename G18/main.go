package main

import (
	"fmt"
	"strings"
)

func main() {
	dataString := "Java Python Golang Scala Elixir Cobol"
	dataArray := strings.Fields(dataString)
	fmt.Println(dataString)

	for _, v := range dataArray {
		fmt.Println(v)
	}
}
