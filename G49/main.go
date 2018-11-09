package main

import (
	crypto "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

func cryptoRand() {
	gen := genCrypto()
	ran := rand.New(rand.NewSource(gen))
	for i := 0; i < 10; i++ {
		first := ran.Int()
		fmt.Printf("Math %d\n", first)
	}
}

func genCrypto() int64 {
	safe, err := crypto.Int(crypto.Reader, big.NewInt(100234))
	if err != nil {
		panic(err)
	}
	return safe.Int64()
}

func main() {
	cryptoRand()
}
