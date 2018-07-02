package main

import "fmt"

func main() {
	fmt.Println(findLargest(5, 6, 8, 1000000))
}

func findLargest(ints ...int) int {
	// ints is a slice of all the arguments
	var largest int
	for _, v := range ints {
		if v > largest {
			largest = v
		}
	}
	return largest
}
