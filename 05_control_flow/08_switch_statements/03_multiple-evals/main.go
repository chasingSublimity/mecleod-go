package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim, or, err, Jenny")
	case "Medhi", "Marcus":
		fmt.Println("Wassup Medhi")
	default:
		fmt.Println("Have you no friends?")
	}
}
