package main

type Car struct {
	Make string
	Model string
	Year int
}

type Engine struct {
	//...
}

func MakeEngine() *Engine {
	//..
}

func MakeCar(mak, model string, year int) *Car {
	car := Car{
		Make: mak,
		Model: model,
		Year: year,
		Engine: MakeEngine(),
		Transmission: 
	}

	// Return a pointer to a local variable
	// In go, this is a "stack escape" and is 
	// copied to the heap so we can return it safely
	return &car

	// If we provide initializers in the same order
	// as declared, we don't need to use the field name
	// return &Car{mak, model, year}
}

func main() {
	car := MakeCar("Ford", "F150", 2024)
	fmt.Printf("car: %v\n", car)	
}
