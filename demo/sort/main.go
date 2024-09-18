package main

import (
	"fmt"
	"sort"
)

type Thingy struct {
	x int
}

type Thingies []Thingy

func (t Thingies) Len() int {
	return len(t)
}

func (t Thingies) Less(i, j int) bool {
	return t[i].x < t[j].x
}

func (t Thingies) Swap(i, j int) {
	// tmp := t[i]
	// t[i] = t[j]
	// t[j] = tmp

	t[i], t[j] = t[j], t[i]
}

func main() {
	thingies := Thingies{
		{10},
		{12},
		{3},
		{16},
	}

	fmt.Printf("thingies: %v\n", thingies)

	sort.Sort(thingies)

	fmt.Printf("sorted thingies: %v\n", thingies)
}
