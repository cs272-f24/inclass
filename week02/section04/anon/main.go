package main

import "fmt"

func main() {
	f := func() {
		fmt.Println("in anonymous function")
	}

	f()

	fmt.Println("in main function")
}
