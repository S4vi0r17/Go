// Ejercicio 2.2: Package flag - Parsing de flags
// Objetivo: Dominar el package flag para opciones de línea de comandos

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ============================================
// PARTE 1: Flags básicos
// ============================================

func demoFlagsBasicos() {
	fmt.Println("=== Flags Básicos ===")

	// Definir flags (devuelven punteros)
	nombre := flag.String("nombre", "proyecto", "nombre del proyecto")
	puerto := flag.Int("puerto", 8080, "puerto del servidor")
	verbose := flag.Bool("verbose", false, "modo verbose")

	// IMPORTANTE: Parsear DESPUÉS de definir todos los flags
	flag.Parse()

	// Usar valores (dereferencia con *)
	fmt.Printf("Nombre: %s\n", *nombre)
	fmt.Printf("Puerto: %d\n", *puerto)
	fmt.Printf("Verbose: %v\n", *verbose)

	// Argumentos restantes (no-flags)
	fmt.Printf("Args extras: %v\n", flag.Args())
	fmt.Printf("Cantidad de extras: %d\n", flag.NArg())
}

// ============================================
// PARTE 2: Flags con variables existentes
// ============================================

func demoFlagsConVariables() {
	fmt.Println("\n=== Flags con Variables ===")

	var host string
	var timeout int
	var debug bool

	// Enlazar flags a variables existentes
	flag.StringVar(&host, "host", "localhost", "dirección del servidor")
	flag.IntVar(&timeout, "timeout", 30, "timeout en segundos")
	flag.BoolVar(&debug, "debug", false, "modo debug")

	flag.Parse()

	fmt.Printf("Host: %s\n", host)
	fmt.Printf("Timeout: %d\n", timeout)
	fmt.Printf("Debug: %v\n", debug)
}

// ============================================
// PARTE 3: Personalizar mensaje de uso
// ============================================

func demoUsagePersonalizado() {
	fmt.Println("\n=== Usage Personalizado ===")

	nombre := flag.String("nombre", "app", "nombre de la aplicación")
	env := flag.String("env", "development", "entorno de ejecución")

	// Personalizar el mensaje de ayuda
	flag.Usage = func() {
		fmt.Println("Mi Aplicación - v1.0.0")
		fmt.Println()
		fmt.Println("Uso: programa [opciones]")
		fmt.Println()
		fmt.Println("Opciones:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Ejemplos:")
		fmt.Println("  programa -nombre=frontend -env=production")
		fmt.Println("  programa -env=staging")
	}

	flag.Parse()

	fmt.Printf("Configurado: %s en %s\n", *nombre, *env)
}

// ============================================
// PARTE 4: Validación de flags
// ============================================

func demoValidacion() {
	fmt.Println("\n=== Validación de Flags ===")

	env := flag.String("env", "development", "entorno")
	puerto := flag.Int("puerto", 8080, "puerto")

	flag.Parse()

	// Validar entorno
	validEnvs := []string{"development", "staging", "production"}
	envValido := false
	for _, e := range validEnvs {
		if *env == e {
			envValido = true
			break
		}
	}
	if !envValido {
		fmt.Fprintf(os.Stderr, "Error: entorno inválido '%s'\n", *env)
		fmt.Fprintf(os.Stderr, "Válidos: %s\n", strings.Join(validEnvs, ", "))
		os.Exit(1)
	}

	// Validar puerto
	if *puerto < 1 || *puerto > 65535 {
		fmt.Fprintf(os.Stderr, "Error: puerto debe estar entre 1 y 65535\n")
		os.Exit(1)
	}

	fmt.Printf("Configuración válida: %s:%d\n", *env, *puerto)
}

// ============================================
// PARTE 5: FlagSet para subcomandos
// ============================================

func demoFlagSetSubcomandos() {
	fmt.Println("\n=== FlagSet para Subcomandos ===")

	// Cada subcomando tiene su propio set de flags
	deployCmd := flag.NewFlagSet("deploy", flag.ExitOnError)
	deployBranch := deployCmd.String("branch", "main", "rama a desplegar")
	deployForce := deployCmd.Bool("force", false, "forzar despliegue")

	statusCmd := flag.NewFlagSet("status", flag.ExitOnError)
	statusVerbose := statusCmd.Bool("verbose", false, "modo detallado")

	if len(os.Args) < 2 {
		fmt.Println("Uso: programa <comando> [opciones]")
		fmt.Println("Comandos: deploy, status")
		return
	}

	switch os.Args[1] {
	case "deploy":
		deployCmd.Parse(os.Args[2:])
		fmt.Printf("Deploy - Rama: %s, Force: %v\n", *deployBranch, *deployForce)
		if deployCmd.NArg() > 0 {
			fmt.Printf("Proyecto: %s\n", deployCmd.Arg(0))
		}

	case "status":
		statusCmd.Parse(os.Args[2:])
		fmt.Printf("Status - Verbose: %v\n", *statusVerbose)

	default:
		fmt.Printf("Comando desconocido: %s\n", os.Args[1])
		os.Exit(1)
	}
}

