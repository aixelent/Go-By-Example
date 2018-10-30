package main

import (
	"fmt"
)

const (
	_ = iota
	KB = 1 << (iota * 10)	// 1 << (1 * 10)
	MB = 1 << (iota * 10)	// 1 << (2 * 10)
	GB = 1 << (iota * 10)	// 1 << (3 * 10)
	TB = 1 << (iota * 10)	// 1 << (4 * 10)
)

func main() {
	fmt.Println("Hello, playground")
	fmt.Printf("%b\t\n", KB)
	fmt.Printf("%d\t\n", KB)

	fmt.Printf("%b\t\n", MB)
	fmt.Printf("%d\t\n", MB)

	fmt.Printf("%b\t\n", GB)
	fmt.Printf("%d\t\n", GB)

	fmt.Printf("%b\t\n", TB)
	fmt.Printf("%d\t\n", TB)
}