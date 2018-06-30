package main

import "fmt"

func main() {
	fmt.Println(greet("Jane ", "Doe"))
}

/* syntax:
keyword [reciever] identifier(parameters) return value
*/
func greet(fname, lname string) string {
	return fmt.Sprint(fname, lname)
}
