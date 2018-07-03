package main

import "fmt"

func main() {
	//array, length, capacity
	mySlice := make([]int, 0, 5)

	fmt.Println("-------------")
	fmt.Println(mySlice)
	fmt.Println(len(mySlice))
	fmt.Println(cap(mySlice)) // capacity doubles when exceeded to a certain point, after which they scale in some smaller proportion
	fmt.Println("-------------")

	for i := 0; i < 80; i++ {
		mySlice = append(mySlice, i)
		fmt.Println("Len:", len(mySlice), "Capacity:", cap(mySlice), "Value ", mySlice[i])
	}
}
