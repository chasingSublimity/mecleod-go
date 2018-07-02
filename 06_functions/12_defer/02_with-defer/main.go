package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}

func world() {
	fmt.Println("world")
}

func main() {
	defer world() // pushes to end of main. (Call stack)?
	hello()
}
