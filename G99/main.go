package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func main() {
	password := "pump312"

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Hash to store:", string(hash))

	// Comparing sum
	if err = bcrypt.CompareHashAndPassword(hash, []byte(password)); err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Password correct.")
}
