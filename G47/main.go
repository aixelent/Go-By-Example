package main

import (
	"fmt"
	"strconv"
)

const bin string = "100101001"

//const flo string = "15.55485"
const dec string = "15"
const hex string = "15F"
const oct string = "10"

func binToDec() {
	b, err := strconv.ParseInt(bin, 2, 64)
	if err != nil {
		panic(err)
	}
	fmt.Println(b)
}

func hexToDec() {
	h, err := strconv.ParseInt(hex, 16, 16)
	if err != nil {
		panic(err)
	}
	fmt.Println(h)
}

func octToHex() {
	o, err := strconv.ParseInt(oct, 8, 16)
	if err != nil {
		panic(err)
	}
	fmt.Println(o)
}

func decToOct() {
	dec, err := strconv.ParseInt(dec, 10, 8)
	if err != nil {
		panic(err)
	}
	fmt.Println(dec)
}

func main() {
	binToDec()
	hexToDec()
	octToHex()
	decToOct()
}
