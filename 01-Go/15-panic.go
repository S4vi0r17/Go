package main

import "fmt"

func main() {
	divide(10, 2)
	divide(10, 0)
}

func divide(a, b int) {
	validateDivisor(b)
	fmt.Println("Dividing", a, "by", b, "equals", a/b)
}

func validateDivisor(b int) {
	if b == 0 {
		panic("Cannot divide by zero")
	}
}
