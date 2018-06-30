package main

import "fmt"

func main() {
	greet("Jane", "doe")
}

func greet(fname, lname string) {
	fmt.Println(fname, lname)
}
