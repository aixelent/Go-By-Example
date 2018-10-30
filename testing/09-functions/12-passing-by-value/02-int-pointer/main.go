package main

import "fmt"

func main() {
	age := 42
	fmt.Println(&age)		// 0xc42001c0c0

	changeMe(&age)

	fmt.Println(&age)		// 0xc42001c0c0
	fmt.Println(age)		// 42
}

func changeMe(z *int) {
	fmt.Println(z)			// 0xc42001c0c0
	fmt.Println(*z)			// 42
	*z = 24
	fmt.Println(z)			// 0xc42001c0c0
	fmt.Println(*z)			// 24
}