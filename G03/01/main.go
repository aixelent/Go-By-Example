package main

import "fmt"

func intEvenOdd(n int) (int, bool) {
	return n, n%2 == 0
}

func main() {
	v1, isEven1 := intEvenOdd(6)
	v2, isEven2 := intEvenOdd(11)
	fmt.Println(v1, isEven1)
	fmt.Println(v2, isEven2)
}
