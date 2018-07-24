package main

import "fmt"

func main() {
	c := make(chan int)
	go func() {
		c <- 1
	}()
	fmt.Println(<-c)
}

/*
By placing the blocking code onto it's own goroutine,
the main thread is allowed to continue. The main thread
is blocked at L10 until it has a value to recieve
from the channel.
*/
