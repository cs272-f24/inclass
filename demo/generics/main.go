package main

import "fmt"

func SumThings[K string, V int | float64](m map[K]V) V {
	var sum V
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	intmap := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	fmap := map[string]float64{
		"baz":  3.14159,
		"buzz": 2.71828,
	}
	sum := SumThings(intmap)
	fmt.Println("Sum of intmap: ", sum)

	fsum := SumThings(fmap)
	fmt.Println("Sum of fmap: ", fsum)
}
