package main

import "fmt"

func main() {
	age := 44
	changeMe(age)
	fmt.Println(age) // 44
}

func changeMe(z int) { // value of age (44) is stored in variable z
	fmt.Println(z)
	z = 24 // value of z is changed, NOT age
}
