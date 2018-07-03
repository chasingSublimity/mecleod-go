package main

import "fmt"

func main() {
	var student []string
	var students [][]string
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil) // true
}

// student is nil because underlying array hasnt been made,
// so address is pointing to nothing
