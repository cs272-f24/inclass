package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}
		close(ch)
	}()

	for x := range ch {
		fmt.Println(x)
	}

	fmt.Println("channel closed, exiting")
}
