package main

import (
	"fmt"
	"sort"
)

type AnInteger struct {
	x int
}

type Integers []AnInteger

func (ints Integers) Len() int {
	return len(ints)
}

func (ints Integers) Less(i, j int) bool {
	return ints[i].x < ints[j].x
}

func (ints Integers) Swap(i, j int) {
	// tmp := ints[i]
	// ints[i] = ints[j]
	// ints[j] = tmp

	ints[i], ints[j] = ints[j], ints[i]
}

func main() {
	ints := Integers{
		{1},
		{42},
		{12},
		{27},
	}

	fmt.Println("integers ", ints)

	sort.Sort(ints)

	fmt.Println("sorted integers: ", ints)
}
