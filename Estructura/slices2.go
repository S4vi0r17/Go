package main

import "fmt"

func main() {

	meses := []string{"Enero", "Febrero", "Marzo", "Abril", "Mayo", "Junio"}

	longitud := len(meses)
	capacidad := cap(meses)

	fmt.Printf("Longitud: %d, Capacidad: %d\n", longitud, capacidad)

	meses = append(meses, "Julio") //Si la estrcutura esta a tope, se genera una nueva estructura con el doble de capacidad

	longitud = len(meses)
	capacidad = cap(meses)

	fmt.Printf("Longitud: %d, Capacidad: %d\n", longitud, capacidad)
	fmt.Printf("Meses: %v\n", meses) //%v es un verbo para imprimir un slice
	fmt.Printf("La dirección de memoria del slice es: %p\n", meses) //%p es un verbo para imprimir la dirección de memoria de un slice
	
}