package main

import "fmt"

func main() {
	// [key]value
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	delete(m, "k2")
	fmt.Println("map:", m)

	_, ok := m["k2"]
	fmt.Println("ok?:", ok)

	v, ok := m["k1"]
	fmt.Println("ok?:", ok, "value: ", v)

	// declare and init in same line with composite literal syntax
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map:", n)
}

// Checking undefined values returns zero value of map value type
