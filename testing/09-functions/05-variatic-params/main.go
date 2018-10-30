package main

import "fmt"

func main() {
	n := average(43, 55, 61, 62, 64, 69, 75)
	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64 = 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
