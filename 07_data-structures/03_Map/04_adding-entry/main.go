package main

import "fmt"

func main() {
	myGreeting := map[string]string{
		"tim":   "Good Morning",
		"Jenny": "Bonjour",
	}

	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
}
