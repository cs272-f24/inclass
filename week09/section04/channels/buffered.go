package main

import "fmt"

func main() {
	// ch is a buffered channel 10 strings wide
	ch := make(chan string, 10)

	// This anonymous goroutine closes over ch
	go func() {
		ch <- "foo"
		ch <- "bar"
		// Without the close() the main goroutine tries to read,
		// the range blocks, and we get a deadlock from the Go runtime
		// ranging over channels stops when the channel closes
		close(ch)

		// Trying to send on a closed channel will panic the runtime
		//ch <- "baz"
	}()

	// 	s1 := <-ch
	// 	fmt.Println("s1: ", s1)
	//
	// 	s2 := <-ch
	// 	fmt.Println("s2: ", s2)

	for s := range ch {
		fmt.Println("s: ", s)
	}

	fmt.Println("done with range loop")
}
