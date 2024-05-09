package main

import (
    "fmt"
	"testsGo/prueba"
)

func main() {
    prueba.Saludar("Eder")

	Greet("Eder")

    resultado := prueba.Sumar(3, 5)
    fmt.Println("El resultado de la suma es:", resultado)

    fmt.Println("El valor de Pi es:", prueba.Pi)
}
