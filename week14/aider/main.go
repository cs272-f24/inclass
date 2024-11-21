package main

import (
	"fmt"
	"sort"
)

func main() {
	// Create an unsorted slice of integers
	numbers := []int{64, 34, 25, 12, 22, 11, 90}
	
	fmt.Println("Before sorting:", numbers)
	
	// Sort the slice
	sort.Ints(numbers)
	
	fmt.Println("After sorting:", numbers)
}
