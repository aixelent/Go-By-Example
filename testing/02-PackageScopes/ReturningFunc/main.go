package main

import "fmt"

/* Mozna zwracac cale funckje z metody. Tak jak ponizej */

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main()  {
	increment := wrapper()
	fmt.Println(increment())
	fmt.Println(increment())
}


