package main

import "fmt"

func main() {
	var salutation string
	var name string
	fmt.Print("Please enter your name: ")
	fmt.Scan(&salutation, &name)
	fmt.Println("Hello, ", salutation, name)
}
