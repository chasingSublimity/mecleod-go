package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		fmt.Println("Err: ", err)
		// log.Println("Err: ", err)
		// log.Fatalln("Err: ", err) // calls os.exit after executing
		// panic(err) // displays stack
	}
}
