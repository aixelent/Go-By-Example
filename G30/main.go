package main

import "fmt"

func main() {
	intSlice := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	displayInt := 21
	fmt.Println("Slice of int: ", intSlice)

	if len(intSlice) == 0 || (intSlice)[0] == displayInt {
		fmt.Println(intSlice)
		return
	} else if (intSlice)[len(intSlice) - 1] == displayInt {
		intSlice = append([]int{displayInt}, intSlice[:len(intSlice)-1]...)
		fmt.Println(intSlice)
		return
	}

	for p, x := range intSlice {
		if x == displayInt {
			intSlice = append([]int{displayInt}, append(intSlice[:p], intSlice[p+1:]...)...)
			break
		}
	}
	fmt.Println("Slice after moving element: ", intSlice)
}
