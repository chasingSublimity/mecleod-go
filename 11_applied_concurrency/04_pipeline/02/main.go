package main

import "fmt"

func main() {
	// set up pipeline and consume output
	for v := range sq(gen(2, 3)) {
		fmt.Println(v)
	}
}

func gen(nums ...int) <-chan int { // recieve only
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func sq(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
