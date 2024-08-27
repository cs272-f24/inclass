package main

import "fmt"

func main() {
	arr := [3]int{200, 400, 600}
	sl := arr[0:1]

	fmt.Println(sl)
	arr[0] = 100
	fmt.Println(sl)
}
