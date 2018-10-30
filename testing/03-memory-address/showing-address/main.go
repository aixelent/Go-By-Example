package main

import "fmt"

func main()  {
	a := 42

	fmt.Println("a - ", a)
	fmt.Println("a's memory - ", &a)
	fmt.Printf("%d decimal location", &a)
}
