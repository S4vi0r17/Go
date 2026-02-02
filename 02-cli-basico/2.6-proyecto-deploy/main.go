// Proyecto 2.6: Deploy CLI Completo
// Un CLI profesional con subcomandos, flags, colores y validación

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"time"
)

// ============================================
// Colores y estilos
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
// Datos simulados
// ============================================

var proyectos = map[string]map[string]string{
	"frontend": {
		"status": "running",
		"branch": "main",
		"port":   "3000",
		"tipo":   "react",
	},
	"backend": {
		"status": "running",
		"branch": "main",
		"port":   "8080",
		"tipo":   "go",
	},
	"api": {
		"status": "stopped",
		"branch": "develop",
		"port":   "9000",
		"tipo":   "go",
	},
	"worker": {
		"status": "running",
		"branch": "main",
		"port":   "9001",
		"tipo":   "python",
	},
}

// ============================================
// Helpers de output
// ============================================

func printHeader(title string) {
	width := 45
	padding := (width - len(title) - 2) / 2
	fmt.Println()
	fmt.Println(Bold + Blue + "╔" + strings.Repeat("═", width-2) + "╗" + Reset)
	fmt.Printf(Bold+Blue+"║"+Reset+"%s%s%s"+Bold+Blue+"║"+Reset+"\n",
		strings.Repeat(" ", padding), title, strings.Repeat(" ", width-padding-len(title)-2))
	fmt.Println(Bold + Blue + "╚" + strings.Repeat("═", width-2) + "╝" + Reset)
}

func success(msg string) {
	fmt.Println(Green + "✓ " + msg + Reset)
}

func errorMsg(msg string) {
	fmt.Fprintln(os.Stderr, Red+"✗ "+msg+Reset)
}

func warning(msg string) {
	fmt.Println(Yellow + "⚠ " + msg + Reset)
}

func step(msg string) {
	fmt.Printf("  %s›%s %s", Cyan, Reset, msg)
}

func stepDone() {
	fmt.Println(Green + " ✓" + Reset)
}

func listarNombresProyectos() string {
	nombres := make([]string, 0, len(proyectos))
	for n := range proyectos {
		nombres = append(nombres, n)
	}
	return strings.Join(nombres, ", ")
}

// ============================================
// AYUDA PRINCIPAL
// ============================================

func mostrarAyuda() {
	fmt.Println()
	fmt.Println(Bold + Blue + "Deploy CLI" + Reset + " - Herramienta de despliegue")
	fmt.Println(Dim + "Versión 1.0.0" + Reset)
	fmt.Println()
	fmt.Println(Bold + "Uso:" + Reset)
	fmt.Println("  deploy <comando> [opciones]")
	fmt.Println()
	fmt.Println(Bold + "Comandos:" + Reset)
	fmt.Println("  deploy   " + Dim + "Despliega un proyecto" + Reset)
	fmt.Println("  status   " + Dim + "Muestra el estado de servicios" + Reset)
	fmt.Println("  logs     " + Dim + "Muestra los logs de un servicio" + Reset)
	fmt.Println("  restart  " + Dim + "Reinicia un servicio" + Reset)
	fmt.Println("  list     " + Dim + "Lista todos los proyectos" + Reset)
	fmt.Println("  help     " + Dim + "Muestra esta ayuda" + Reset)
	fmt.Println()
	fmt.Println(Bold + "Ejemplos:" + Reset)
	fmt.Println("  deploy deploy frontend")
	fmt.Println("  deploy status -verbose")
	fmt.Println("  deploy logs api -lines=50")
	fmt.Println("  deploy restart backend -graceful")
	fmt.Println()
}

// ============================================
// COMANDO: deploy
// ============================================

