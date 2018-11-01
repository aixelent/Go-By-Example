package main

import "fmt"

func main()  {
	st := "Hello"
	fmt.Println(st)

	fmt.Println(ReverseString(st))
}

func ReverseString(st string) string {
	var rev string
	for i := len(st) - 1; i >= 0; i-- {
		rev += string(st[i])
	}
	return rev
}
