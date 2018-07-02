package main

import "fmt"

func main() {
	m := make([]string, 1, 25)

	fmt.Println(m) // [ ]
	changeMe(m)
	fmt.Println(m) // [Todd]
}

/*
	slices already contain references to memory addresses
	of an underlying data structure (array),
	so passing a mem address is not necessary for mutation.
	ie: slice is a _reference type_
*/
func changeMe(z []string) {
	z[0] = "Todd"
	fmt.Println(z) //Todd
}
