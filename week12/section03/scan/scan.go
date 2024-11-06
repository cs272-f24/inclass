package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetAnswer(question string) string {
	return "here's your answer!"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nCatalog search> ")

	for scanner.Scan() {
		question := scanner.Text()
		answer := GetAnswer(question)
		fmt.Println(answer)
		fmt.Print("\nCatalog search> ")
	}
}
