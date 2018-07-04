package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book
	res, err := http.Get("http://www.gutenberg.org/files/2701/2701-0.txt")
	if err != nil {
		log.Fatal(err)
	}

	//scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close() // close right before main exits.

	// set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)
	// create slice to hold counts
	buckets := make([]int, 250) // slice of len 200 and cap of 200
	// advances token, returns false when scan is finished (end of text or err)
	for scanner.Scan() {
		n := HashBucket(scanner.Text())
		buckets[n]++
	}

	fmt.Println(buckets[65:123]) // a to z
}

// HashBucket -- converts words into ints
func HashBucket(word string) int {
	return int(word[0])
}
