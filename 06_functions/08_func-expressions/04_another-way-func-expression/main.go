package main

import "fmt"

func makeGreeter() func() string {
	return func() string { // returns anon func
		return "Hello!"
	}
}

func main() {
	greet := makeGreeter()
	fmt.Println(greet())
}
