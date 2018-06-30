package main

import "fmt"

// Contact - explanatory comment
type Contact struct {
	greeting string
	name     string
}

// SwitchOnType - explanatory comment
func SwitchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case Contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Sager")

	var t = Contact{"Good to see you,", "Blake"}
	SwitchOnType(t)
}
