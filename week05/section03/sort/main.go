package main

import (
	"fmt"
	"sort"
)

type Person struct {
	name   string
	height int
}

type People []Person

func (ppl People) Len() int {
	return len(ppl)
}

func (ppl People) Less(i, j int) bool {
	return ppl[i].height < ppl[j].height
}

func (ppl People) Swap(i, j int) {
	// tmp := ppl[i]
	// ppl[i] = ppl[j]
	// ppl[j] = tmp

	ppl[i], ppl[j] = ppl[j], ppl[i]
}

func main() {
	people := People{
		{
			"Phil", 70,
		},
		{
			"StephC", 77,
		},
		{
			"DannyD", 58,
		},
	}

	fmt.Println("People: ", people)

	sort.Sort(people)

	fmt.Println("Sorted people: ", people)
}
