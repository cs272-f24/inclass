package main

import "fmt"

const (
	_        = iota
	KB int64 = 1 << (10 * iota)
	MB
	GB
)

func main() {
	fmt.Println("KB: ", KB)
	fmt.Println("MB: ", MB)
}
