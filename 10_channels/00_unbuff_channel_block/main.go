package main

import (
	"fmt"
	"time"
)

func main() {
	// makes channel, over which we can communicate an int
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i // pass value into channel, block
		}
	}()

	go func() {
		for {
			fmt.Println(<-c) // take value from channel, unblock
		}
	}()

	time.Sleep(time.Second)
}
