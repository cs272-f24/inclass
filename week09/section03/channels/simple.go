package main

import "fmt"

func foo(ch chan int) {
	// Block until written
	ch <- 99
}

func main() {
	// Create a unbuffered (synchronous) channel of integers
	ch := make(chan int)

	// Pass the channel to a goroutine
	go foo(ch)

	// Block the main goroutine to Read the channel and put the value into i
	i := <-ch

	// Useful example
	// quit := make(chan bool)
	// <- quit

	fmt.Println("i: ", i)
}
