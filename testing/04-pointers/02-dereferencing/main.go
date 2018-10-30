package main

import "fmt"

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	var b *int = &a			// referencja

	fmt.Println(b)
	fmt.Println(*b)			// dereferencja czyli zamiast podawac adres w pamieci podajemy wartosc spod tego adresu
}