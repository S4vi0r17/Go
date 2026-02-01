package main

import (
	"fmt"
)

func main() {
	var nombre string = "servidor"
	var puerto int = 8080

	// Constantes
	const VERSION = "1.0.0"

	// Declaración corta (solo dentro de funciones)
	activo := true
	conexiones := 10.5

	fmt.Printf("Nombre: %s (tipo: %T)\n", nombre, nombre)             // string
	fmt.Printf("Puerto: %d (tipo: %T)\n", puerto, puerto)             // int
	fmt.Printf("Version: %s (tipo: %T)\n", VERSION, VERSION)          // string
	fmt.Printf("Activo: %t (tipo: %T)\n", activo, activo)             // bool
	fmt.Printf("Conexiones: %f (tipo: %T)\n", conexiones, conexiones) // float64

	// Tipos básicos
	// string, int, int64, float64, bool, byte, rune

	// Para debugging usar %v,
	// Para salida user-facing: %s, %d, %.2f, etc.
}
