package main

import "fmt"

func main() {
	// The data type of s is inferred from the type of
	// the expression on the right-hand-side (RHS)
	s := "Hello CS 272"
	fmt.Println(s)

	var t string = "Another string"
	fmt.Println(t)

	// An example with an expression
	i := 1 + 2
	i = i + 3
	fmt.Println(i)

	var j int = 3 + 4

	// Refer to the j which was declared above
	j = j + 4
	fmt.Println(j)
}
