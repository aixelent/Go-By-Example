package main

import "fmt"

// makeGreeter zwraca funkcje anonimowa, ktora zwraca string
func makeGreeter() func() string {
	// zwracamy funkcje anonimowa, ktora zwraca string
	return func() string {
		return "Hello world"
	}
}

func main() {
	// makeGreeter przypisuje do funkcji greet i wyswietlam
	greet := makeGreeter()
	fmt.Println(greet())
}
