package main

import "fmt"

func main() {
	sl := []int{}
	sl = append(sl, 99, 100, 101)

	for pos, val := range sl {
		fmt.Printf("sl[%d] = %d\n", pos, val)
	}
}
