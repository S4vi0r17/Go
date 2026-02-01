package main

import (
	"fmt"
)

func main() {
	// Array estático
	numeros := [3]int{1, 2, 3}
	fmt.Println(numeros)

	// Otra forma de declarar un array
	array2 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	// Slice (array dinámico)
	proyectos := []string{"frontend", "backend", "api"}
	proyectos = append(proyectos, "mobile")

	// Iterar
	for i, proyecto := range proyectos {
		fmt.Printf("%d: %s\n", i, proyecto)
	}

	// Slices are like arrays, but they can grow and shrink, also they don't data, they are just references(like pointers) to an array
	backend := proyectos[1:3]
	fmt.Println(backend)

	food := [5]string{"🍕", "🍓", "🍌", "🍑", "🍔"}
	fruits := food[1:4] // [🍓, 🍌, 🍑]

	fmt.Println("food: ", food)
	fmt.Println("fruits: ", fruits)
	fmt.Println("len(fruits): ", len(fruits))
	fmt.Println("cap(fruits): ", cap(fruits)) // from the first element in the slice to the last element in the array

	fruits = append(fruits, "🍇", "🍐", "🥑")
	// * If we append more elements than the capacity of the slice, the slice will be reallocated and the reference will change *

	fmt.Println("\nfood: ", food)
	fmt.Println("fruits: ", fruits)
	fmt.Println("len(fruits): ", len(fruits))
	fmt.Println("cap(fruits): ", cap(fruits))

	// Other way to declare a slice
	emoji := make([]string, 3)
	/*
		s := make([]int, 5) // Crea un slice de enteros con longitud 5 y capacidad 5
		s := make([]int, 5, 10) // Crea un slice de enteros con longitud 5 y capacidad 10
	*/
	emoji[0] = "🥵"
	emoji[1] = "🥶"
	emoji[2] = "🤮"
	emoji = append(emoji, "🤯", "🤬")
	fmt.Println(emoji)
	fmt.Println("len(emoji): ", len(emoji))
	fmt.Println("cap(emoji): ", cap(emoji))

	// Map (diccionario)
	estados := map[string]string{
		"frontend": "running",
		"backend":  "stopped",
		"api":      "running",
	}

	// Acceso
	estado, existe := estados["frontend"]
	if existe {
		fmt.Println(estado)
	} else {
		fmt.Println("No existe")
	}

	// Iterar map
	for nombre, estado := range estados {
		fmt.Printf("%s: %s\n", nombre, estado)
	}

	// Other way to declare a map
	var fruitsMap = make(map[string]string)
	fruitsMap["apple"] = "🍎"
	fruitsMap["banana"] = "🍌"
	fruitsMap["peach"] = "🍑"
	fmt.Println(fruitsMap)

	// Delete an element from a map
	delete(fruitsMap, "banana")
	fmt.Println(fruitsMap)
}
