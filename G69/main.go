package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomSlice(size int) []int {
	s := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		s[i] = rand.Intn(99999)
	}
	return s
}

func main() {
	s := randomSlice(100)
	fmt.Printf("Random numbers: %v\n", s)
}
