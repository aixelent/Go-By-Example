package main

import (
	"fmt"
	"strings"
)

func main() {
	// s1 == s2 - result 0
	// s1 < s2 - result -1
	// s1 > s2 - result 1
	fmt.Println(strings.Compare("first", "first"))
	fmt.Println(strings.Compare("other", "another"))
	fmt.Println(strings.Compare("Maci", "maci"))
}
