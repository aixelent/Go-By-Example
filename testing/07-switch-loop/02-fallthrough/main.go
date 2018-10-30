package main

import "fmt"

func main() {
	switch "Adrian" {
	case "Monika":
		fmt.Println("Hi Moni")
	case "Adrian":
		fmt.Println("Hi Adrian")
		fallthrough
	case "Ania":
		fmt.Println("Hi Ann")
	case "Jurek":
		fmt.Println("Hi Jurek")
	default:
		fmt.Println("How are you?")
	}
}
