package main

import (
	"fmt"
)

func main() {

	//Colección de datos no ordenados con clave-valor

	mapa := make(map[int]string)

	mapa[0] = "Domingo"
	mapa[1] = "Lunes"
	mapa[2] = "Martes"
	mapa[3] = "Miercoles"
	mapa[4] = "Jueves"
	mapa[5] = "Viernes"
	mapa[6] = "Sabado"
	
	mapa[4] = "Jueves 2"

	fmt.Println(mapa[4])

	//Eliminar un elemento del mapa
	delete(mapa, 4)

	fmt.Println(mapa)

	//usuarios := make(map[string] []int)

	//usuarios["Juan Perez"] = []int{10, 8, 9, 7, 10}

	//fmt.Println(usuarios)

	//Recorrer un mapa

	usuarios := map[int] string {}

	usuarios[1] = "Usuario 1"
	usuarios[2] = "Usuario 2"
	usuarios[3] = "Usuario 3"
	usuarios[4] = "Usuario 4"

	for id, username := range usuarios { //range devuelve dos valores, el indice y el valor

		fmt.Println(id, username)
	}

	
}
