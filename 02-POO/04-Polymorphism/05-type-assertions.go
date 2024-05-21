package main

import (
	"fmt"
	"strings"
)

type exampler interface {
	Example()
}

func wrapper(i interface{}) {
	fmt.Printf("value: %v, type: %T\n", i, i)
	v, ok := i.(string)
	if ok {
		fmt.Printf("\t%s\n", strings.ToUpper(v))
	}
}

func main() {
	// var i exampler
	// fmt.Printf("value: %v, type: %T\n", i, i) // value: <nil>, type: <nil>

	wrapper(42)
	wrapper("hello")
	wrapper(3.1416)
	wrapper(true)
}
