package _1

import (
	"fmt"
	"strings"
)

const words string = "Save time while PyCharm takes care of the routine. Focus on the bigger things and embrace the keyboard-centric approach to get the most of PyCharm’s many productivity features."

func main() {
	sep := strings.Fields(words)
	for index, word := range sep {
		fmt.Printf("Word %d: %s \n", index, word)
		//fmt.Printf("Word %s \n", word)
	}
}
