package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []string{"Java", "Python", "Golang", "Scala", "Elixir", "Cobol"}
	sort.Sort(sort.StringSlice(data))
	fmt.Println(data)

	// Scala after sorting on position 5
	fmt.Println("Scala on position: ", sort.StringSlice(data).Search("Scala"))

	num := []int{1, 7, 5, 3, 9, 13, 11, 15, 21, 19, 17, 25, 23, 31, 27, 31}
	fmt.Println(num)
	sort.Sort(sort.IntSlice(num))
	// after sorting
	fmt.Println(num)

	fmt.Println("7 is on position: ", sort.IntSlice(num).Search(7))
}
