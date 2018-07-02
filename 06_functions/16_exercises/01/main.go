package main

import "fmt"

func main() {
	fmt.Println(divideByTwo(4))
	fmt.Println(divideByTwo(5))
	fmt.Println(divideByTwo(6))
}

func divideByTwo(x int) (int, bool) {
	divided := x / 2
	isEven := x%2 == 0
	return divided, isEven
}
