package main

import "fmt"

func main() {
	ch := make(chan int, 2) // buffered

	ch <- 1
	ch <- 2

	x := <-ch // in order
	fmt.Println(x)

	y := <-ch
	fmt.Println(y)
}
