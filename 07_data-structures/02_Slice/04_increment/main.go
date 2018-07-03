package main

import "fmt"

func main() {
	mySlice := make([]int, 2)
	fmt.Println(mySlice)
	fmt.Println(mySlice[0])
	mySlice[0] = 7
	fmt.Println(mySlice[0])
	mySlice[0]++
	fmt.Println(mySlice[0])
}
