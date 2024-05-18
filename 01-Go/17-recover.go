package main

import "fmt"

func Recover() {
	divideInt(10, 2)
	divideInt(10, 0)
}

func divideInt(a, b int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from", r)
		}
	}()
	validateDivisor(b)
	fmt.Println("Dividing", a, "by", b, "equals", a/b)
}
