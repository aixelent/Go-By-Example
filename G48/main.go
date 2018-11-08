package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randMath() {
	ran := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 50; i++ {
		first := ran.Int()
		fmt.Printf("Math: %d\n", first)
	}
}

func main() {
	randMath()
}
