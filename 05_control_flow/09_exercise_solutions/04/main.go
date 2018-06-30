package main

import "fmt"

func main() {
	var small int
	var big int
	fmt.Print("Enter a small number, followed by a larger number: ")
	fmt.Scan(&small, &big)
	fmt.Print(big % small)
}
