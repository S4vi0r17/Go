// Ejercicio 7: Manejo de Errores
// Objetivo: Dominar el patrón de errores de Go (no hay excepciones)

package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================
// PARTE 1: El patrón básico de errores
// ============================================

// En Go, las funciones que pueden fallar retornan (valor, error)
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("división por cero")
	}
	return a / b, nil
}

// ============================================
// PARTE 2: Errores personalizados
// ============================================

// Errores como variables (para comparación)
var (
	ErrNoEncontrado      = errors.New("elemento no encontrado")
	ErrConexionFallida   = errors.New("conexión fallida")
	ErrPermisosDenegados = errors.New("permisos denegados")
	ErrTimeout           = errors.New("operación timeout")
)

func buscarProyecto(nombre string, proyectos []string) (int, error) {
	for i, p := range proyectos {
		if p == nombre {
			return i, nil
		}
	}
	return -1, ErrNoEncontrado
}

// ============================================
// PARTE 3: Errores con contexto (wrapping)
// ============================================

func conectarServidor(host string) error {
	// Simular fallo
	if host == "" {
		return fmt.Errorf("conectando a servidor: %w", ErrConexionFallida)
	}
	if host == "blocked.com" {
		return fmt.Errorf("conectando a %s: %w", host, ErrPermisosDenegados)
	}
	return nil
}

func desplegarProyecto(proyecto, servidor string) error {
	err := conectarServidor(servidor)
	if err != nil {
		// Envolver error con más contexto
		return fmt.Errorf("desplegando %s: %w", proyecto, err)
	}
	return nil
}

// ============================================
// PARTE 4: Verificar tipo de error
// ============================================

func manejarErrorConexion(err error) {
	if err == nil {
		fmt.Println("Sin error")
		return
	}

	// errors.Is verifica si el error ES o CONTIENE cierto error
	if errors.Is(err, ErrConexionFallida) {
		fmt.Println("Error de conexión - intentando reconectar...")
	} else if errors.Is(err, ErrPermisosDenegados) {
		fmt.Println("Error de permisos - contacta al administrador")
	} else if errors.Is(err, ErrTimeout) {
		fmt.Println("Timeout - intenta más tarde")
	} else {
		fmt.Printf("Error desconocido: %v\n", err)
	}
}

// ============================================
// PARTE 5: Tipos de error personalizados
// ============================================

// Error como struct (más información)
type ErrorValidacion struct {
	Campo   string
	Mensaje string
	// Valor   interface{}
	Valor any // Es lo mismo que interface{}
}

// Implementar interface error
func (e *ErrorValidacion) Error() string {
	return fmt.Sprintf("validación fallida en '%s': %s (valor: %v)",
		e.Campo, e.Mensaje, e.Valor)
}

func validarPuerto(puerto int) error {
	if puerto < 1 || puerto > 65535 {
		return &ErrorValidacion{
			Campo:   "puerto",
			Mensaje: "debe estar entre 1 y 65535",
			Valor:   puerto,
		}
	}
	if puerto < 1024 {
		return &ErrorValidacion{
			Campo:   "puerto",
			Mensaje: "puertos < 1024 requieren permisos de root",
			Valor:   puerto,
		}
	}
	return nil
}

// Extraer información del error personalizado
func manejarErrorValidacion(err error) {
	var errVal *ErrorValidacion
	if errors.As(err, &errVal) {
		fmt.Printf("Campo con error: %s\n", errVal.Campo)
		fmt.Printf("Valor problemático: %v\n", errVal.Valor)
		fmt.Printf("Mensaje: %s\n", errVal.Mensaje)
	} else {
		fmt.Printf("Otro error: %v\n", err)
	}
}

// ============================================
// PARTE 6: Múltiples errores
// ============================================

type ErrorMultiple struct {
	errores []error
}

func (e *ErrorMultiple) Error() string {
	if len(e.errores) == 0 {
		return "sin errores"
	}

	// msg := fmt.Sprintf("%d errores:\n", len(e.errores))
	// for i, err := range e.errores {
	// 	msg += fmt.Sprintf("  %d. %v\n", i+1, err)
	// }
	// return msg
	//

	var b strings.Builder
	fmt.Fprintf(&b, "%d errores:\n", len(e.errores))
	for i, err := range e.errores {
		fmt.Fprintf(&b, "  %d. %v\n", i+1, err)
	}
	return b.String()
}

func (e *ErrorMultiple) Add(err error) {
	if err != nil {
		e.errores = append(e.errores, err)
	}
}

func (e *ErrorMultiple) HasErrors() bool {
	return len(e.errores) > 0
}

