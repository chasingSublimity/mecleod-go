package main

import "fmt"

type foo int

func main() {
	var myAge foo
	myAge = 44
	fmt.Printf("%T %v \n", myAge, myAge)

	var yourAge int
	yourAge = 29
	fmt.Printf("%T %v \n", yourAge, yourAge)

	// this doesn't work
	// fmt.Println(myAge + yourAge) // -> invalid operation: myAge + yourAge (mismatched types foo and int)

	// this does
	fmt.Println(int(myAge) + yourAge)

	// so does this
	fmt.Println(myAge + foo(yourAge))

}
