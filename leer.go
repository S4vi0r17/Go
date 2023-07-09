package main

import "fmt"

func main() {

	var nombre string // se usa %s para string
	var edad int // se usa %d para enteros
	var altura float32 // se usa %f para flotantes

	fmt.Println("Ingresa tu nombre: ")
	//fmt.Println("Ingresa una oracion: ")
	//fmt.Scanf("%s", &nombre) para scanf se usa print para leer el dato, no println porque no hace salto de línea
	fmt.Scanln(&nombre) //no lee espacio en blanco, los usa como separador de argumentos

	fmt.Println("Ingresa tu edad: ")
	fmt.Scanln(&edad)

	
	fmt.Println("Ingresa tu altura: ")
	fmt.Scanln(&altura)

	fmt.Println("El nombre ingresado es:", nombre)
											  // %f - %.2f solo para 2 decimales
	fmt.Printf("Hola %s, tienes %d años y mides %.2f metros\n", nombre, edad, altura)

}
