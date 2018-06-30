package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("https://www.facebook.com/")

	page, _ := ioutil.ReadAll(res.Body)

	res.Body.Close()
	fmt.Printf("%s", page)
}
