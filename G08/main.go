package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	intss := []int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19, 21, 23, 25, 27, 29, 31}
	str := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k"}

	fmt.Println(intss, "\nCapacity: ", cap(intss))
	fmt.Println(str, "\nCapacity: ", cap(str))

	for i := range intss {
		num := intss[i]
		text := strconv.Itoa(num)
		str = append(str, text)
	}

	result := strings.Join(str, " ")
	fmt.Println(result)
}
