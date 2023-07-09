package main

import (
	"fmt"
	"reflect"
)

func main() {

	// Strings
	// var cadena string = "Hola Mundo"
	cadena := "Hola, probando strings"

	fmt.Println(cadena)
	fmt.Println(len(cadena)) // Longitud de la cadena

	// Caracteres
	// var caracter byte = 'a'
	caracter := 'a' // rune -> int32

	fmt.Println(caracter)
	fmt.Println(string(caracter)) // Convertir de byte a string
	fmt.Println(reflect.TypeOf(caracter))

	primerCaracter := cadena[0] //Char  -> uint8
	fmt.Println(primerCaracter)
	fmt.Println(string(primerCaracter))
	fmt.Println(reflect.TypeOf(primerCaracter))
	fmt.Printf("%c\n", primerCaracter) // Imprimir caracter

	ultimoCaracter := cadena[len(cadena)-1]
	fmt.Println(ultimoCaracter)
	fmt.Println(string(ultimoCaracter))
	fmt.Println(reflect.TypeOf(ultimoCaracter))
	fmt.Printf("%c\n", ultimoCaracter) // %c caracter
}