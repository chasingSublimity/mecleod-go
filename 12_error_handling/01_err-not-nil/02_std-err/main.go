package main

import (
	"fmt"
	"log"
	"os"
)

func init() {
	nf, err := os.Create("log.txt") // create file
	if err != nil {
		fmt.Println(err)
	}
	log.SetOutput(nf) // set output
}

func main() {
	_, err := os.Open("no-file.txt")
	if err != nil {
		log.Println("err: ", err)
	}
}
