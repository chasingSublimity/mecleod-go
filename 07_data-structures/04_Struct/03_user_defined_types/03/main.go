package main

import "fmt"

// declaration
type person struct {
	first string // fields
	last  string // `json:"-"`            // json tag, dont include in json
	age   int    // `json:"wisdom score"` // json tag, rename field
}

func main() {
	p1 := person{"james", "bond", 20} // shorthand notation with struct literal
	p2 := person{"miss", "moneypenny", 20}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}

// struct is aggregate/composite type
