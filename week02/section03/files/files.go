package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f, err := os.Create("/bin/🤣foo.txt")
	if err != nil {
		// something failed
		log.Fatalf("os.Create returned %v\n", err)
	}
	defer f.Close()

	_, err = f.Write([]byte("🤓foo"))
	if err != nil {
		fmt.Printf("f.Write returned %v\n", err)
	}
}
