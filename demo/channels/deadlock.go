package main

import "fmt"

func foo(ch chan int) {
	ch <- 5 // blocks
}

func main() {
	ch := make(chan int) // unbuffered
	foo(ch)
	x := <-ch // blocks
	fmt.Println(x)
}
