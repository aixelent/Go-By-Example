package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}

	domain := os.Args[1]
	addr, err := net.ResolveIPAddr("ip", domain)
	if err != nil {
		err.Error()
		os.Exit(1)
	}

	fmt.Printf("Address: %s\n", addr)
	os.Exit(0)
}
