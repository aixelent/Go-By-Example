package main

import "fmt"

func main() {
	greet("Adrian", "Monika")
}

func greet(firName string, secName string) {
	fmt.Println(firName, secName)
}