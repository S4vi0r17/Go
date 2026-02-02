// Ejercicio 2: Funciones y Métodos
// Objetivo: Dominar funciones, retornos múltiples y métodos

package main

import (
	"errors"
	"fmt"
	"strings"
)

// ============================================
// PARTE 1: Funciones básicas
// ============================================

// Función sin parámetros ni retorno
func saludar() {
	fmt.Println("¡Hola!")
}

// Función con parámetros
func saludarA(nombre string) {
	fmt.Printf("¡Hola %s!\n", nombre)
}

// Función con retorno
func sumar(a, b int) int {
	return a + b
}

// Función con múltiples parámetros del mismo tipo
func concatenar(a, b, c string) string {
	return a + b + c
}

// ============================================
// PARTE 2: Retornos múltiples (MUY COMÚN EN GO)
// ============================================

// Retorno múltiple básico
func dividir(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

// Retornos nombrados
func estadisticas(numeros []int) (suma int, promedio float64, cantidad int) {
	cantidad = len(numeros)
	if cantidad == 0 {
		return 0, 0, 0
	}

	for _, n := range numeros {
		suma += n
	}
	promedio = float64(suma) / float64(cantidad)
	return // Retorna las variables nombradas
}

// ============================================
// PARTE 3: Funciones como valores
// ============================================

// Función que recibe otra función como parámetro
func aplicar(numeros []int, operacion func(int) int) []int {
	resultado := make([]int, len(numeros))
	for i, n := range numeros {
		resultado[i] = operacion(n)
	}
	return resultado
}

// ============================================
// PARTE 4: Métodos (funciones asociadas a tipos)
// ============================================

// Definir un tipo
type Servidor struct {
	Nombre string
	Puerto int
	Activo bool
}

// Método del tipo Servidor (receiver por valor)
func (s Servidor) Direccion() string {
	return fmt.Sprintf("%s:%d", s.Nombre, s.Puerto)
}

// Método que lee estado
func (s Servidor) Estado() string {
	if s.Activo {
		return "running"
	}
	return "stopped"
}

// Método que modifica estado (receiver por puntero)
func (s *Servidor) Iniciar() {
	s.Activo = true
	fmt.Printf("Servidor %s iniciado\n", s.Nombre)
}

func (s *Servidor) Detener() {
	s.Activo = false
	fmt.Printf("Servidor %s detenido\n", s.Nombre)
}

// ============================================
// MAIN
// ============================================

func main() {
	fmt.Println("=== Funciones Básicas ===")
	saludar()
	saludarA("Developer")
	fmt.Printf("5 + 3 = %d\n", sumar(5, 3))
	fmt.Printf("Concatenado: %s\n", concatenar("Hola", " ", "Mundo"))

	fmt.Println("\n=== Retornos Múltiples ===")

	// Patrón común: valor, error
	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", resultado)
	}

	// Manejando error de división por cero
	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Retornos múltiples con estadísticas
	nums := []int{10, 20, 30, 40, 50}
	suma, promedio, cantidad := estadisticas(nums)
	fmt.Printf("Números: %v\n", nums)
	fmt.Printf("Suma: %d, Promedio: %.2f, Cantidad: %d\n", suma, promedio, cantidad)

	fmt.Println("\n=== Funciones como Valores ===")

	numeros := []int{1, 2, 3, 4, 5}

	// Función anónima para duplicar
	duplicados := aplicar(numeros, func(n int) int {
		return n * 2
	})
	fmt.Printf("Original: %v\n", numeros)
	fmt.Printf("Duplicados: %v\n", duplicados)

	// Función anónima para elevar al cuadrado
	cuadrados := aplicar(numeros, func(n int) int {
		return n * n
	})
	fmt.Printf("Cuadrados: %v\n", cuadrados)

	fmt.Println("\n=== Métodos ===")

	servidor := Servidor{
		Nombre: "localhost",
		Puerto: 8080,
		Activo: false,
	}

	fmt.Printf("Servidor: %s\n", servidor.Direccion())
	fmt.Printf("Estado: %s\n", servidor.Estado())

	servidor.Iniciar()
	fmt.Printf("Estado: %s\n", servidor.Estado())

	servidor.Detener()
	fmt.Printf("Estado: %s\n", servidor.Estado())

	// ============================================
	// EJERCICIOS: Completa el código
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea una función 'esPar' que reciba un int y retorne bool
	// func esPar(n int) bool { ... }

	// TODO 2: Crea una función 'parsearComando' que reciba un string como
	// "deploy frontend" y retorne (comando string, argumento string, error)
	// Usa strings.Split() para separar
	// Si no tiene 2 partes, retorna error
	// func parsearComando(input string) (string, string, error) { ... }

	// TODO 3: Crea un tipo 'Proyecto' con campos: Nombre, Ruta, Tipo
	// Agrega un método 'Resumen' que retorne un string descriptivo

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un tipo 'Calculadora' con un campo 'resultado' (float64)
	// Agrega métodos:
	// - Sumar(n float64) que suma n al resultado
	// - Restar(n float64) que resta n del resultado
	// - Multiplicar(n float64) que multiplica el resultado por n
	// - Dividir(n float64) error que divide el resultado entre n (maneja división por cero)
	// - Resultado() float64 que retorna el resultado actual
	// - Reset() que pone el resultado en 0
	//
	// Ejemplo de uso:
	// calc := Calculadora{}
	// calc.Sumar(10)
	// calc.Multiplicar(2)
	// fmt.Println(calc.Resultado()) // 20

	// Tu código aquí...
}

// Función auxiliar que podrías necesitar
func ejemplo() {
	// strings.Split divide un string
	partes := strings.Split("deploy frontend", " ")
	fmt.Println(partes)      // ["deploy", "frontend"]
	fmt.Println(len(partes)) // 2
}
