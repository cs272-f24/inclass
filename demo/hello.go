package main

import (
	"fmt"
)

func main() {

	var arr [3]int
	printArray(arr)

	var arr2 [3]int
	printArray(arr2)

	sl := []int{}
	sl = append(sl, 1)
	fmt.Printf("s: %v\n", sl)

	s := "Hello CS 272! ❤️"
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %c\n", i, s[i])
	}
	//fmt.Println(os.Executable())
}
