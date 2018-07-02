package main

import "fmt"

func main() {
	name := "Todd"
	fmt.Println(name) // Todd

	changeMe(name)

	fmt.Println(name) // Todd
}

func changeMe(z string) { // no pointer. NB: strings are immutable
	fmt.Println(z) // Todd
	z = "Rocky"
	fmt.Println(z) // Rocky
}
