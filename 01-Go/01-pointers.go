package main

import "fmt"

func main() {
	var x int
	var y *int

	x = 10
	y = &x

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("*y: ", *y)
	fmt.Printf("Type: %T, Value: %v, Dereferenced: %v\n", y, y, *y)

	*y = 20
	fmt.Println("\nx: ", x)
	fmt.Println("y: ", y)
	fmt.Println("*y: ", *y)
	fmt.Printf("Type: %T, Value: %v, Dereferenced: %v\n", y, y, *y)
}