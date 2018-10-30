package main

import (
	"fmt"
)

func main() {
	b := true

	if b {
		fmt.Println("Like this or like this...")
	}

	if food := "Chocolate"; b {
		fmt.Println(food)
	}

	//if err := file.Chmod(0664); err != nil {
	//	log.Print(err)
	//	return err
	//}
}
