// Ejercicio 1: Variables y Tipos
// Objetivo: Familiarizarte con el sistema de tipos de Go

package main

import "fmt"

func main() {
	// ============================================
	// PARTE 1: Declaración de variables
	// ============================================

	// Declaración explícita con var
	var nombre string = "mi-servidor"
	var puerto int = 8080
	var activo bool = true
	var temperatura float64 = 36.5

	fmt.Println("=== Declaración Explícita ===")
	fmt.Printf("Nombre: %s (tipo: %T)\n", nombre, nombre)
	fmt.Printf("Puerto: %d (tipo: %T)\n", puerto, puerto)
	fmt.Printf("Activo: %v (tipo: %T)\n", activo, activo)
	fmt.Printf("Temperatura: %.1f (tipo: %T)\n", temperatura, temperatura)

	// ============================================
	// PARTE 2: Declaración corta con :=
	// ============================================

	// Solo funciona dentro de funciones
	servidor := "localhost"
	conexiones := 100
	ratio := 0.95

	fmt.Println("\n=== Declaración Corta ===")
	fmt.Printf("Servidor: %s\n", servidor)
	fmt.Printf("Conexiones: %d\n", conexiones)
	fmt.Printf("Ratio: %.2f\n", ratio)

	// ============================================
	// PARTE 3: Valores por defecto (zero values)
	// ============================================

	var textoVacio string  // ""
	var numeroVacio int    // 0
	var boolVacio bool     // false
	var floatVacio float64 // 0.0

	fmt.Println("\n=== Zero Values ===")
	fmt.Printf("string vacío: '%s'\n", textoVacio)
	fmt.Printf("int vacío: %d\n", numeroVacio)
	fmt.Printf("bool vacío: %v\n", boolVacio)
	fmt.Printf("float64 vacío: %f\n", floatVacio)

	// ============================================
	// PARTE 4: Constantes
	// ============================================

	const VERSION = "2.0.0"
	const MAX_RETRIES = 3
	const PI = 3.14159

	fmt.Println("\n=== Constantes ===")
	fmt.Printf("VERSION: %s\n", VERSION)
	fmt.Printf("MAX_RETRIES: %d\n", MAX_RETRIES)
	fmt.Printf("PI: %f\n", PI)

	// Las constantes NO pueden cambiar
	// VERSION = "3.0.0" // ERROR: cannot assign to VERSION

	// ============================================
	// PARTE 5: Conversión de tipos
	// ============================================

	var entero int = 42
	var flotante float64 = float64(entero) // Conversión explícita
	var texto string = fmt.Sprintf("%d", entero)

	fmt.Println("\n=== Conversión de Tipos ===")
	fmt.Printf("int a float64: %d -> %.1f\n", entero, flotante)
	fmt.Printf("int a string: %d -> '%s'\n", entero, texto)

	// ============================================
	// PARTE 6: Tipos básicos de Go
	// ============================================
	// string, int, int8, int16, int32, int64
	// uint, uint8, uint16, uint32, uint64
	// float32, float64
	// bool, byte (alias de uint8), rune (alias de int32)

	// ============================================
	// NOTAS de fmt
	// ============================================
	// %v  - valor en formato por defecto
	// %+v - incluye nombres de campos en structs
	// %#v - representación Go del valor
	// %T  - tipo del valor
	// %s  - string
	// %d  - entero base 10
	// %f  - float (%.2f para 2 decimales)
	// %t  - bool
	// %p  - puntero
	//
	// Para debugging usar %v o %+v
	// Para salida user-facing: %s, %d, %.2f, etc.

	// ============================================
	// EJERCICIO: Completa el código
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Declara una variable 'proyecto' de tipo string con valor "mi-app"
	// var proyecto ...

	// TODO 2: Declara una variable 'version' usando := con valor "1.0.0"
	// version := ...

	// TODO 3: Declara una variable 'maxConexiones' de tipo int con valor 1000
	// var maxConexiones ...

	// TODO 4: Declara una variable 'debugMode' de tipo bool con valor false
	// var debugMode ...

	// TODO 5: Imprime todas las variables que creaste
	// fmt.Printf(...)

	// ============================================
	// RETO FINAL
	// ============================================

	// Crea variables que representen la configuración de un servidor:
	// - hostname (string)
	// - port (int)
	// - useSSL (bool)
	// - timeout (float64) en segundos
	// - maxConnections (int)
	//
	// Luego imprime un resumen de la configuración

	fmt.Println("\n=== RETO: Configuración del Servidor ===")
	// Tu código aquí...
}
