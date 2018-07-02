package main

import "fmt"

func main() {
	half := func(x int) (int, bool) {
		divided := x / 2
		isEven := x%2 == 0
		return divided, isEven
	}
	fmt.Println(half(4))
	fmt.Println(half(6))
	fmt.Println(half(7))
}
