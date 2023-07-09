package main

import "fmt"

//Constantes fuera creo
const pi = 3.14 

func main() {

	//var nombre, apellido, pais = "Eduardo", "Garcia", "Perú"
	nombre, apellido, pais := "Eduardo", "Garcia", "Perú"

	edad , altura := 21, 1.75

	var nickName string = "Ryuk"

	fmt.Println(nombre, apellido, pais)
	fmt.Println(edad, altura)
	fmt.Println(nickName)

	fmt.Println("Mi nombre es", nombre, apellido, "y soy de", pais)

	//Uso de constantes

	println(pi)
}
