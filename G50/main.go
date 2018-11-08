package main

import (
	"fmt"
	"math/rand"
)

func randString(n int) string {
	bytes := make([]byte, n)
	for i := 0; i < 20; i++ {
		bytes[i] = byte(randInt(97, 122))
	}
	return string(bytes)
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min)
}

func main() {
	fmt.Println(randString(20))
}
