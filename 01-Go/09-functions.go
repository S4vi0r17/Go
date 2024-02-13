package main

import (
	"fmt"
)

func main() {
	// Call the function
	sayHello("Gustavo")
	fmt.Println(sum(5, 5))
	fmt.Println(sumAndSub(5, 5))
	fmt.Println(sumAndSubNamed(5, 5))

	emoji := "🦎"
	fmt.Println(emoji)
	byReference(&emoji)
	fmt.Println(emoji)

	fmt.Println(sumAll(1, 2, 3, 4, 5))
}

// Function without return
func sayHello(name string) {
	fmt.Println("Hello", name)
}

// Function with return
func sum(a int, b int) int {
	return a + b
}

// Function with multiple return
func sumAndSub(a int, b int) (int, int) {
	return a + b, a - b
}

// Function with named return
func sumAndSubNamed(a int, b int) (sum int, sub int) {
	sum = a + b
	sub = a - b
	return
}

// Funcions by reference
func byReference(emoji *string) {
	*emoji = "👋"
}

// Function with variadic parameters
func sumAll(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