func validarConfiguracion(host string, puerto int, usuario string) error {
	errores := &ErrorMultiple{}

	if host == "" {
		errores.Add(errors.New("host es requerido"))
	}

	if err := validarPuerto(puerto); err != nil {
		errores.Add(err)
	}

	if len(usuario) < 3 {
		errores.Add(errors.New("usuario debe tener al menos 3 caracteres"))
	}

	if errores.HasErrors() {
		return errores
	}
	return nil
}

// ============================================
// PARTE 7: Patrones comunes
// ============================================

// Patrón: Retorno temprano
func procesarArchivo(path string) error {
	// Verificar existencia
	_, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("verificando archivo: %w", err)
	}

	// Leer contenido
	data, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("leyendo archivo: %w", err)
	}

	// Procesar (simulado)
	if len(data) == 0 {
		return errors.New("archivo vacío")
	}

	fmt.Printf("Procesado: %d bytes\n", len(data))
	return nil
}

// Patrón: Ignorar error intencionalmente
func obtenerPuertoODefault(s string) int {
	puerto, err := strconv.Atoi(s)
	if err != nil {
		return 8080 // Default si hay error
	}
	return puerto
}

// ============================================
// MAIN
// ============================================

func main() {
	fmt.Println("=== Error Básico ===")

	resultado, err := dividir(10, 2)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("10 / 2 = %.2f\n", resultado)
	}

	resultado, err = dividir(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	}

	fmt.Println("\n=== Errores Predefinidos ===")

	proyectos := []string{"frontend", "backend", "api"}

	idx, err := buscarProyecto("backend", proyectos)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Proyecto encontrado en índice: %d\n", idx)
	}

	idx, err = buscarProyecto("mobile", proyectos)
	if err != nil {
		if errors.Is(err, ErrNoEncontrado) {
			fmt.Println("El proyecto no existe")
		}
	}

	fmt.Println("\n=== Errores con Contexto ===")

	err = desplegarProyecto("frontend", "")
	if err != nil {
		fmt.Println("Error completo:", err)
		// Output: desplegando frontend: conectando a servidor: conexión fallida

		// Verificar error interno
		if errors.Is(err, ErrConexionFallida) {
			fmt.Println("→ Causa raíz: problema de conexión")
		}
	}

	fmt.Println("\n=== Manejo por Tipo ===")

	manejarErrorConexion(nil)
	manejarErrorConexion(desplegarProyecto("api", ""))
	manejarErrorConexion(desplegarProyecto("api", "blocked.com"))
	manejarErrorConexion(errors.New("error raro"))

	fmt.Println("\n=== Error Personalizado (Struct) ===")

	err = validarPuerto(80)
	if err != nil {
		manejarErrorValidacion(err)
	}

	err = validarPuerto(99999)
	if err != nil {
		manejarErrorValidacion(err)
	}

	fmt.Println("\n=== Múltiples Errores ===")

	err = validarConfiguracion("", -1, "ab")
	if err != nil {
		fmt.Println(err)
	}

	err = validarConfiguracion("localhost", 8080, "admin")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Configuración válida")
	}

	fmt.Println("\n=== Puerto con Default ===")

	fmt.Printf("Puerto de '3000': %d\n", obtenerPuertoODefault("3000"))
	fmt.Printf("Puerto de 'invalid': %d\n", obtenerPuertoODefault("invalid"))

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea errores predefinidos para:
	// - ErrArchivoNoExiste
	// - ErrFormatoInvalido
	// - ErrSinPermisos

	// TODO 2: Crea una función 'leerConfigYAML(path string) (*Config, error)'
	// que retorne diferentes errores según el problema:
	// - Si el archivo no existe
	// - Si el formato es inválido
	// - Si faltan campos requeridos

	// TODO 3: Crea un tipo de error 'ErrorSSH' que contenga:
	// - Host string
	// - Puerto int
	// - Causa error
	// E implementa el método Error()

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un sistema de deploy con manejo robusto de errores:
	//
	// 1. Tipos de error:
	//    - ErrorConexion{Host, Puerto, Causa}
	//    - ErrorDeploy{Proyecto, Paso, Causa}
	//    - ErrorValidacion (ya definido arriba)
	//
	// 2. Función 'Deploy(proyecto, host string, puerto int) error' que:
	//    - Valide parámetros (usa validarConfiguracion como base)
	//    - Simule conexión (falla si host contiene "fail")
	//    - Simule pasos de deploy (git pull, build, restart)
	//    - Retorne errores apropiados con contexto
	//
	// 3. Función 'ManejarDeploy(err error)' que:
	//    - Identifique el tipo de error
	//    - Imprima instrucciones específicas para cada caso
	//    - Use errors.Is y errors.As apropiadamente

	// Tu código aquí...
}
