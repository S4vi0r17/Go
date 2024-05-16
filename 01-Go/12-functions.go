package main

import "fmt"

func Functions4() {
	x := func() {
		fmt.Println("Hello from anonymous function")
	}
	x()

	func(text string) {
		fmt.Println(text)
	}("Hello from anonymous function with parameter")
}
