package main

import (
	"fmt"
	"reflect"
)

func main() {
	firstMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}
	secondMap := map[int]int{
		1: 1,
		2: 2,
		3: 3,
		4: 4,
		5: 5,
	}

	fmt.Println("Map equals?: ", reflect.DeepEqual(firstMap, secondMap))
}
