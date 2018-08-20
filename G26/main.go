package main

import (
	"fmt"
	"reflect"
)

func main() {
	num1 := []int{1, 2, 4, 6, 8}
	num2 := []int{0, 3, 5, 7, 9}

	fmt.Println("Slices equals?: ", reflect.DeepEqual(num1, num2))
}
