package main

import "fmt"

func main() {
	s := "foo"
	var intf interface{}

	intf = s
	fmt.Println("intf: ", intf)
	switch intf.(type) {
	case string:
		fmt.Println("intf is a string")
	default:
		fmt.Println("intf is not a string")
	}
	s2 := intf.(string)
	fmt.Println("s2:", s2)

	// PANIC i := intf.(int)
	//fmt.Println("i: ", i)

}
