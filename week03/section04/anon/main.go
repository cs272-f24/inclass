package main

import "fmt"

func main() {

	f := func() func() int {
		i := 0

		return func() int {
			i++
			return i
		}
	}

	g := f()

	fmt.Println(g())
	fmt.Println(g())
}
