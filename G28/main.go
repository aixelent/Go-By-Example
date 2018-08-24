package main

import "fmt"

func sliceElements(elements []int) []int {
	container := make(map[int]bool)
	result := []int{}

	for values := range elements {
		if container[elements[values]] == true {
		} else {
			container[elements[values]] = true
			result = append(result, elements[values])
		}
	}
	return result
}

func main() {
	elem := []int{1, 3, 5, 7, 9, 9, 11, 13, 15, 13, 17, 19, 21, 19, 21, 19, 13}
	fmt.Println(elem)
	result := sliceElements(elem)
	fmt.Println(result)
}
