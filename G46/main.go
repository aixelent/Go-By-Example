package main

import (
	"fmt"
	"strconv"
)

const vInt = "12345"
const vFloat = "15.5"
const bin = "0111010"

func convertFloat() {
	flo, err := strconv.ParseFloat(vFloat, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Float: %f\n", flo)
}

func convertInt() {
	in, err := strconv.Atoi(vInt)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Int: %d\n", in)
}

func convertBin() {
	b, err := strconv.ParseInt(bin, 2, 32)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Binary: %b\n", b)
}

func main() {
	convertFloat()
	convertInt()
	convertBin()
}
