package main

import "fmt"

func Defer() {
	// defer: A defer statement defers the execution of a function until the surrounding function returns.
	// defer fmt.Println("world")

	// fmt.Println("hello")

	// defer fmt.Println("1")
	// defer fmt.Println("2")
	// defer fmt.Println("3")

	a := 5
	defer fmt.Println("deferred print:", a)
	a = 10
	fmt.Println("regular print:", a)
}
