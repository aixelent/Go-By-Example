package main

import (
	"fmt"
	"net"
)

var ip = "140.82.118.4"
var domain = "http://github.com"

func byIP() {
	// by address
	ips, err := net.LookupIP(ip)
	if err != nil {
		fmt.Println(err)
	}
	for _, ip := range ips {
		fmt.Println(ip.String())
	}
}

func byDomain() {
	addr, err := net.LookupAddr(ip)
	if err != nil {
		fmt.Println(err)
	}
	for _, ip := range addr {
		fmt.Println(ip)
	}
	//return err
}

func byHost() {
	addr, err := net.LookupHost(domain)
	if err, ok := err.(*net.DNSError); ok {
		if err.Timeout() {
			fmt.Println("Timed out")
		} else if err.Temporary() {
			fmt.Println("Temporary")
		} else {
			fmt.Println("Generic error", err)
		}
		return
	}
	fmt.Println(addr)
}

func main() {
	fmt.Println("By IP:\n")
	byIP()
	fmt.Println("By domain:\n")
	byDomain()
	fmt.Println("By host:")
	byHost()
}
