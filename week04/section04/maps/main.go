package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 1
	m["two"] = 2
	fmt.Println(m)

	m2 := map[string]string{}
	m2["foo"] = "bar"
	m2["baz"] = "something"
	fmt.Println(m2)

	// iterate over a map
	for key, val := range m2 {
		fmt.Printf("m2[%q] = %q\n", key, val)
	}

	s := m2["nope"]
	fmt.Printf("val: %s\n", s) // empty string due to zero-initialization

	k := "nope"
	_, ok := m2[k] // use a variable for key
	if ok {
		fmt.Printf("key %q exists\n", k)
	} else {
		fmt.Printf("key %q does not exist\n", k)
	}
}
