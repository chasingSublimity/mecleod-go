package main

import "fmt"

func main() {
	greeting := func() {
		fmt.Println("Hellow World!")
	}
	greeting()
	fmt.Printf("%T\n", greeting) // => func()
}
