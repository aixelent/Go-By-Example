package main

import (
	"fmt"
	"reflect"
)

type myStruct struct {
	a string
	b []int
}

func main() {
	st1 := myStruct{a: "Charlotte de Witte", b: []int{1, 3, 5, 7, 9}}
	st2 := myStruct{a: "Charlotte de Witte", b: []int{1, 3, 5, 7, 9}}
	fmt.Println("Struct equal: ", reflect.DeepEqual(st1, st2))
}
