package main

import "fmt"

func main() {

	/*
	var saludo string = "Hola, ¿cómo estás?"
	for i, c := range saludo {
		fmt.Printf("%d: %s byte: %d\n", i, string(c), c) // imprime el índice y el carácter
	}
	*/



	var sentence string
	fmt.Println("Ingresa una frase:")
	fmt.Scanln(&sentence)

	fmt.Println("Frase ingresada:", sentence)

	//fmt.Scanln(&oracion) // no lee espacios en blanco, los usa como separador de argumentos


}
