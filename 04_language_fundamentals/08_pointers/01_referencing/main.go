package main

import (
	"fmt"
)

func main() {
	a := 43

	fmt.Println(a)
	fmt.Println(&a)

	// type pointer to an int
	var b = &a

	fmt.Println(b)
}
