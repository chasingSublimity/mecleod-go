package main

import "fmt"

func greet(name string) { // name is a parameter
	fmt.Println(name)
}

func main() {
	greet("Jane") // "Jane" is an argument
	greet("John") // "John" is an argument
}

// greet is delcared with a parameter and called with argument
