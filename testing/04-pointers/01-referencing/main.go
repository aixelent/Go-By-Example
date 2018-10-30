package main

import "fmt"

func main()  {
	a := 42

	fmt.Println(a)
	fmt.Println(&a)		// adres do pamieci 'a'

	// wskaznik na int i przypisany do tego adres z wartoscia oryginalnej zmiennej 'a'
	// a = 3			// to nadpisze oryginalna wartosc zmiennej 'a' i wszystkich wskaznikow

	var b *int = &a

	fmt.Println(b)
}
