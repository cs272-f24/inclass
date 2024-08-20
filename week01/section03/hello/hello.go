package main

import (
	"fmt"
)

// myFunc() returns a string
func myFunc() string {
	return "myfunc returns a string"
}

func main() {
	// Declare s of type string
	var s string = "Hello CS 272!"
	fmt.Println(s)

	var t string
	if true {
		t = "this"
	} else {
		t = "that"
	}
	fmt.Println(t)

	// Short declaration
	// type is inferred from right hand side (RHS)
	u := "Hello CS 272!"
	fmt.Println(u)

	v := myFunc()
	fmt.Println(v)
}
