package main

import "fmt"

func main() {
	s := "Hello world ✅"
	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %c\n", i, s[i])
	}

	for pos, char := range s {
		fmt.Printf("s[%d] = %c", pos, char)
	}
}
