package main

import "fmt"

// User defined type, with underlying type of struct
type person struct { // struct is an aggregate type .
	first string // first, last, and age are all _fields_
	last  string
	age   int
}

func main() {
	p1 := person{"James", "Bond", 20} // assigned in order
	p2 := person{"Miss", "MoneyPenny", 18}
	fmt.Println(p1.first, p1.last, p1.age)
	fmt.Println(p2.first, p2.last, p2.age)
}
