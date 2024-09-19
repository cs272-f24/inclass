package main

import "fmt"

const (
	_ = uint64(1 << (10 * iota))
	KB
	MB
	GB
	TB
	PB
)

const (
	Pi = 3.14159
	E  = 2.71828
)

func main() {
	fmt.Printf("KB: %d\n", KB)
	fmt.Printf("PB: %d\n", PB)
}
