package main

import "fmt"

func main() {
	s := "foo"
	var i interface{}
	i = s

	fmt.Printf("i: %v\n", i)
}
