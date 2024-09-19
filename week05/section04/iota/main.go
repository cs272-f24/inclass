package main

import "fmt"

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("PB: %d\n", PB)
}
