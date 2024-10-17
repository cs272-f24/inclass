package main

import "fmt"

type Person struct {
	Name   string
	Height int
}

func AddPhil(ch chan Person) {
	ch <- Person{
		"Phil",
		69,
	}
}

func AddSteph(ch chan Person) {
	ch <- Person{
		"Steph",
		75,
	}
}

func main() {
	ch1 := make(chan Person, 10000)
	ch2 := make(chan Person, 10)

	go AddPhil(ch1)
	go AddSteph(ch2)

outer:

	for {
		select {
		case p1 := <-ch1:
			fmt.Println("Person on ch1: ", p1)
		case p2 := <-ch2:
			fmt.Println("Person on ch2: ", p2)
		default:
			break outer
		}
	}
}
