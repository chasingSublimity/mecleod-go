package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"Tim":   "Good Morning",
		"Jenny": "Hiya",
	}
	fmt.Println(myGreeting["Jenny"])
}
