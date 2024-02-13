package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")

	if err != nil {
		fmt.Println("Failed to create the file")
		return
	}

	_, err = file.Write([]byte("Hello, World!"))

	defer file.Close()

	if err != nil {
		fmt.Println("Failed to write to the file")
		return
	}
}
