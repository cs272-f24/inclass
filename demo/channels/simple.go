package main

import "fmt"

func foo(ch chan int) {
	ch <- 5
}

func main() {
	ch := make(chan int) // unbuffered
	go foo(ch)
	x := <-ch
	fmt.Println(x)
}
