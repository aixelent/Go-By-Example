package main

import "fmt"

func main() {
	myFriendsName := "Mar"

	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Length 2")
	case len(myFriendsName) == 4:
		fmt.Println("Length 4")
	case len(myFriendsName) == 6:
		fmt.Println("Length 6")
	case len(myFriendsName) == 8:
		fmt.Println("Length 8")
	case len(myFriendsName) == 10:
		fmt.Println("Length 10")
	case myFriendsName == "Mar":
		fmt.Println("Length 13")
	default:
		fmt.Println("Nothing")
	}
}