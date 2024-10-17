package main

import "fmt"

type Dog struct {
	Name string
	Bark string
}

func Spot(ch chan Dog) {
	ch <- Dog{"Spot", "arf!"}
}

func Milo(ch chan Dog) {
	ch <- Dog{"Milo", "woof!"}
}

func main() {
	ch1 := make(chan Dog, 10)
	ch2 := make(chan Dog, 10)

	go Spot(ch1)
	go Milo(ch2)

outer:
	for {
		select {
		case d := <-ch1:
			fmt.Println("ch1 Dog: ", d)
		case d := <-ch2:
			fmt.Println("ch2 Dog: ", d)
		default:
			break outer
		}
	}
}
