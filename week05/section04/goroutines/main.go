package main

import (
	"fmt"
	"time"
)

func counter(s string) {
	for i := range 10 {
		fmt.Printf("%s: %d\n", s, i)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	go counter("goroutine")
	counter("main goroutine")

	go func() {
		for i := range 10 {
			fmt.Printf("Anon %d\n", i)
		}
	}()

	time.Sleep(10 * time.Millisecond)
}
