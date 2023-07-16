package main

import (
	"fmt"
)

func main() {

	var calificacion int

	fmt.Println("Ingresa tu calificacion: ")
	fmt.Scanln(&calificacion)

	calificacion = int(calificacion)

	/*
		switch {

		case calificacion == 10:
		  fmt.Println("Felicitaciones, tienes 10")

		case calificacion >= 7 && calificacion <= 9:
		  fmt.Println("Muy bien, buen trabajo")

		case calificacion >= 5 && calificacion < 7:
		  fmt.Println("Bien, pero puedes mejorar")

		default:
		  fmt.Println("Desaprobaste")
		}
	*/

	switch calificacion {

	case 10:
		fmt.Println("Felicitaciones, tienes 10")
	case 9, 8, 7:
		fmt.Println("Muy bien, buen trabajo")
	case 6, 5:
		fmt.Println("Bien, pero puedes mejorar")
	default:
		fmt.Println("Desaprobaste")
	}
}

//Ver el video 5