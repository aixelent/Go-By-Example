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

	args := os.Args[1]
	ip := net.ParseIP(args)
	if ip == nil {
		fmt.Println("Bad address format.")
	} else {
		fmt.Println("IP: ", ip.String())
	}
	os.Exit(0)
}
