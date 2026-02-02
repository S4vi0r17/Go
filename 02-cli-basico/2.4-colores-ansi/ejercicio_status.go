// Ejercicio 04: Status con Colores
// CLI que muestra el estado de servicios con colores

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// ============================================
// Colores ANSI
// ============================================

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
	Dim    = "\033[2m"
)

// ============================================
// Estructuras de datos
// ============================================

type Servicio struct {
	Nombre      string
	Estado      string // running, stopped, warning, error
	Puerto      int
	CPU         float64
	Memoria     int // MB
	Uptime      string
	UltimoCheck time.Time
}

// ============================================
// Datos de ejemplo
// ============================================

func obtenerServicios() []Servicio {
	return []Servicio{
		{
			Nombre:      "nginx",
			Estado:      "running",
			Puerto:      80,
			CPU:         2.5,
			Memoria:     128,
			Uptime:      "5d 12h",
			UltimoCheck: time.Now(),
		},
		{
			Nombre:      "postgres",
			Estado:      "running",
			Puerto:      5432,
			CPU:         15.3,
			Memoria:     512,
			Uptime:      "5d 12h",
			UltimoCheck: time.Now(),
		},
		{
			Nombre:      "redis",
			Estado:      "running",
			Puerto:      6379,
			CPU:         0.8,
			Memoria:     64,
			Uptime:      "5d 12h",
			UltimoCheck: time.Now(),
		},
		{
			Nombre:      "api-backend",
			Estado:      "running",
			Puerto:      8080,
			CPU:         45.2,
			Memoria:     256,
			Uptime:      "2h 30m",
			UltimoCheck: time.Now(),
		},
		{
			Nombre:      "worker-jobs",
			Estado:      "warning",
			Puerto:      9000,
			CPU:         78.5,
			Memoria:     384,
			Uptime:      "1h 15m",
			UltimoCheck: time.Now(),
		},
		{
			Nombre:      "cache-server",
			Estado:      "stopped",
			Puerto:      11211,
			CPU:         0,
			Memoria:     0,
			Uptime:      "-",
			UltimoCheck: time.Now().Add(-10 * time.Minute),
		},
		{
			Nombre:      "mail-service",
			Estado:      "error",
			Puerto:      25,
			CPU:         0,
			Memoria:     0,
			Uptime:      "-",
			UltimoCheck: time.Now().Add(-5 * time.Minute),
		},
	}
}

// ============================================
// Funciones de display
// ============================================

func colorPorEstado(estado string) string {
	switch estado {
	case "running":
		return Green
	case "stopped":
		return Red
	case "warning":
		return Yellow
	case "error":
		return Red + Bold
	default:
		return Reset
	}
}

func iconoPorEstado(estado string) string {
	switch estado {
	case "running":
		return "●"
	case "stopped":
		return "○"
	case "warning":
		return "◐"
	case "error":
		return "✗"
	default:
		return "?"
	}
}

func barraProgreso(porcentaje float64, ancho int) string {
	lleno := int(porcentaje / 100 * float64(ancho))
	vacio := ancho - lleno

	color := Green
	if porcentaje > 70 {
		color = Yellow
	}
	if porcentaje > 90 {
		color = Red
	}

	return color + strings.Repeat("█", lleno) + Dim + strings.Repeat("░", vacio) + Reset
}

func imprimirHeader() {
	fmt.Println()
	fmt.Println(Bold + Blue + "╔══════════════════════════════════════════════════════════════╗" + Reset)
	fmt.Println(Bold + Blue + "║" + Reset + "                    " + Bold + "SYSTEM STATUS MONITOR" + Reset + "                    " + Bold + Blue + "║" + Reset)
	fmt.Println(Bold + Blue + "╚══════════════════════════════════════════════════════════════╝" + Reset)
	fmt.Println()
}

func imprimirResumen(servicios []Servicio) {
	running := 0
	stopped := 0
	warning := 0
	errored := 0

	for _, s := range servicios {
		switch s.Estado {
		case "running":
			running++
		case "stopped":
			stopped++
		case "warning":
			warning++
		case "error":
			errored++
		}
	}

	fmt.Println(Bold + "Resumen:" + Reset)
	fmt.Printf("  %s● Running: %d%s  ", Green, running, Reset)
	fmt.Printf("%s◐ Warning: %d%s  ", Yellow, warning, Reset)
	fmt.Printf("%s○ Stopped: %d%s  ", Dim, stopped, Reset)
	fmt.Printf("%s✗ Error: %d%s\n", Red, errored, Reset)
	fmt.Println()
}

