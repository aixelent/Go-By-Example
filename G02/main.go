package main

import "fmt"

func main()  {
	var col, lin int

	for lin = 1; lin <= 12; lin++ {
		for col = 1; col <= 12; col++ {
			fmt.Println(col, " x ", lin, " = ", col *lin)
		}
	}
}
