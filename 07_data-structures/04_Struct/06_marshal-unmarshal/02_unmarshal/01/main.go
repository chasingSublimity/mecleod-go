package main

import (
	"encoding/json"
	"fmt"
)

// Person -- note
type Person struct {
	First string
	Last  string
	Age   int
}

func main() {
	var p1 Person
	// all default values
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	// conversion -- json string -> slice of bites
	bs := []byte(`{"First":"James", "Last":"Bond", "Age":20}`)
	json.Unmarshal(bs, &p1) // Unmarshal json into struct at memory address of p1

	fmt.Println("---------------")
	fmt.Println(p1.First)
	fmt.Println(p1.Last)
	fmt.Println(p1.Age)
	fmt.Printf("%T \n", p1)
}
