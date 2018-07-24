package main

import "fmt"

func main() {
	c := make(chan int)
	c <- 1 // this blocks, waiting for a reciever
	fmt.Println(<-c)
}

// Why does this result in a dead lock?
