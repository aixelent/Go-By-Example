package main

import (
	"crypto/md5"
	"crypto/sha512"
	"fmt"
	"io/ioutil"
	"log"
)

var fileName = "file1.txt"

func readBytes() {
	f, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Printf("MD5: %x\n", md5.Sum(f))
	fmt.Printf("SHA512: %x\n", sha512.Sum512(f))
	fmt.Printf("SHA512_256: %x\n", sha512.Sum512_256(f))
}

func main() {
	readBytes()
}
