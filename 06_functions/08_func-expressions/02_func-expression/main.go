package main

import "fmt"

func main() {

	// expressions are the only way to have nested functions
	greeting := func() { // func expression
		fmt.Println("Hello world")
	}

	greeting()
}
