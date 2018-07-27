package main

import "fmt"

func main() {
	fmt.Print("Enter first number to swap: ")
	var first int
	fmt.Scanln(&first);

	fmt.Print("Enter second number to swap: ")
	var second int
	fmt.Scanln(&second)

	first = first - second
	second = first + second
	first = second - first

	fmt.Println("First number now is: ", first)
	fmt.Print("Second number now is: ", second, "\n")
}