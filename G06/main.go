package main

import "fmt"

func main() {
	ints := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	fmt.Println("Slice:", ints)

	last := ints[:1]
	fmt.Println(last)

	var a int

	// Assign to the first element '0' index
	// last element
	ints[0] = ints[len(ints)-1]

	// Give the last element 'a' value
	// undeclared 'a' so value is '0'
	ints[len(ints)-1] = a

	// Adding the slice before first element
	ints = ints[:len(ints)-1]

	fmt.Println(ints)
}
