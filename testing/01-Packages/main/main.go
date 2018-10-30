package main

import (
	"fmt"
	"github.com/golang/example/stringutil"
	"Go-By-Example/Testing/01-Packages/stringsy"
)

/* Wolam metode Reverse z reverse.go, ktora wola metode reverseString z reverseTwo.go
   Wielkosc liter. Duza litera eksport poza pakiet, mala w obszarze pakietu
   Wiec z poziomu main nie mozna wywolac bezposrednio reverseString */
func main() {
	fmt.Println(stringutil.Reverse("litugnirts yb desreveR"))
	fmt.Println(stringsy.Reverse("olleH"))
	fmt.Println(stringsy.MyName)
}
