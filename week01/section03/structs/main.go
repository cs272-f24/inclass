package main

import "fmt"

type Person struct {
	Name   string
	Height int
}

func printPerson(p *Person) {
	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Height: %v\n", p.Height)
}

func makePerson(name string, height int) *Person {
	return &Person{name, height}
}

func main() {
	p := Person{
		Name:   "Phil",
		Height: 70,
	}
	printPerson(&p)

	p2 := Person{
		"Greg",
		71,
	}
	printPerson(&p2)
}