func comandoDeploy(args []string) {
	fs := flag.NewFlagSet("deploy", flag.ExitOnError)
	branch := fs.String("branch", "main", "rama a desplegar")
	force := fs.Bool("force", false, "forzar despliegue")
	skipTests := fs.Bool("skip-tests", false, "omitir tests")
	env := fs.String("env", "development", "entorno de destino")

	fs.Usage = func() {
		fmt.Println(Bold + "deploy deploy" + Reset + " - Despliega un proyecto")
		fmt.Println()
		fmt.Println("Uso: deploy deploy <proyecto> [opciones]")
		fmt.Println()
		fmt.Println("Opciones:")
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Proyectos disponibles: " + listarNombresProyectos())
	}

	fs.Parse(args)

	if fs.NArg() < 1 {
		errorMsg("Se requiere nombre del proyecto")
		fs.Usage()
		os.Exit(1)
	}

	proyecto := fs.Arg(0)

	if _, existe := proyectos[proyecto]; !existe {
		errorMsg(fmt.Sprintf("Proyecto '%s' no encontrado", proyecto))
		fmt.Println("Proyectos disponibles: " + listarNombresProyectos())
		os.Exit(1)
	}

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
		errorMsg(fmt.Sprintf("Entorno '%s' inválido", *env))
		fmt.Println("Entornos válidos: " + strings.Join(validEnvs, ", "))
		os.Exit(1)
	}

	printHeader("DEPLOY")
	fmt.Println()
	fmt.Printf("  Proyecto: %s%s%s\n", Bold, proyecto, Reset)
	fmt.Printf("  Rama:     %s\n", *branch)
	fmt.Printf("  Entorno:  %s\n", *env)
	fmt.Println()

	if *force {
		warning("Modo forzado activado")
	}

	if *env == "production" {
		warning("Desplegando a PRODUCCIÓN")
	}

	fmt.Println()

	// Simular pasos de deploy
	pasos := []struct {
		nombre   string
		duracion time.Duration
	}{
		{"Conectando al servidor", 300 * time.Millisecond},
		{"Haciendo git pull", 500 * time.Millisecond},
		{"Instalando dependencias", 800 * time.Millisecond},
	}

	if !*skipTests {
		pasos = append(pasos, struct {
			nombre   string
			duracion time.Duration
		}{"Ejecutando tests", 1000 * time.Millisecond})
	}

	pasos = append(pasos,
		struct {
			nombre   string
			duracion time.Duration
		}{"Compilando", 600 * time.Millisecond},
		struct {
			nombre   string
			duracion time.Duration
		}{"Reiniciando servicio", 400 * time.Millisecond},
	)

	for _, paso := range pasos {
		step(paso.nombre + "...")
		time.Sleep(paso.duracion)
		stepDone()
	}

	fmt.Println()
	success(fmt.Sprintf("Deploy de %s completado exitosamente", proyecto))
	fmt.Println()
}

// ============================================
// COMANDO: status
// ============================================

func comandoStatus(args []string) {
	fs := flag.NewFlagSet("status", flag.ExitOnError)
	verbose := fs.Bool("verbose", false, "mostrar información detallada")
	fs.Parse(args)

	printHeader("ESTADO DE SERVICIOS")
	fmt.Println()

	running := 0
	stopped := 0

	for nombre, info := range proyectos {
		statusColor := Green
		statusIcon := "●"
		if info["status"] == "stopped" {
			statusColor = Red
			stopped++
		} else {
			running++
		}

		fmt.Printf("  %s%s%s %-12s %s%s%s",
			statusColor, statusIcon, Reset,
			nombre,
			statusColor, info["status"], Reset)

		if *verbose {
			fmt.Printf("  %s(:%s, %s, rama: %s)%s",
				Dim, info["port"], info["tipo"], info["branch"], Reset)
		}

		fmt.Println()
	}

	fmt.Println()
	fmt.Printf("  %s%d running%s | %s%d stopped%s\n",
		Green, running, Reset, Red, stopped, Reset)
	fmt.Println()
}

// ============================================
// COMANDO: logs
// ============================================

func comandoLogs(args []string) {
	fs := flag.NewFlagSet("logs", flag.ExitOnError)
	lines := fs.Int("lines", 10, "número de líneas a mostrar")
	follow := fs.Bool("follow", false, "seguir logs en tiempo real")
	level := fs.String("level", "", "filtrar por nivel (info/warn/error)")
	fs.Parse(args)

	if fs.NArg() < 1 {
		errorMsg("Se requiere nombre del servicio")
		os.Exit(1)
	}

	servicio := fs.Arg(0)

	if _, existe := proyectos[servicio]; !existe {
		errorMsg(fmt.Sprintf("Servicio '%s' no encontrado", servicio))
		os.Exit(1)
	}

	fmt.Printf("\n%sLogs de %s%s (últimas %d líneas)\n", Bold, servicio, Reset, *lines)
	fmt.Println(Dim + strings.Repeat("─", 60) + Reset)

	// Logs simulados
	logs := []struct {
		nivel   string
		mensaje string
		tiempo  string
	}{
		{"INFO", "Servidor iniciado en puerto " + proyectos[servicio]["port"], "10:30:01"},
		{"INFO", "Conexión a base de datos establecida", "10:30:02"},
		{"DEBUG", "Procesando request GET /api/users", "10:30:15"},
		{"INFO", "Request completado en 45ms", "10:30:15"},
		{"WARN", "Cache miss para key: user_123", "10:31:00"},
		{"INFO", "Regenerando cache...", "10:31:01"},
		{"DEBUG", "Procesando request POST /api/login", "10:32:00"},
		{"INFO", "Usuario autenticado: admin", "10:32:01"},
		{"ERROR", "Timeout en conexión externa", "10:33:00"},
		{"INFO", "Reintentando conexión...", "10:33:01"},
		{"INFO", "Conexión restaurada", "10:33:05"},
		{"DEBUG", "Health check OK", "10:34:00"},
	}

	count := 0
	for _, log := range logs {
		if count >= *lines {
			break
		}

		// Filtrar por nivel si se especifica
		if *level != "" && strings.ToLower(log.nivel) != strings.ToLower(*level) {
			continue
		}

		color := Reset
		switch log.nivel {
		case "ERROR":
			color = Red
		case "WARN":
			color = Yellow
		case "DEBUG":
			color = Dim
		case "INFO":
			color = Cyan
		}

		fmt.Printf("%s%s [%s] %s%s\n", color, log.tiempo, log.nivel, log.mensaje, Reset)
		count++
	}

	if *follow {
		fmt.Println()
		fmt.Println(Yellow + "Siguiendo logs... (Ctrl+C para salir)" + Reset)
	}
	fmt.Println()
}

