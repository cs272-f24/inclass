package main

import "fmt"

func foo(ch chan int) {
	// Write a value into the channel
	ch <- 42
}

func main() {
	// Create an unbuffered/synchronous channel
	// So-called because we didn't provide a channel length to make()
	ch := make(chan int)

	// Call foo() in a goroutine
	go foo(ch)

	// Block on reading channel
	val := <-ch

	// useful example
	//<-quit

	fmt.Println("Val: ", val)

}
