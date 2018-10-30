package main

import "fmt"

func main() {
	switch "Adrian" {
	case "Monika", "Adrian" :
		fmt.Println("Hi Moni")
	case "Adrian 2", "Kamil":
		fmt.Println("Hi Adrian")
	case "Ania", "Marta":
		fmt.Println("Hi Ann")
	case "Jurek", "Magda":
		fmt.Println("Hi Jurek")
	default:
		fmt.Println("How are you?")
	}
}
