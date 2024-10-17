package main

import (
	"fmt"
	"reflect"
)

type Foo struct {
	Name   string
	Height int
}

func PrintTypes(f interface{}) {
	values := reflect.ValueOf(f)
	types := values.Type()

	for i := range values.NumField() {
		field := types.Field(i)
		value := values.Field(i)
		typeName := value.Type().Name()
		fmt.Printf("Field: %v, type: %v, value: %v\n", field.Name, typeName, value)
	}
}

func main() {
	f := Foo{"Steph Curry", 77}
	PrintTypes(f)
}
