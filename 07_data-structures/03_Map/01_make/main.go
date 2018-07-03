package main

import "fmt"

func main() {
	var myGreeting = make(map[string]string)
	myGreeting["Tim"] = "good morning."
	myGreeting["Jenny"] = "Bonjour"

	fmt.Println(myGreeting)
}
