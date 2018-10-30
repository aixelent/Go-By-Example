package main

import "fmt"

func main() {
	if false {
		fmt.Println("First")
	} else if true {
		fmt.Println("Second")
	} else {
		fmt.Println("Third")
	}
}
