package main

import "fmt"

// return function that returns an int
func wrapper() func() int {
	x := 0
	return func() int { // returned function
		x++
		return x
	}
}

func main() {
	increment := wrapper() // increment is now equal to the function returned by wrapper
	fmt.Println(increment())
	fmt.Println(increment())
}
