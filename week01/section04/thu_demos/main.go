package main

import "fmt"

func main() {

	isl := make([]int, 3)
	isl = append(isl, 2, 4, 6)
	fmt.Printf("isl contents: %v\n", isl)

	s := "Hello ✅"
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %c\n", i, s[i])
	}

	for pos, char := range s {
		fmt.Printf("s[%d] = %c\n", pos, char)
	}
}
