package main

import "fmt"

func main() {
	myFriendsName := "Mar"

	switch { // no expression
	case len(myFriendsName) == 2: // checks for true
		fmt.Println("Wassyp my friend with name length of 2")
	default:
		fmt.Println("Default")
	}

}
