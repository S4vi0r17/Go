package main

import "fmt"

func main() {
	var arreglo [5]int

	arreglo[0] = 1
	arreglo[1] = 2
	arreglo[2] = 3
	arreglo[3] = 4
	arreglo[4] = 5

	fmt.Println(arreglo)

	// Otra forma de declarar un arreglo
	arreglo2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(arreglo2)
}