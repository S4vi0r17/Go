package main

import "fmt"

//Los slices son como los arrays pero sin definir el tamaño y dinamicos

func main() {

	numeros := [] int {1, 2, 3, 4} //Referencia a un arreglo base

	numeros = append(numeros, 5)
	numeros = append(numeros, 6)
	numeros = append(numeros, 7)
	numeros = append(numeros, 8)
	numeros = append(numeros, 9)
	numeros = append(numeros, 10)

	nuevoSlice := numeros[0:5]

	numeros[0] = 100 //El slice es una referencia a un array, por lo que si modificamos el slice, modificamos el array
	nuevoSlice[1] = 200

	fmt.Println(numeros)
	fmt.Println(nuevoSlice)
}
