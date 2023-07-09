package main

import (
	"fmt"
	"strconv"
)

func main() {

	// go build -> un archivo ejecutable
	// go run main.go lo ejecuta de una

	//fmt.Println("Backend is running...")

	var nombre string
	var edad int

	nombre = "Jorge"
	edad = 22

	fmt.Println("Tu nombre: "+ nombre)
	fmt.Println("Tu edad: " + strconv.Itoa(edad)) //strconv.Itoa convierte un int a string
}
