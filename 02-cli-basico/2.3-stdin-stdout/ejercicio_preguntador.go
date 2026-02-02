// Ejercicio 03: Preguntador Interactivo
// CLI que hace preguntas y muestra un resumen

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ============================================
// Helpers de input
// ============================================

func pregunta(reader *bufio.Reader, prompt string) string {
	fmt.Print(prompt)
	respuesta, _ := reader.ReadString('\n')
	return strings.TrimSpace(respuesta)
}

func preguntaRequerida(reader *bufio.Reader, prompt string) string {
	for {
		respuesta := pregunta(reader, prompt)
		if respuesta != "" {
			return respuesta
		}
		fmt.Fprintln(os.Stderr, "Este campo es requerido")
	}
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

func preguntaNumero(reader *bufio.Reader, prompt string, min, max int) int {
	for {
		input := pregunta(reader, prompt)
		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Por favor ingresa un número válido\n")
			continue
		}
		if num < min || num > max {
			fmt.Fprintf(os.Stderr, "El número debe estar entre %d y %d\n", min, max)
			continue
		}
		return num
	}
}

func preguntaOpcion(reader *bufio.Reader, prompt string, opciones []string) string {
	fmt.Println(prompt)
	for i, op := range opciones {
		fmt.Printf("  %d. %s\n", i+1, op)
	}

	seleccion := preguntaNumero(reader, "Selección: ", 1, len(opciones))
	return opciones[seleccion-1]
}

// ============================================
// Estructuras de datos
// ============================================

type Perfil struct {
	Nombre        string
	Edad          int
	Email         string
	Ocupacion     string
	Experiencia   int
	Lenguajes     []string
	ProyectoFav   string
	UsaLinux      bool
	Newsletter    bool
}

// ============================================
// Main - Cuestionario
// ============================================

func main() {
	reader := bufio.NewReader(os.Stdin)
	perfil := Perfil{}

	fmt.Println()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║     CUESTIONARIO DE DESARROLLADOR      ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	// Datos personales
	fmt.Println("=== Datos Personales ===")
	perfil.Nombre = preguntaRequerida(reader, "Nombre completo: ")
	perfil.Edad = preguntaNumero(reader, "Edad: ", 10, 120)

	// Email con validación básica
	for {
		perfil.Email = pregunta(reader, "Email: ")
		if strings.Contains(perfil.Email, "@") && strings.Contains(perfil.Email, ".") {
			break
		}
		fmt.Fprintln(os.Stderr, "Email inválido, intenta de nuevo")
	}

	// Datos profesionales
	fmt.Println("\n=== Datos Profesionales ===")

	ocupaciones := []string{"Estudiante", "Junior Developer", "Senior Developer", "Tech Lead", "DevOps", "Otro"}
	perfil.Ocupacion = preguntaOpcion(reader, "Ocupación:", ocupaciones)

	perfil.Experiencia = preguntaNumero(reader, "Años de experiencia: ", 0, 50)

	// Lenguajes (múltiples respuestas)
	fmt.Println("\n=== Lenguajes de Programación ===")
	fmt.Println("Ingresa los lenguajes que conoces (uno por línea, línea vacía para terminar):")

	for {
		lang := pregunta(reader, "> ")
		if lang == "" {
			break
		}
		perfil.Lenguajes = append(perfil.Lenguajes, lang)
	}

	// Preferencias
	fmt.Println("\n=== Preferencias ===")
	perfil.ProyectoFav = preguntaConDefault(reader, "Tu proyecto favorito de Go", "BubbleTea")
	perfil.UsaLinux = preguntaSiNo(reader, "¿Usas Linux como sistema principal?", true)
	perfil.Newsletter = preguntaSiNo(reader, "¿Te gustaría recibir nuestro newsletter?", false)

	// Resumen
	fmt.Println()
	fmt.Println("╔════════════════════════════════════════╗")
	fmt.Println("║           RESUMEN DEL PERFIL           ║")
	fmt.Println("╚════════════════════════════════════════╝")
	fmt.Println()

	fmt.Printf("  Nombre:       %s\n", perfil.Nombre)
	fmt.Printf("  Edad:         %d años\n", perfil.Edad)
	fmt.Printf("  Email:        %s\n", perfil.Email)
	fmt.Printf("  Ocupación:    %s\n", perfil.Ocupacion)
	fmt.Printf("  Experiencia:  %d años\n", perfil.Experiencia)

	if len(perfil.Lenguajes) > 0 {
		fmt.Printf("  Lenguajes:    %s\n", strings.Join(perfil.Lenguajes, ", "))
	}

	fmt.Printf("  Proyecto fav: %s\n", perfil.ProyectoFav)
	fmt.Printf("  Usa Linux:    %v\n", perfil.UsaLinux)
	fmt.Printf("  Newsletter:   %v\n", perfil.Newsletter)

	// Confirmación
	fmt.Println()
	if preguntaSiNo(reader, "¿La información es correcta?", true) {
		fmt.Println("\n✓ Perfil guardado exitosamente!")
	} else {
		fmt.Println("\nPerfil descartado. Vuelve a ejecutar el programa.")
		os.Exit(0)
	}
}
