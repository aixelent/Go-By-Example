package main

import (
	"fmt"
	"log"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}

	_, ip4, err := net.ParseCIDR(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	getMask := ip4.Mask
	if len(getMask) != 4 {
		panic("Length must be 4 bytes")
	}

	fmt.Println(fmt.Sprintf("%d", getMask[:]))
}
