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

	host := os.Args[1]
	addrs, err := net.LookupHost(host)
	if err != nil {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}

	fmt.Printf("%s\n", addrs)
}
