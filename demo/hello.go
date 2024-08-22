package main

import (
	"fmt"
)

func main() {
	s := "Hello CS 272! ❤️"
	for i := 0; i < len(s); i++ {
		fmt.Printf("s[%d] = %c\n", i, s[i])
	}
	//fmt.Println(os.Executable())
}
