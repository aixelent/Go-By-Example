package main

import "fmt"

func main() {
	str := "BLAKE2b_512BLAKE2b_384BLAKE2b_256BLAKE2s_256SHA512_256SHA512_224"
	newStr := make([]rune, len(str))
	i, added := 0, false

	for _, r := range str {
		if r >= '0' && r <= '9' {
			if added {
				continue
			}
			added, newStr[i] = true, 'x'
		} else {
			added, newStr[i] = false, r
		}
		i++
	}
	fmt.Println(string(newStr[:i]))
}
