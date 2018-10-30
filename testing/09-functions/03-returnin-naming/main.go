package main

import "fmt"

func main() {
	fmt.Println(greet("Adrian", "Monika"))
}

func greet(firName string, secName string) (s string) {
	s = fmt.Sprint(firName, secName)
	return
}