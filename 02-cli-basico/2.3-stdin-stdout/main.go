// Ejercicio 2.3: Stdin, Stdout, Stderr
// Objetivo: Dominar la entrada/salida estándar

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// ============================================
// PARTE 1: Stdout vs Stderr
// ============================================

// Stdout: salida normal del programa
// Stderr: errores y diagnósticos (no se mezcla en pipes)
//
// Ejemplo de uso:
// programa > output.txt 2> errors.txt
// programa 2>&1 | grep "algo"

func demoStdoutStderr() {
	fmt.Println("=== Stdout vs Stderr ===")

	// Salida normal (stdout)
	fmt.Println("Esto va a stdout")
	fmt.Fprintln(os.Stdout, "Esto también va a stdout")

	// Errores (stderr)
	fmt.Fprintln(os.Stderr, "Esto va a stderr (errores)")

	// fprintf permite más control
	fmt.Fprintf(os.Stdout, "Stdout: %s\n", "mensaje")
	fmt.Fprintf(os.Stderr, "Stderr: %s\n", "error")
}

// ============================================
// PARTE 2: Leer de Stdin
// ============================================

func demoStdinBasico() {
	fmt.Println("\n=== Stdin Básico ===")

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Ingresa tu nombre: ")
	nombre, _ := reader.ReadString('\n')
	nombre = strings.TrimSpace(nombre) // Quitar \n y espacios

	fmt.Printf("Hola, %s!\n", nombre)
}

// ============================================
// PARTE 3: Scanner para leer líneas
// ============================================

func demoScanner() {
	fmt.Println("\n=== Scanner ===")

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Escribe líneas (Ctrl+D o 'salir' para terminar):")

	lineNum := 1
	for scanner.Scan() {
		linea := scanner.Text()

		if linea == "salir" || linea == "exit" {
			break
		}

		fmt.Printf("[%d] Leído: %s\n", lineNum, linea)
		lineNum++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error leyendo:", err)
	}

	fmt.Println("Fin de la lectura")
}

// ============================================
// PARTE 4: Preguntas interactivas
// ============================================

func pregunta(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	respuesta, _ := reader.ReadString('\n')
	return strings.TrimSpace(respuesta)
}

func preguntaConDefault(reader *bufio.Reader, prompt, defaultVal string) string {
	fmt.Printf("%s [%s]: ", prompt, defaultVal)
	respuesta, _ := reader.ReadString('\n')
	respuesta = strings.TrimSpace(respuesta)
	if respuesta == "" {
		return defaultVal
	}
	return respuesta
}

func preguntaSiNo(reader *bufio.Reader, prompt string, defaultSi bool) bool {
	defaultStr := "s/N"
	if defaultSi {
		defaultStr = "S/n"
	}

	fmt.Printf("%s [%s]: ", prompt, defaultStr)
	respuesta, _ := reader.ReadString('\n')
	respuesta = strings.ToLower(strings.TrimSpace(respuesta))

	if respuesta == "" {
		return defaultSi
	}

	return respuesta == "s" || respuesta == "si" || respuesta == "y" || respuesta == "yes"
}

func demoPreguntas() {
	fmt.Println("\n=== Preguntas Interactivas ===")

	reader := bufio.NewReader(os.Stdin)

	nombre := pregunta(reader, "Nombre del proyecto: ")
	version := preguntaConDefault(reader, "Versión", "1.0.0")
	usarGit := preguntaSiNo(reader, "¿Inicializar Git?", true)
	privado := preguntaSiNo(reader, "¿Proyecto privado?", false)

	fmt.Println("\n--- Resumen ---")
	fmt.Printf("Proyecto: %s v%s\n", nombre, version)
	fmt.Printf("Git: %v\n", usarGit)
	fmt.Printf("Privado: %v\n", privado)
}

// ============================================
// PARTE 5: Leer de pipe
// ============================================

func demoPipe() {
	fmt.Println("\n=== Leer de Pipe ===")

	// Detectar si hay datos en stdin (pipe)
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// Hay datos en pipe
		fmt.Println("Leyendo de pipe...")

		scanner := bufio.NewScanner(os.Stdin)
		lineCount := 0
		wordCount := 0

		for scanner.Scan() {
			line := scanner.Text()
			lineCount++
			wordCount += len(strings.Fields(line))
			fmt.Println(line) // Echo de la línea
		}

		fmt.Fprintf(os.Stderr, "\n--- Estadísticas ---\n")
		fmt.Fprintf(os.Stderr, "Líneas: %d\n", lineCount)
		fmt.Fprintf(os.Stderr, "Palabras: %d\n", wordCount)
	} else {
		// Input interactivo
		fmt.Println("No hay pipe, modo interactivo")
		fmt.Println("Usa: echo 'texto' | go run main.go")
	}
}

