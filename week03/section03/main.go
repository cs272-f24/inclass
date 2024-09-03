package main

import "fmt"

type Car struct {
	Make  string
	Model string
	Year  int
}

func MakeCar(mak, model string, year int) *Car {
	// This creates a local variable of type Car
	// Returning address tells compiler to copy
	// local variable to a heap block
	// "stack escape"
	//return &Car{Make: mak, Model: model, Year: year}

	// We can also allocate heap blocks directly
	// using new()
	c := new(Car)
	c.Make = mak
	c.Model = model
	c.Year = year
	return c
}

// c is a method "receiver" and that's all we need
// to add a method to the Car struct
func (c *Car) Print() {
	fmt.Printf("car: %v\n", *c)
}

func main() {
	c := MakeCar("Ford", "Mustang", 1964)

	c.Print()
}
