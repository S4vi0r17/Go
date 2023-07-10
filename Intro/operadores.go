package main

import "fmt"

const ( 
	Domingo int = iota + 1 // 0 + 1 = 1
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

/*
	const Domingo = 0, Lunes = 1, Martes = 2, Miercoles = 3, Jueves = 4, Viernes = 5, Sabado = 6
*/

func main() {

	/*
		&& -> AND
		|| -> OR
		! -> NOT
	*/

	resultado := !(5 > 10) && 10 > 5 //!false -> true 
	fmt.Println(resultado)

	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)
}