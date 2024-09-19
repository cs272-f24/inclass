package main

import "fmt"

func SumNumbers[K string, V int | float64](m map[K]V) V {
	var sum V
	sum = 0
	for _, v := range m {
		sum += v
	}
	return sum
}

func main() {
	m := map[string]int{
		"foo": 1,
		"bar": 2,
	}

	sum := SumNumbers(m)
	fmt.Println("sum of ints ", sum)

	n := map[string]float64{
		"pi": 3.14159,
		"e":  2.71828,
	}

	fsum := SumNumbers(n)
	fmt.Println("sum of floats ", fsum)
}
