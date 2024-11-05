package main

import (
	"bufio"
	"fmt"
	"os"
)

func GetAnswer(question string) string {
	return "Here's your answer"
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("\nWhat? >")
	for scanner.Scan() {
		question := scanner.Text()
		answer := GetAnswer(question)
		fmt.Println(answer)
		fmt.Print("\nWhat? >")
	}
}
