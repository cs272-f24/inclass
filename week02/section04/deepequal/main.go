package main

import (
	"fmt"
	"reflect"
)

type Car struct {
	Make  string
	Model string
}

func main() {
	c1 := Car{"Ford", "Mustang"}
	//c2 := Car{"Chevy", "Camaro"}
	c2 := Car{"Ford", "Mustang"}

	if reflect.DeepEqual(c1, c2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
