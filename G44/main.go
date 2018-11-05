package main

import (
	"fmt"
	"strings"
)

const words string = "Save time while PyCharm takes care of the routine. Focus on the bigger things and embrace the keyboard-centric approach to get the most of PyCharm’s many productivity features."

func main() {
	word := strings.NewReplacer("takes", "LoL", "most", "LoL", "PyCharm", "LoL")
	output := word.Replace(words)
	fmt.Println(output)
}
