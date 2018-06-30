package main

import "fmt"

func main() {
	switch "Marcus" {
	case "Tim":
		fmt.Println("Wassup Tim")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Marcus":
		fmt.Println("Wassup Marcus")
		fallthrough
	case "Jenny":
		fmt.Println("Wassup Jenny")
		fallthrough
	case "Blake":
		fmt.Println("Wassup Blake")
	default:
		fmt.Println("Have you no friends?")
	}
}
