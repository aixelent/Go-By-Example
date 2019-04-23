package main

import (
	"fmt"
	"net"
)

func main() {
	ip := "127.0.0.1"
	fmt.Printf("IP: %s", net.ParseIP(ip))
}
