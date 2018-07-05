package main

import "fmt"

// Person - helpful comment here
type Person struct {
	First string
	Last  string
	Age   int
}

// DoubleZero - helpful comment here
type DoubleZero struct {
	Person               // embedded type
	First         string // overwrites First of innner type (Person)
	LicenseToKill bool
}

func main() {
	p1 := DoubleZero{
		Person: Person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}
	fmt.Println(p1.First) // -> "Double Zero Seven"
}
