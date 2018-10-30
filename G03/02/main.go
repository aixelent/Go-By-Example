package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Scan("Enter number(s) to check if the're even or odd:")
	numRead := bufio.NewReader(os.Stdin)

	// Reading multiple integers (with spaces)
	num, err := numRead.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	numSlice := strings.Fields(num)
	for _, v := range numSlice {
		i, _ := strconv.Atoi(v)
		if i%2 == 0 {
			fmt.Println(v, "is even")
		} else {
			fmt.Println(v, "is odd")
		}
	}
}