func imprimirTabla(servicios []Servicio) {
	// Header de la tabla
	fmt.Println(Dim + "┌────────────────┬──────────┬────────┬─────────────────────┬────────────┐" + Reset)
	fmt.Printf(Dim+"│"+Reset+" %-14s "+Dim+"│"+Reset+" %-8s "+Dim+"│"+Reset+" %-6s "+Dim+"│"+Reset+" %-19s "+Dim+"│"+Reset+" %-10s "+Dim+"│"+Reset+"\n",
		Bold+"Servicio"+Reset, Bold+"Estado"+Reset, Bold+"Puerto"+Reset, Bold+"CPU"+Reset, Bold+"Uptime"+Reset)
	fmt.Println(Dim + "├────────────────┼──────────┼────────┼─────────────────────┼────────────┤" + Reset)

	// Filas
	for _, s := range servicios {
		color := colorPorEstado(s.Estado)
		icono := iconoPorEstado(s.Estado)

		// Nombre y estado
		fmt.Printf(Dim+"│"+Reset+" %s%s%s %-12s ", color, icono, Reset, s.Nombre)
		fmt.Printf(Dim+"│"+Reset+" %s%-8s%s ", color, s.Estado, Reset)

		// Puerto
		if s.Puerto > 0 {
			fmt.Printf(Dim+"│"+Reset+" %-6d ", s.Puerto)
		} else {
			fmt.Printf(Dim+"│"+Reset+" %-6s ", "-")
		}

		// CPU con barra
		if s.Estado == "running" || s.Estado == "warning" {
			barra := barraProgreso(s.CPU, 10)
			fmt.Printf(Dim+"│"+Reset+" %s %5.1f%% ", barra, s.CPU)
		} else {
			fmt.Printf(Dim+"│"+Reset+" %-19s ", "-")
		}

		// Uptime
		fmt.Printf(Dim+"│"+Reset+" %-10s "+Dim+"│"+Reset+"\n", s.Uptime)
	}

	fmt.Println(Dim + "└────────────────┴──────────┴────────┴─────────────────────┴────────────┘" + Reset)
}

func imprimirAlertas(servicios []Servicio) {
	var alertas []string

	for _, s := range servicios {
		switch s.Estado {
		case "error":
			alertas = append(alertas, fmt.Sprintf("%s✗ ERROR:%s %s está caído", Red+Bold, Reset, s.Nombre))
		case "warning":
			alertas = append(alertas, fmt.Sprintf("%s⚠ WARN:%s %s tiene alta carga (CPU: %.1f%%)", Yellow, Reset, s.Nombre, s.CPU))
		case "stopped":
			alertas = append(alertas, fmt.Sprintf("%s○ INFO:%s %s está detenido", Dim, Reset, s.Nombre))
		}
	}

	if len(alertas) > 0 {
		fmt.Println()
		fmt.Println(Bold + "Alertas:" + Reset)
		for _, a := range alertas {
			fmt.Println("  " + a)
		}
	}
}

// ============================================
// Main
// ============================================

func main() {
	// Verificar argumento --help
	if len(os.Args) > 1 && (os.Args[1] == "-h" || os.Args[1] == "--help") {
		fmt.Println("status - Muestra el estado de los servicios")
		fmt.Println()
		fmt.Println("Uso: status [opciones]")
		fmt.Println()
		fmt.Println("Opciones:")
		fmt.Println("  -h, --help     Muestra esta ayuda")
		fmt.Println("  -s, --simple   Muestra versión simplificada")
		return
	}

	servicios := obtenerServicios()

	imprimirHeader()
	imprimirResumen(servicios)
	imprimirTabla(servicios)
	imprimirAlertas(servicios)

	fmt.Println()
	fmt.Printf("%sÚltima actualización: %s%s\n", Dim, time.Now().Format("2006-01-02 15:04:05"), Reset)
	fmt.Println()
}
