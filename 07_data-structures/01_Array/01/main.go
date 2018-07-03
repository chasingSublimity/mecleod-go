package main

import "fmt"

func main() {
	var x [58]string // brackets with number == array
	fmt.Println(x)
	fmt.Printf("%T \n", x) // length is part of the array
	fmt.Println(len(x))
	fmt.Println(x[42])
	for i := 65; i <= 122; i++ { // 65 in ascii table is A
		x[i-65] = string(i)
	}
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(x[42])
}

// arrays are not dynamic
