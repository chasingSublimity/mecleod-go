package main

import "fmt"

func main() {
	c := gen(10)

	f := factorial(c)

	for n := range f {
		fmt.Println(n)
	}
}

func gen(n int) chan int {
	out := make(chan int)
	go func() {
		for i := 1; i < n; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func factorial(in chan int) chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- fact(n)
		}
		close(out)
	}()
	return out
}

func fact(n int) int {
	total := 1
	for i := n; i > 0; i-- {
		total *= i
	}
	return total
}
