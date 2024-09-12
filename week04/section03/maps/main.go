package main

import "fmt"

func foo() map[string]int {
	return map[string]int{
		"foo": 1,
	}
}

func main() {
	// Use make when creating a go type which requires data structure setup
	m := make(map[string]int)
	m["foo"] = 1
	fmt.Println(m)

	// Initialize to empty with {}
	m2 := map[string]string{}
	m2["bar"] = "baz"
	m2["this"] = "that"
	fmt.Println(m2)

	// Range over key, val. No val means just key
	for key, val := range m2 {
		fmt.Printf("m2[%q] = %q\n", key, val)
	}

	k := "nope"
	v, ok := m2[k]
	// OK tells us whether the map contains the key
	if ok {
		fmt.Println("m2 contains nope")
	} else {
		fmt.Println("m2 does not contain nope")
	}
	fmt.Printf("m2[%q] = %q\n", "nope", v)

	// map is a normal data type, here as return type
	fmt.Printf("from function: %v\n", foo())

	// Wow. Heterogeneous maps using go's any type
	m3 := map[string]any{}
	m3["something"] = 1
	m3["another"] = "huh"
	fmt.Printf("m3: %v\n", m3)
}
