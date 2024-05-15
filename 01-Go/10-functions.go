package main

// import (
// 	"errors"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	content, err := os.ReadFile("file.txt")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println(content)
// 	fmt.Println(string(content))

// 	println()

// 	result, err := divide(100, 0)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}
// 	fmt.Println(result)
// }

// func divide(a, b float64) (float64, error) {
// 	if b == 0 {
// 		// return 0, fmt.Errorf("Cannot divide by zero")
// 		return 0, errors.New("Cannot divide by zero")
// 	}
// 	return a / b, nil
// }
