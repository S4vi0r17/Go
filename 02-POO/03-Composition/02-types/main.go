package main

import (
	"fmt"
)

type course struct {
	name string
}

func (c course) Print() {
	fmt.Printf("%+v\n", c)
}

// Declaration of alias
type myAlias = course

// Definition of type
type myType course

type newBool bool

func (b newBool) String() string {
	if b {
		return "VERDADERO"
	}
	return "FALSO"
}

func main() {
	c := myType{name: "Go"}
	// c.Print()
	fmt.Printf("El tipo es: %T\n", c)

	var b newBool = true
	// fmt.Println(b) Overwrite the method String
	fmt.Println(b.String())
}
