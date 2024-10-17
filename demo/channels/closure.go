package main

import "fmt"

func main() {
	ch := make(chan int) // unbuffered

	go func() {
		ch <- 5
	}()

	x := <-ch
	fmt.Println(x)
}
