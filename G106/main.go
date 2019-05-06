package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}
	// godaddy.com:80
	args := os.Args[1]

	addr, err := net.ResolveTCPAddr("tcp4", args)
	//ifErr("Bad request:", err)
	if err != nil {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}

	tcpConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}
	_, err = tcpConn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Printf("Bad request: %s\n", os.Args[0])
		os.Exit(1)
	}

	b, err := ioutil.ReadAll(tcpConn)
	fmt.Println(string(b))
	os.Exit(0)
}
