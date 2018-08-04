package main

import (
	"fmt"
	"log"
)

// NorgateMathError comment bla
type NorgateMathError struct {
	lat, long string
	err       error
}

func (n *NorgateMathError) Error() string {
	return fmt.Sprintf("a norgate math error occured: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	_, err := Sqrt(-10.23)
	if err != nil {
		log.Println(err)
	}
}

// Sqrt is a stub and needs no comment
func Sqrt(f float64) (float64, error) {
	if f < 0 {
		nme := fmt.Errorf("norgate math again: %v", f)
		return 0, &NorgateMathError{"50.2289 N", "99.4656 W", nme}
	}
	// implementation
	return 42, nil

}