// ============================================
// PARTE 6: Formulario completo
// ============================================

type Configuracion struct {
	Nombre       string
	Email        string
	Edad         int
	Framework    string
	Experiencia  string
	Newsletter   bool
}

func formularioCompleto() {
	fmt.Println("\n=== Formulario de Registro ===")
	fmt.Println()

	reader := bufio.NewReader(os.Stdin)
	var config Configuracion

	// Nombre (requerido)
	for {
		config.Nombre = pregunta(reader, "Nombre completo: ")
		if config.Nombre != "" {
			break
		}
		fmt.Fprintln(os.Stderr, "Error: El nombre es requerido")
	}

	// Email (requerido)
	for {
		config.Email = pregunta(reader, "Email: ")
		if strings.Contains(config.Email, "@") {
			break
		}
		fmt.Fprintln(os.Stderr, "Error: Email inválido")
	}

	// Edad (opcional, default 0)
	edadStr := preguntaConDefault(reader, "Edad", "0")
	fmt.Sscanf(edadStr, "%d", &config.Edad)

	// Framework favorito (lista de opciones)
	fmt.Println("\nFrameworks disponibles:")
	fmt.Println("  1. BubbleTea")
	fmt.Println("  2. Cobra")
	fmt.Println("  3. Chi")
	fmt.Println("  4. Gin")
	opcion := preguntaConDefault(reader, "Selecciona (1-4)", "1")

	frameworks := map[string]string{
		"1": "BubbleTea",
		"2": "Cobra",
		"3": "Chi",
		"4": "Gin",
	}
	if fw, ok := frameworks[opcion]; ok {
		config.Framework = fw
	} else {
		config.Framework = "Otro"
	}

	// Experiencia (múltiples líneas)
	fmt.Println("\nDescribe tu experiencia con Go (línea vacía para terminar):")
	var experiencia []string
	for {
		linea := pregunta(reader, "> ")
		if linea == "" {
			break
		}
		experiencia = append(experiencia, linea)
	}
	config.Experiencia = strings.Join(experiencia, "\n")

	// Newsletter
	config.Newsletter = preguntaSiNo(reader, "¿Suscribirse al newsletter?", true)

	// Mostrar resumen
	fmt.Println("\n╔══════════════════════════════════════╗")
	fmt.Println("║         Resumen de Registro          ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Printf("\nNombre:     %s\n", config.Nombre)
	fmt.Printf("Email:      %s\n", config.Email)
	if config.Edad > 0 {
		fmt.Printf("Edad:       %d\n", config.Edad)
	}
	fmt.Printf("Framework:  %s\n", config.Framework)
	if config.Experiencia != "" {
		fmt.Printf("Experiencia:\n%s\n", config.Experiencia)
	}
	fmt.Printf("Newsletter: %v\n", config.Newsletter)

	// Confirmar
	fmt.Println()
	if preguntaSiNo(reader, "¿Confirmar registro?", true) {
		fmt.Println("\nRegistro completado!")
	} else {
		fmt.Println("\nRegistro cancelado")
		os.Exit(0)
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	// Descomentar la función que quieras probar:

	demoStdoutStderr()
	// demoStdinBasico()
	// demoScanner()
	// demoPreguntas()
	// demoPipe()
	// formularioCompleto()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea un programa que lea líneas de stdin y las numere:
	// Entrada: "hola" -> Salida: "1: hola"
	// Usa: echo -e "a\nb\nc" | go run main.go

	// TODO 2: Crea un "filtro" que lea de stdin y solo imprima líneas
	// que contengan una palabra específica pasada como argumento:
	// cat archivo.txt | go run main.go "error"

	// TODO 3: Crea un programa interactivo que pregunte datos de un
	// servidor (host, puerto, usuario) y genere un archivo de config

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un CLI interactivo "setup wizard" para configurar un proyecto:
	//
	// 1. Pregunta nombre del proyecto (requerido)
	// 2. Pregunta descripción (opcional)
	// 3. Selecciona tipo de proyecto (1. CLI, 2. Web, 3. API, 4. TUI)
	// 4. Pregunta autor y email
	// 5. Lista de dependencias (ingresa hasta línea vacía)
	// 6. ¿Usar git? (S/n)
	// 7. ¿Crear README? (S/n)
	//
	// Al final:
	// - Muestra un resumen completo
	// - Pide confirmación
	// - Simula la creación de archivos (solo print)
	//
	// Bonus: Genera un go.mod y main.go reales

	// Tu código aquí...
}
