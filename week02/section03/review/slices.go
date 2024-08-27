package main

import "fmt"

func main() {
	arr := [3]int{2, 4, 6}
	sl := arr[0:3] // sl could be a int*[] or int**
	arr[0] = 1     // expect to yield [1 4]
	//fmt.Println(sl)

	sl = append(sl, 8) // expect to yield [1 4 8]
	fmt.Println(sl)

	fmt.Println(arr)
}
