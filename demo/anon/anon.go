package main

import "fmt"

func main() {
	var f = func() {
		fmt.Println("Hello from a function")
	}

	fmt.Println("Hello from main")
	f()
}
