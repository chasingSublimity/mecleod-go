package main

import (
	"encoding/json"
	"fmt"
)

// Person - Exported
type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"James", "bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(string(bs))
}
