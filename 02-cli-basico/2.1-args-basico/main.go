// Ejercicio 2.1: os.Args - Argumentos de línea de comandos
// Objetivo: Entender cómo Go recibe argumentos del usuario

package main

import (
	"fmt"
	"os"
	"strconv"
)

// ============================================
// PARTE 1: Qué es os.Args
// ============================================

// os.Args es un slice de strings que contiene:
// - os.Args[0]: nombre del programa
// - os.Args[1:]: argumentos pasados por el usuario
//
// Ejemplo: go run main.go deploy frontend --force
// os.Args = ["main.go", "deploy", "frontend", "--force"]

func demoArgsBasico() {
	fmt.Println("=== os.Args Básico ===")

	fmt.Printf("Programa: %s\n", os.Args[0])
	fmt.Printf("Total argumentos: %d\n", len(os.Args))

	if len(os.Args) > 1 {
		fmt.Println("\nArgumentos recibidos:")
		for i, arg := range os.Args[1:] {
			fmt.Printf("  [%d] %s\n", i+1, arg)
		}
	} else {
		fmt.Println("\nNo se pasaron argumentos")
		fmt.Println("Prueba: go run main.go hola mundo 123")
	}
}

// ============================================
// PARTE 2: Patrón de comandos
// ============================================

func demoComandos() {
	fmt.Println("\n=== Patrón de Comandos ===")

	if len(os.Args) < 2 {
		fmt.Println("Uso: programa <comando> [argumentos]")
		fmt.Println("\nComandos disponibles:")
		fmt.Println("  deploy  - Despliega un proyecto")
		fmt.Println("  status  - Muestra estado")
		fmt.Println("  help    - Muestra ayuda")
		return
	}

	comando := os.Args[1]

	switch comando {
	case "deploy":
		if len(os.Args) < 3 {
			fmt.Println("Uso: programa deploy <proyecto>")
			os.Exit(1)
		}
		proyecto := os.Args[2]
		fmt.Printf("Desplegando proyecto: %s\n", proyecto)

	case "status":
		fmt.Println("Estado: OK")

	case "help":
		fmt.Println("Ayuda del programa...")

	default:
		fmt.Printf("Comando desconocido: %s\n", comando)
		os.Exit(1)
	}
}

// ============================================
// PARTE 3: Calculadora simple con os.Args
// ============================================

func calculadora() {
	fmt.Println("\n=== Calculadora con os.Args ===")

	// Verificar argumentos mínimos
	if len(os.Args) < 4 {
		fmt.Println("Uso: calculadora <operación> <número1> <número2>")
		fmt.Println("\nOperaciones: add, sub, mul, div")
		fmt.Println("Ejemplo: go run main.go add 5 3")
		return
	}

	operacion := os.Args[1]

	// Parsear números
	a, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: '%s' no es un número válido\n", os.Args[2])
		os.Exit(1)
	}

	b, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: '%s' no es un número válido\n", os.Args[3])
		os.Exit(1)
	}

	var resultado float64

	switch operacion {
	case "add", "+":
		resultado = a + b
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, resultado)

	case "sub", "-":
		resultado = a - b
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, resultado)

	case "mul", "*", "x":
		resultado = a * b
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, resultado)

	case "div", "/":
		if b == 0 {
			fmt.Fprintln(os.Stderr, "Error: división por cero")
			os.Exit(1)
		}
		resultado = a / b
		fmt.Printf("%.2f / %.2f = %.2f\n", a, b, resultado)

	default:
		fmt.Fprintf(os.Stderr, "Operación desconocida: %s\n", operacion)
		os.Exit(1)
	}
}

// ============================================
// PARTE 4: Exit Codes
// ============================================

// Los exit codes comunican el resultado al sistema operativo:
// 0 - Éxito
// 1 - Error general
// 2 - Mal uso del comando (argumentos incorrectos)
// 126 - Comando no ejecutable
// 127 - Comando no encontrado
// 130 - Terminado con Ctrl+C

func demoExitCodes() {
	fmt.Println("\n=== Exit Codes ===")

	// Simular diferentes escenarios
	fmt.Println("Código 0: os.Exit(0) - Éxito")
	fmt.Println("Código 1: os.Exit(1) - Error general")
	fmt.Println("Código 2: os.Exit(2) - Argumentos incorrectos")

	// En un programa real:
	// if err != nil {
	//     fmt.Fprintln(os.Stderr, "Error:", err)
	//     os.Exit(1)
	// }
	// os.Exit(0) // o simplemente retornar de main()
}

// ============================================
// MAIN
// ============================================

func main() {
	// Descomentar la función que quieras probar:

	demoArgsBasico()
	// demoComandos()
	// calculadora()
	demoExitCodes()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea un programa que reciba un nombre como argumento
	// y salude: "Hola, <nombre>!"
	// Si no hay argumento, muestra error y sale con código 1

	// TODO 2: Crea un conversor de temperatura:
	// programa celsius 100  -> 100°C = 212°F
	// programa fahrenheit 32 -> 32°F = 0°C

	// TODO 3: Crea un programa que acepte múltiples archivos como argumentos
	// y muestre información de cada uno (nombre y si existe o no)
	// Usa os.Stat() para verificar existencia

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un CLI "mini-git" con los siguientes comandos:
	//
	// programa init           -> "Repositorio inicializado"
	// programa add <archivo>  -> "Archivo <archivo> agregado al staging"
	// programa add .          -> "Todos los archivos agregados"
	// programa commit <msg>   -> "Commit creado: <msg>"
	// programa status         -> "Rama: main, Sin cambios"
	// programa log            -> Muestra últimos 3 "commits" ficticios
	//
	// Requisitos:
	// - Validar que cada comando tenga los argumentos necesarios
	// - Mostrar ayuda si no se pasa comando
	// - Usar exit codes apropiados
	// - Errores van a stderr, output normal a stdout

	// Tu código aquí...
}
