package main

import "fmt"

func main()  {
	var num int
	var rev int = 0

	fmt.Println("Type numberic value: ")
	fmt.Scan(&num)
	temp := num

	for {
		rest := num % 10
		rev = rev * 10 + rest
		num /= 10

		if(num == 0) {
			break
		}
	}

	if(temp == rev) {
		fmt.Println("Value is a Palindrome", temp)
	} else {
		fmt.Println("Value isn't a Palindrome", temp)
	}
}
