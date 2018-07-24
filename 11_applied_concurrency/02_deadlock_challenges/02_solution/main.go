package main

import "fmt"

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
}

// No deadlock here. We close the channel,
// signaling that no other values will be coming.
// Concurrently, we are pulling values off the channel
// and loggin them to the console.
