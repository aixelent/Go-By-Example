package main

import "fmt"

/*
	[]string - slice
    [50]string - array
*/

func main() {
	var x [58]string
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])

	x[42] = "Dupa"
	fmt.Println(x)
	fmt.Println(x[42])
	// for i := 65; i <= 122; i++ {
	// 	x[i-65] = string(i)
	// }
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[42])
}