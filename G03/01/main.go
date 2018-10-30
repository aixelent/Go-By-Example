package main

import "fmt"

func main() {
	intEvenOdd := func(n int) (int, bool) {
		return n, n%2 == 0
	}

	v1, isEven1 := intEvenOdd(6)
	v2, isEven2 := intEvenOdd(7)
	fmt.Println(v1, isEven1)
	fmt.Println(v2, isEven2)
}
