package main

import "fmt"

// Square -- bla
type Square struct {
	side float64
}

func (z Square) area() float64 { // Square implements shape interface b/c of this method
	return z.side * z.side
}

// Shape -- bla
type Shape interface {
	area() float64
}

func info(z Shape) { // expects Shape interface
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := Square{10}
	info(s)
}
