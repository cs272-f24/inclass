package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create an unsorted slice of integers
	numbers := []int{3, 1, 4, 1, 5, 9, 2, 6, 5, 3}
	
	fmt.Println("Before sorting:", numbers)
	
	// Sort the slice in-place
	sort.Ints(numbers)
	
	fmt.Println("After sorting:", numbers)
}
