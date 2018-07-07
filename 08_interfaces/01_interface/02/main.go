package main

import (
	"fmt"
	"math"
)

// Square -- bla
type Square struct {
	side float64
}

// Shape -- bla
type Shape interface {
	area() float64 // anything with this method implements the Shape interface
}

// Circle - bla
type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	return s.side * s.side
}

// Triangle - Bla
type Triangle struct {
	height float64
	base   float64
}

func (t Triangle) area() float64 {
	return (t.height * t.base) / 2
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func main() {
	s := Square{10}
	c := Circle{10}
	t := Triangle{15, 10}
	info(s)
	info(c)
	info(t)
}
