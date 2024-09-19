package main

import (
	"fmt"
	"time"
)

func CountUp(s string) {
	for i := range 10 {
		fmt.Printf("%s: %d\n", s, i)
		time.Sleep(10 * time.Millisecond)
	}
}

func main() {
	go CountUp("from goroutine")
	CountUp("from main goroutine")

	go func() {
		for i := range 10 {
			fmt.Printf("anon: %d\n", i)
		}
	}()

	time.Sleep(1 * time.Microsecond)
}
