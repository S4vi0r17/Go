// Ejercicio 2.4: Colores ANSI en Terminal
// Objetivo: Crear CLIs visualmente atractivos

package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"
)

// ============================================
// PARTE 1: Códigos ANSI básicos
// ============================================

// Los códigos ANSI son secuencias de escape que la terminal interpreta
// Formato: \033[<código>m  o  \x1b[<código>m

const (
	// Reset
	Reset = "\033[0m"

	// Colores de texto
	Black   = "\033[30m"
	Red     = "\033[31m"
	Green   = "\033[32m"
	Yellow  = "\033[33m"
	Blue    = "\033[34m"
	Magenta = "\033[35m"
	Cyan    = "\033[36m"
	White   = "\033[37m"

	// Colores brillantes
	BrightBlack   = "\033[90m"
	BrightRed     = "\033[91m"
	BrightGreen   = "\033[92m"
	BrightYellow  = "\033[93m"
	BrightBlue    = "\033[94m"
	BrightMagenta = "\033[95m"
	BrightCyan    = "\033[96m"
	BrightWhite   = "\033[97m"

	// Estilos
	Bold      = "\033[1m"
	Dim       = "\033[2m"
	Italic    = "\033[3m"
	Underline = "\033[4m"
	Blink     = "\033[5m"
	Reverse   = "\033[7m"
	Hidden    = "\033[8m"
	Strike    = "\033[9m"

	// Colores de fondo
	BgBlack   = "\033[40m"
	BgRed     = "\033[41m"
	BgGreen   = "\033[42m"
	BgYellow  = "\033[43m"
	BgBlue    = "\033[44m"
	BgMagenta = "\033[45m"
	BgCyan    = "\033[46m"
	BgWhite   = "\033[47m"
)

func demoColoresBasicos() {
	fmt.Println("=== Colores Básicos ===")

	fmt.Println(Red + "Texto rojo" + Reset)
	fmt.Println(Green + "Texto verde" + Reset)
	fmt.Println(Yellow + "Texto amarillo" + Reset)
	fmt.Println(Blue + "Texto azul" + Reset)
	fmt.Println(Magenta + "Texto magenta" + Reset)
	fmt.Println(Cyan + "Texto cyan" + Reset)
}

func demoEstilos() {
	fmt.Println("\n=== Estilos ===")

	fmt.Println(Bold + "Texto en negrita" + Reset)
	fmt.Println(Dim + "Texto tenue" + Reset)
	fmt.Println(Italic + "Texto en cursiva" + Reset)
	fmt.Println(Underline + "Texto subrayado" + Reset)
	fmt.Println(Strike + "Texto tachado" + Reset)
	fmt.Println(Reverse + "Texto invertido" + Reset)
}

func demoCombinaciones() {
	fmt.Println("\n=== Combinaciones ===")

	fmt.Println(Bold + Red + "Error: algo falló" + Reset)
	fmt.Println(Bold + Green + "Éxito: operación completada" + Reset)
	fmt.Println(Bold + Yellow + "Advertencia: revisar config" + Reset)
	fmt.Println(Bold + Blue + Underline + "Título importante" + Reset)
	fmt.Println(BgRed + White + " CRÍTICO " + Reset + " Servidor caído")
	fmt.Println(BgGreen + Black + " OK " + Reset + " Tests pasaron")
}

// ============================================
// PARTE 2: Funciones helper
// ============================================

// Colorizar texto
func colorize(text, color string) string {
	return color + text + Reset
}

// Mensajes con formato
func success(msg string) {
	fmt.Println(Green + "✓ " + msg + Reset)
}

func errorMsg(msg string) {
	fmt.Fprintln(os.Stderr, Red+"✗ "+msg+Reset)
}

func warning(msg string) {
	fmt.Println(Yellow + "⚠ " + msg + Reset)
}

func info(msg string) {
	fmt.Println(Cyan + "ℹ " + msg + Reset)
}

