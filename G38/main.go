package main

import "fmt"

func main() {
	langs := make(map[int]string)
	langs[0] = "Cobol"
	langs[1] = "Go"
	langs[2] = "Java"
	langs[3] = "Python"
	langs[4] = "Rust"
	fmt.Println("Legth:", len(langs))
	fmt.Println(langs)

	delete(langs, 0)

	fmt.Println("Legth:", len(langs))
	fmt.Println(langs)
}
