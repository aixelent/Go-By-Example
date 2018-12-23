package main

import (
	"Go-By-Example/G57/calc"
	"fmt"
	"strconv"
)

func main() {
	var a, b, res string

	for {
		fmt.Println("First value:")
		fmt.Scanln(&a)
		fmt.Println("Second value:")
		fmt.Scanln(&b)
		fmt.Println("Choose operation (+, -, *, /):")
		fmt.Scanln(&res)

		num1, err := strconv.Atoi(a)
		if err != nil {
			panic(err)
		}

		num2, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}

		fmt.Println(calc.Calculate(num1, num2, res))
	}
}
