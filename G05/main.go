package main

import "fmt"

func main() {
	ints := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	fmt.Println("Slice:", ints)

	first := ints[len(ints) - 1]
	fmt.Println(first)

	remove := ints[:len(ints) - 1]
	fmt.Println(remove)
}