func debug(msg string) {
	fmt.Println(Dim + "• " + msg + Reset)
}

func demoHelpers() {
	fmt.Println("\n=== Funciones Helper ===")

	success("Deploy completado")
	errorMsg("No se pudo conectar")
	warning("Versión antigua detectada")
	info("Procesando 42 archivos")
	debug("Conexión establecida en 45ms")
}

// ============================================
// PARTE 3: Barras de progreso
// ============================================

func progressBar(current, total int, width int) string {
	percent := float64(current) / float64(total)
	filled := int(percent * float64(width))
	empty := width - filled

	bar := strings.Repeat("█", filled) + strings.Repeat("░", empty)
	return fmt.Sprintf("[%s] %3.0f%%", bar, percent*100)
}

func spinner(frame int) string {
	frames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	return frames[frame%len(frames)]
}

func demoProgreso() {
	fmt.Println("\n=== Barras de Progreso ===")

	for i := 0; i <= 10; i++ {
		fmt.Printf("\r%s Descargando... %s", Cyan, progressBar(i, 10, 20)+Reset)
		// time.Sleep(200 * time.Millisecond)
	}
	fmt.Println()

	fmt.Println("\nSpinners:")
	for i := 0; i < 10; i++ {
		fmt.Printf("%s ", spinner(i))
	}
	fmt.Println()
}

// ============================================
// PARTE 4: Tablas con colores
// ============================================

func printTable(headers []string, rows [][]string) {
	// Calcular anchos
	widths := make([]int, len(headers))
	for i, h := range headers {
		widths[i] = len(h)
	}
	for _, row := range rows {
		for i, cell := range row {
			if len(cell) > widths[i] {
				widths[i] = len(cell)
			}
		}
	}

	// Línea separadora
	separator := "+"
	for _, w := range widths {
		separator += strings.Repeat("-", w+2) + "+"
	}

	// Imprimir header
	fmt.Println(separator)
	fmt.Print("|")
	for i, h := range headers {
		fmt.Printf(" %s%-*s%s |", Bold+Cyan, widths[i], h, Reset)
	}
	fmt.Println()
	fmt.Println(separator)

	// Imprimir filas
	for _, row := range rows {
		fmt.Print("|")
		for i, cell := range row {
			// Colorizar según contenido
			color := Reset
			if cell == "running" {
				color = Green
			} else if cell == "stopped" {
				color = Red
			} else if cell == "warning" {
				color = Yellow
			}
			fmt.Printf(" %s%-*s%s |", color, widths[i], cell, Reset)
		}
		fmt.Println()
	}
	fmt.Println(separator)
}

func demoTabla() {
	fmt.Println("\n=== Tabla con Colores ===")

	headers := []string{"Servicio", "Estado", "Puerto", "Uptime"}
	rows := [][]string{
		{"frontend", "running", "3000", "5h 23m"},
		{"backend", "running", "8080", "5h 23m"},
		{"database", "running", "5432", "12h 45m"},
		{"cache", "stopped", "6379", "-"},
		{"worker", "warning", "9000", "1h 02m"},
	}

	printTable(headers, rows)
}

// ============================================
// PARTE 5: Cajas y bordes
// ============================================

func box(title, content string, width int) {
	// Caracteres de caja
	topLeft := "╔"
	topRight := "╗"
	bottomLeft := "╚"
	bottomRight := "╝"
	horizontal := "═"
	vertical := "║"

	// Línea superior con título
	titleLen := len(title)
	padding := (width - titleLen - 4) / 2
	fmt.Println(Bold + Blue + topLeft + strings.Repeat(horizontal, padding+1) + " " + title + " " + strings.Repeat(horizontal, width-padding-titleLen-4) + topRight + Reset)

	// Contenido
	lines := strings.Split(content, "\n")
	for _, line := range lines {
		fmt.Printf("%s%s%s %-*s %s%s%s\n", Bold, Blue, vertical, width-4, line, vertical, Reset, "")
	}

	// Línea inferior
	fmt.Println(Bold + Blue + bottomLeft + strings.Repeat(horizontal, width-2) + bottomRight + Reset)
}

