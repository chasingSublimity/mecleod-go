package main

import (
	"errors"
	"log"
)

// ErrNorthgateMath contains an error message
var ErrNorthgateMath = errors.New("norgate math: square root of negative number")

func main() {
	_, err := Sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

// Sqrt is a stub and needs no comment
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, ErrNorthgateMath
	}
	// implementation
	return 42, nil

}
