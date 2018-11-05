package G41

import (
	"fmt"
	"strings"
)

const string = "Save time while PyCharm takes care of the routine. Focus on the bigger things and embrace the keyboard-centric approach to get the most of PyCharm’s many productivity features."

func main() {
	keyword := "Python"
	isIt := strings.Contains(string, keyword)
	fmt.Printf("\n\nThe \"%s\" contains \"%s\": %t \n", string, keyword, isIt)

	keyword = "PyCharm"
	isIt = strings.Contains(string, keyword)
	fmt.Printf("\n\nThe \"%s\" contains \"%s\": %t \n", string, keyword, isIt)
}