func demoBox() {
	fmt.Println("\n=== Cajas y Bordes ===")

	box("Información", "Usuario: admin\nServidor: localhost\nPuerto: 8080", 30)
	fmt.Println()
	box("Estado", "Todo funcionando\ncorrectamente", 25)
}

// ============================================
// PARTE 6: Detectar soporte de colores
// ============================================

func supportsColor() bool {
	// En Windows, necesitamos habilitar VT100
	if runtime.GOOS == "windows" {
		// Windows 10+ soporta ANSI
		return true
	}

	// Verificar si es una terminal
	if os.Getenv("TERM") == "dumb" {
		return false
	}

	// Verificar NO_COLOR
	if os.Getenv("NO_COLOR") != "" {
		return false
	}

	return true
}

// Wrapper condicional
func colorIf(text, color string) string {
	if supportsColor() {
		return color + text + Reset
	}
	return text
}

// ============================================
// PARTE 7: Ejemplo de CLI con colores
// ============================================

func statusCLI() {
	fmt.Println("\n=== CLI de Estado con Colores ===")

	fmt.Println()
	fmt.Println(Bold + "  Deploy Manager v1.0" + Reset)
	fmt.Println(Dim + "  " + strings.Repeat("─", 30) + Reset)
	fmt.Println()

	servicios := []struct {
		nombre string
		estado string
		puerto int
	}{
		{"frontend", "running", 3000},
		{"backend", "running", 8080},
		{"api", "stopped", 9000},
		{"worker", "running", 9001},
	}

	for _, s := range servicios {
		icon := Green + "●" + Reset
		estado := Green + s.estado + Reset
		if s.estado == "stopped" {
			icon = Red + "●" + Reset
			estado = Red + s.estado + Reset
		}

		fmt.Printf("  %s %-12s %s %s:%d%s\n",
			icon, s.nombre, estado, Dim, s.puerto, Reset)
	}

	fmt.Println()
	fmt.Printf("  %s%d servicios activos%s, %s%d detenidos%s\n",
		Green, 3, Reset, Red, 1, Reset)
	fmt.Println()
}

// ============================================
// MAIN
// ============================================

func main() {
	demoColoresBasicos()
	demoEstilos()
	demoCombinaciones()
	demoHelpers()
	demoProgreso()
	demoTabla()
	demoBox()
	statusCLI()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea un logger con niveles coloreados:
	// log.Debug("mensaje") -> gris
	// log.Info("mensaje")  -> cyan
	// log.Warn("mensaje")  -> amarillo
	// log.Error("mensaje") -> rojo
	// log.Fatal("mensaje") -> rojo + exit

	// TODO 2: Crea una función que dibuje un árbol de archivos:
	// proyecto/
	// ├── src/
	// │   ├── main.go
	// │   └── utils.go
	// └── README.md

	// TODO 3: Crea un diff con colores:
	// - línea eliminada (rojo)
	// + línea agregada (verde)
	// línea sin cambios (normal)

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un "dashboard" de sistema:
	//
	// ╔════════════════════════════════════════╗
	// ║           SYSTEM DASHBOARD             ║
	// ╚════════════════════════════════════════╝
	//
	// CPU:    [████████░░░░░░░░░░░░] 40%
	// Memory: [████████████░░░░░░░░] 60%
	// Disk:   [██████████████████░░] 90% ⚠
	//
	// ┌─ Services ──────────────────────────────┐
	// │ ● nginx      running   :80              │
	// │ ● postgres   running   :5432            │
	// │ ○ redis      stopped   :6379            │
	// └─────────────────────────────────────────┘
	//
	// Last update: 2024-01-15 10:30:45
	//
	// Usa colores apropiados para cada métrica
	// Barra de disco en rojo si > 80%
	// Servicios en verde/rojo según estado

	// Tu código aquí...
}