// ============================================
// PARTE 6: Ejemplo completo - Configurador de servidor
// ============================================

func configuradorServidor() {
	fmt.Println("\n=== Configurador de Servidor ===")

	// Definir todos los flags
	host := flag.String("host", "localhost", "dirección del servidor")
	port := flag.Int("port", 8080, "puerto del servidor")
	env := flag.String("env", "development", "entorno (development/staging/production)")
	verbose := flag.Bool("verbose", false, "modo verbose")
	workers := flag.Int("workers", 4, "número de workers")
	timeout := flag.Int("timeout", 30, "timeout en segundos")

	flag.Usage = func() {
		fmt.Println("Servidor HTTP - Configuración")
		fmt.Println()
		fmt.Println("Uso: servidor [opciones]")
		fmt.Println()
		fmt.Println("Opciones:")
		flag.PrintDefaults()
		fmt.Println()
		fmt.Println("Ejemplos:")
		fmt.Println("  servidor -port=3000 -env=production")
		fmt.Println("  servidor -host=0.0.0.0 -verbose")
	}

	flag.Parse()

	// Validar entorno
	validEnvs := []string{"development", "staging", "production"}
	envValid := false
	for _, e := range validEnvs {
		if *env == e {
			envValid = true
			break
		}
	}
	if !envValid {
		fmt.Fprintf(os.Stderr, "Error: entorno inválido '%s'\n", *env)
		fmt.Fprintf(os.Stderr, "Entornos válidos: %s\n", strings.Join(validEnvs, ", "))
		os.Exit(1)
	}

	// Validar puerto
	if *port < 1 || *port > 65535 {
		fmt.Fprintf(os.Stderr, "Error: puerto debe estar entre 1 y 65535\n")
		os.Exit(1)
	}

	// Mostrar configuración
	fmt.Println()
	fmt.Println("╔══════════════════════════════════════╗")
	fmt.Println("║     Configuración del Servidor       ║")
	fmt.Println("╚══════════════════════════════════════╝")
	fmt.Println()
	fmt.Printf("  Host:      %s\n", *host)
	fmt.Printf("  Puerto:    %d\n", *port)
	fmt.Printf("  Entorno:   %s\n", *env)
	fmt.Printf("  Workers:   %d\n", *workers)
	fmt.Printf("  Timeout:   %ds\n", *timeout)
	fmt.Printf("  Verbose:   %v\n", *verbose)
	fmt.Println()

	if flag.NArg() > 0 {
		fmt.Printf("  Extras:    %v\n", flag.Args())
		fmt.Println()
	}

	url := fmt.Sprintf("http://%s:%d", *host, *port)
	fmt.Printf("Servidor configurado en %s\n", url)

	if *env == "production" {
		fmt.Println("ADVERTENCIA: Modo producción activado")
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	// Descomentar la función que quieras probar:
	// NOTA: Solo puede haber un flag.Parse() por ejecución

	// demoFlagsBasicos()
	// demoFlagsConVariables()
	// demoUsagePersonalizado()
	// demoValidacion()
	// demoFlagSetSubcomandos()
	configuradorServidor()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea un programa con flags para configurar una base de datos:
	// -host, -port, -user, -password, -database, -ssl (bool)
	// Valida que user y database no estén vacíos

	// TODO 2: Agrega soporte para leer variables de entorno como fallback:
	// Si el flag no se pasa, busca DB_HOST, DB_PORT, etc.
	// Prioridad: flag > env > default

	// TODO 3: Crea un programa con subcomandos usando FlagSet:
	// programa user create -name=Juan -email=juan@email.com
	// programa user delete -id=123
	// programa user list -limit=10

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un CLI de gestión de configuración:
	//
	// programa config get <key>           - Obtiene un valor
	// programa config set <key> <value>   - Establece un valor
	// programa config list                - Lista toda la config
	// programa config delete <key>        - Elimina una config
	//
	// Flags globales:
	// -file    Archivo de configuración (default: config.json)
	// -format  Formato de salida (json/yaml/text)
	//
	// Flags por subcomando:
	// config list -filter=<prefix>
	// config set -force (sobrescribe sin confirmar)
	//
	// Bonus: Guarda la configuración en un archivo JSON real

	// Tu código aquí...
}
