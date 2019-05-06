package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

func checkError(message string, err error) {
	if err != nil {
		log.Fatalln("", err)
	}
}

func HandleRequest(conn net.Conn) {
	msg, err := bufio.NewReader(conn).ReadString('\n')
	checkError("Error reading", err)
	fmt.Print("Message received from: ", string(msg))
	conn.Close()
}

func main() {
	listener, err := net.Listen(TYPE, HOST+":"+PORT)
	checkError("Error starting server: ", err)
	defer listener.Close()

	log.Println("Listening:" + HOST + ":" + PORT)
	for {
		conn, err := listener.Accept()
		checkError("Error accepting: ", err)
		go HandleRequest(conn)
	}
}
