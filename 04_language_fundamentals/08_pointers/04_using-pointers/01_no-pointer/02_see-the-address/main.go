package main

import "fmt"

func zero(x int) {
	fmt.Printf("%p\n", &x)
	x = 0
}

func main() {
	x := 5
	fmt.Printf("%p\n", &x)

	zero(x)

	fmt.Println(x)
}
