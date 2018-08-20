package main

import (
	"fmt"
	"sort"
)

func main() {
	num := []int{1, 7, 5, 3, 9, 13, 11, 15, 21, 19, 17, 25, 23, 31, 27, 31}
	for index, elem := range num {
		fmt.Println(index, " - ", elem)
	}

	num = append(num, 6)
	sort.Sort(sort.IntSlice(num))
	for index, elem := range num {
		fmt.Println(index, " - ", elem)
	}
	fmt.Println(num)
}
