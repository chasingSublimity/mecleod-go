package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good Morning",
		1: "Bonjour",
		2: "Hiya",
		3: "LEEEEEROOOOYYY",
	}

	for key, val := range myGreeting {
		fmt.Println(key, "-", val)
	}
}
