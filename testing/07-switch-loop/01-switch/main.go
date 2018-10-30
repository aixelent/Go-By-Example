package main

import "fmt"

func main() {
	switch "Adrian" {
	case "Monika":
		fmt.Println("Hi Moni")
	case "Adrian":
		fmt.Println("Hi Adrian")
	case "Ania":
		fmt.Println("Hi Ann")
	default:
		fmt.Println("How are you?")
	}
}