// ============================================
// COMANDO: restart
// ============================================

func comandoRestart(args []string) {
	fs := flag.NewFlagSet("restart", flag.ExitOnError)
	graceful := fs.Bool("graceful", true, "reinicio graceful")
	timeout := fs.Int("timeout", 30, "timeout en segundos")
	fs.Parse(args)

	if fs.NArg() < 1 {
		errorMsg("Se requiere nombre del servicio")
		os.Exit(1)
	}

	servicio := fs.Arg(0)

	if _, existe := proyectos[servicio]; !existe {
		errorMsg(fmt.Sprintf("Servicio '%s' no encontrado", servicio))
		os.Exit(1)
	}

	modo := "graceful"
	if !*graceful {
		modo = "forzado"
	}

	fmt.Printf("\nReiniciando %s%s%s (modo: %s, timeout: %ds)\n",
		Bold, servicio, Reset, modo, *timeout)
	fmt.Println()

	step("Enviando señal de parada...")
	time.Sleep(500 * time.Millisecond)
	stepDone()

	step("Esperando que termine...")
	time.Sleep(800 * time.Millisecond)
	stepDone()

	step("Iniciando servicio...")
	time.Sleep(500 * time.Millisecond)
	stepDone()

	fmt.Println()
	success(fmt.Sprintf("%s reiniciado correctamente", servicio))
	fmt.Println()
}

// ============================================
// COMANDO: list
// ============================================

func comandoList() {
	printHeader("PROYECTOS")
	fmt.Println()

	for nombre, info := range proyectos {
		statusColor := Green
		if info["status"] == "stopped" {
			statusColor = Red
		}

		fmt.Printf("  %s●%s %-12s %s%-10s%s %s(:%s, rama: %s)%s\n",
			statusColor, Reset,
			nombre,
			Bold, info["tipo"], Reset,
			Dim, info["port"], info["branch"], Reset)
	}
	fmt.Println()
}

// ============================================
// MAIN - Router
// ============================================

func main() {
	if len(os.Args) < 2 {
		mostrarAyuda()
		os.Exit(0)
	}

	comando := os.Args[1]

	switch comando {
	case "deploy":
		comandoDeploy(os.Args[2:])
	case "status":
		comandoStatus(os.Args[2:])
	case "logs":
		comandoLogs(os.Args[2:])
	case "restart":
		comandoRestart(os.Args[2:])
	case "list":
		comandoList()
	case "help", "-h", "--help":
		mostrarAyuda()
	default:
		errorMsg(fmt.Sprintf("Comando desconocido: %s", comando))
		fmt.Println("Usa 'deploy help' para ver comandos disponibles")
		os.Exit(1)
	}
}

// ============================================
// EJERCICIOS FINALES
// ============================================

// Este proyecto es una base para expandir. Considera:
//
// 1. Guardar configuración de proyectos en un archivo YAML
// 2. Implementar ejecución real de comandos con os/exec
// 3. Agregar autenticación SSH para servidores remotos
// 4. Crear un comando "init" que configure un nuevo proyecto
// 5. Agregar un comando "rollback" para revertir deploys
// 6. Implementar notificaciones (Slack, email)
// 7. Agregar métricas y health checks reales
// 8. Crear modo watch que monitoree cambios
