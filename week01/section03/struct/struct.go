package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	p.Name = "Phil"
	p.Age = 57

	p2 := Person{
		"Sanjana",
		20,
	}

	fmt.Printf("p: %v\n", p)
	fmt.Printf("p2: %v\n", p2)
}
