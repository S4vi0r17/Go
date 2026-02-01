package main

import (
	"errors"
	"fmt"
)

// Función básica
func saludar(nombre string) string {
	return "Hola " + nombre
}

// Múltiples retornos (patrón común en Go)
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

// Método (función asociada a un tipo)
type Servidor struct {
	Nombre string
	Puerto int
}

func (s Servidor) Iniciar() string {
	return fmt.Sprintf("Iniciando %s en puerto %d", s.Nombre, s.Puerto)
}

func main() {
	nombre := "Juan"
	fmt.Println(saludar(nombre))

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(resultado)
	}

	servidor := Servidor{
		Nombre: "Mi Servidor",
		Puerto: 8080,
	}
	fmt.Println(servidor.Iniciar())
}
