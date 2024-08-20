package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var p Person
	p.Name = "Phil"
	p.Age = 100

	p2 := Person{"Jet", 20}

	fmt.Printf("p2: %v\n", p)
}
