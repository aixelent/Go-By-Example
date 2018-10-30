package main

import "fmt"

func main() {
	fmt.Println(greet("Adrian", "Monika"))
}

func greet(firName, secName string) string {
	return fmt.Sprint(firName, secName)
}