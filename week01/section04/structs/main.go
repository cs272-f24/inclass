package main

import "fmt"

type Person struct {
	Name   string
	Height int
}

func PrintPerson(p *Person) {
	fmt.Printf("Name: %v\n", p.Name)
	fmt.Printf("Height: %d\n", p.Height)
}

func main() {
	p := Person{
		Name:   "Phil",
		Height: 70,
	}

	PrintPerson(&p)
	p2 := Person{
		"Greg", 71,
	}

	PrintPerson(&p2)
}
