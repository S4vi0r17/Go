package main

import "fmt"

func Panic() {
	div(10, 2)
	div(10, 0)
	div(20, 4)
}

func div(a, b int) {
	validateDivisor(b)
	fmt.Println("Dividing", a, "by", b, "equals", a/b)
}

func validateDivisor(b int) {
	if b == 0 {
		panic("Cannot divide by zero")
	}
}
