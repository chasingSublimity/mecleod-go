package main

import "fmt"

func main() {
	x := 43
	fmt.Println(x)
	{
		fmt.Println(x)
		y := "blabalbalbalbala"
		fmt.Println(y)
	}
	// fmt.Println(y) // errors, outside scope
}
