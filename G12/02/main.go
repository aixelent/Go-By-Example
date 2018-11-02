package main

import (
	"fmt"

	"Go-By-Example/G12/02/stringsx"

	"github.com/golang/example/stringutil"
)

func main() {
	fmt.Println(stringutil.Reverse("olleH"))
	fmt.Println(stringsx.Reverse("Hello"))
	fmt.Println(stringsx.Reverse(stringsx.Name))
}
