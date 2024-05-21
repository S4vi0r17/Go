package main

import (
	"fmt"
)

type exampler interface {
	Example()
}

func wrapper(i interface{}) {
	fmt.Printf("value: %v, type: %T\n", i, i)
}

func main() {
	// var i exampler
	// fmt.Printf("value: %v, type: %T\n", i, i) // value: <nil>, type: <nil>

	wrapper(42)
	wrapper("hello")
	wrapper(3.1416)
	wrapper(true)
}
