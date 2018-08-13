package main

import (
	"fmt"
	"sort"
)

func main() {
	ints1 := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	ints2 := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32}

	fmt.Println(ints1, "\nCapacity: ", cap(ints1))
	fmt.Println(ints2, "\nCapacity: ", cap(ints2))

	a := append(ints1, ints2...)
	sort.Ints(a)
	fmt.Println(a, "\nCapacity: ", cap(a))
}
