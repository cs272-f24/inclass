package main

import "fmt"

func foo(ch chan string) {
	ch <- "foo"
	ch <- "bar"
	close(ch)
}

func bar(ch chan string) {
	ch <- "one"
	ch <- "two"
	close(ch)
}

func main() {
	// make a buffered channel of up to 10 strings
	ch := make(chan string, 10)

	// Call goroutine to produce data on the channel
	go foo(ch)
	go bar(ch)

	// Range over the channel until the channel is closed
	for s := range ch {
		fmt.Println("s: ", s)
	}
}
