// Ejercicio 2.5: CLI con Subcomandos
// Objetivo: Crear CLIs profesionales con múltiples comandos

package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

// ============================================
// Colores
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
// PARTE 1: Estructura básica de subcomandos
// ============================================

// Patrón común:
// programa <subcomando> [flags] [argumentos]
//
// Ejemplos:
// git commit -m "mensaje"
// docker run -p 8080:80 nginx
// kubectl get pods -n default

func mostrarAyudaGlobal() {
	fmt.Println(Bold + "Mi CLI" + Reset + " - Herramienta de ejemplo")
	fmt.Println()
	fmt.Println(Bold + "Uso:" + Reset)
	fmt.Println("  programa <comando> [opciones]")
	fmt.Println()
	fmt.Println(Bold + "Comandos:" + Reset)
	fmt.Println("  deploy   Despliega un proyecto")
	fmt.Println("  status   Muestra estado de servicios")
	fmt.Println("  logs     Muestra logs de un servicio")
	fmt.Println("  config   Gestiona configuración")
	fmt.Println("  help     Muestra esta ayuda")
	fmt.Println()
	fmt.Println(Bold + "Flags globales:" + Reset)
	fmt.Println("  -v, --verbose   Modo verbose")
	fmt.Println("  -h, --help      Muestra ayuda")
	fmt.Println()
	fmt.Println("Usa 'programa <comando> --help' para más información")
}

// ============================================
// PARTE 2: Comando deploy
// ============================================

type DeployOptions struct {
	Branch    string
	Force     bool
	SkipTests bool
	Env       string
}

func comandoDeploy(args []string) {
	fs := flag.NewFlagSet("deploy", flag.ExitOnError)

	opts := DeployOptions{}
	fs.StringVar(&opts.Branch, "branch", "main", "rama a desplegar")
	fs.StringVar(&opts.Branch, "b", "main", "rama a desplegar (corto)")
	fs.BoolVar(&opts.Force, "force", false, "forzar despliegue")
	fs.BoolVar(&opts.Force, "f", false, "forzar despliegue (corto)")
	fs.BoolVar(&opts.SkipTests, "skip-tests", false, "omitir tests")
	fs.StringVar(&opts.Env, "env", "development", "entorno de destino")
	fs.StringVar(&opts.Env, "e", "development", "entorno (corto)")

	fs.Usage = func() {
		fmt.Println(Bold + "deploy" + Reset + " - Despliega un proyecto")
		fmt.Println()
		fmt.Println("Uso: programa deploy <proyecto> [opciones]")
		fmt.Println()
		fmt.Println("Argumentos:")
		fmt.Println("  proyecto    Nombre del proyecto a desplegar")
		fmt.Println()
		fmt.Println("Opciones:")
		fs.PrintDefaults()
		fmt.Println()
		fmt.Println("Ejemplos:")
		fmt.Println("  programa deploy frontend")
		fmt.Println("  programa deploy backend -b develop -e staging")
		fmt.Println("  programa deploy api --force --skip-tests")
	}

	fs.Parse(args)

	if fs.NArg() < 1 {
		fmt.Fprintln(os.Stderr, Red+"Error: se requiere nombre del proyecto"+Reset)
		fs.Usage()
		os.Exit(1)
	}

	proyecto := fs.Arg(0)

	fmt.Printf("\n%sDesplegando %s%s\n", Bold, proyecto, Reset)
	fmt.Printf("  Rama:    %s\n", opts.Branch)
	fmt.Printf("  Entorno: %s\n", opts.Env)
	fmt.Printf("  Force:   %v\n", opts.Force)
	fmt.Printf("  Tests:   %v\n", !opts.SkipTests)
	fmt.Println()

	// Simular pasos
	pasos := []string{
		"Conectando al servidor",
		"Clonando repositorio",
		"Instalando dependencias",
	}
	if !opts.SkipTests {
		pasos = append(pasos, "Ejecutando tests")
	}
	pasos = append(pasos, "Compilando", "Desplegando")

	for _, paso := range pasos {
		fmt.Printf("  %s›%s %s... %s✓%s\n", Cyan, Reset, paso, Green, Reset)
	}

	fmt.Printf("\n%s✓ Deploy completado%s\n", Green+Bold, Reset)
}

// ============================================
// PARTE 3: Comando status
// ============================================

type StatusOptions struct {
	All     bool
	Format  string
	Verbose bool
}

func comandoStatus(args []string) {
	fs := flag.NewFlagSet("status", flag.ExitOnError)

	opts := StatusOptions{}
	fs.BoolVar(&opts.All, "all", false, "mostrar todos los servicios")
	fs.BoolVar(&opts.All, "a", false, "mostrar todos (corto)")
	fs.StringVar(&opts.Format, "format", "table", "formato de salida (table/json/yaml)")
	fs.StringVar(&opts.Format, "o", "table", "formato (corto)")
	fs.BoolVar(&opts.Verbose, "verbose", false, "mostrar información detallada")
	fs.BoolVar(&opts.Verbose, "v", false, "verbose (corto)")

	fs.Usage = func() {
		fmt.Println(Bold + "status" + Reset + " - Muestra estado de servicios")
		fmt.Println()
		fmt.Println("Uso: programa status [servicio] [opciones]")
		fmt.Println()
		fmt.Println("Opciones:")
		fs.PrintDefaults()
	}

	fs.Parse(args)

	servicios := []struct {
		nombre string
		estado string
		puerto int
		uptime string
	}{
		{"frontend", "running", 3000, "5h 23m"},
		{"backend", "running", 8080, "5h 23m"},
		{"api", "stopped", 9000, "-"},
		{"worker", "running", 9001, "2h 15m"},
	}

	// Filtrar por servicio si se especifica
	filtro := ""
	if fs.NArg() > 0 {
		filtro = fs.Arg(0)
	}

	fmt.Println()
	fmt.Println(Bold + "Estado de Servicios" + Reset)
	fmt.Println(Dim + strings.Repeat("─", 40) + Reset)

	for _, s := range servicios {
		if filtro != "" && s.nombre != filtro {
			continue
		}

		icon := Green + "●" + Reset
		estado := Green + s.estado + Reset
		if s.estado == "stopped" {
			icon = Red + "●" + Reset
			estado = Red + s.estado + Reset
		}

		fmt.Printf("  %s %-12s %s", icon, s.nombre, estado)
		if opts.Verbose {
			fmt.Printf("  %s:%d  %s", Dim, s.puerto, s.uptime+Reset)
		}
		fmt.Println()
	}
	fmt.Println()
}

// ============================================
// PARTE 4: Comando logs
// ============================================

func comandoLogs(args []string) {
	fs := flag.NewFlagSet("logs", flag.ExitOnError)

	lines := fs.Int("lines", 10, "número de líneas")
	follow := fs.Bool("follow", false, "seguir logs en tiempo real")
	fs.IntVar(lines, "n", 10, "líneas (corto)")
	fs.BoolVar(follow, "f", false, "follow (corto)")

	fs.Usage = func() {
		fmt.Println(Bold + "logs" + Reset + " - Muestra logs de un servicio")
		fmt.Println()
		fmt.Println("Uso: programa logs <servicio> [opciones]")
		fmt.Println()
		fmt.Println("Opciones:")
		fs.PrintDefaults()
	}

	fs.Parse(args)

	if fs.NArg() < 1 {
		fmt.Fprintln(os.Stderr, Red+"Error: se requiere nombre del servicio"+Reset)
		fs.Usage()
		os.Exit(1)
	}

	servicio := fs.Arg(0)

	fmt.Printf("\n%sLogs de %s%s (últimas %d líneas)\n", Bold, servicio, Reset, *lines)
	fmt.Println(Dim + strings.Repeat("─", 50) + Reset)

	// Logs simulados
	logs := []struct {
		nivel   string
		mensaje string
	}{
		{"INFO", "Servidor iniciado en puerto 8080"},
		{"INFO", "Conexión a base de datos establecida"},
		{"DEBUG", "Request GET /api/users - 45ms"},
		{"WARN", "Cache miss para key: user_123"},
		{"INFO", "Regenerando cache"},
		{"ERROR", "Timeout en conexión externa"},
		{"INFO", "Reintentando conexión"},
		{"INFO", "Conexión restaurada"},
		{"DEBUG", "Request POST /api/login - 120ms"},
		{"INFO", "Usuario autenticado: admin"},
	}

	for i := 0; i < *lines && i < len(logs); i++ {
		log := logs[i]
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
		fmt.Printf("%s[%s] %s%s\n", color, log.nivel, log.mensaje, Reset)
	}

	if *follow {
		fmt.Println()
		fmt.Println(Yellow + "Siguiendo logs... (Ctrl+C para salir)" + Reset)
	}
}

// ============================================
// PARTE 5: Comando config con subsubcomandos
// ============================================

func comandoConfig(args []string) {
	if len(args) < 1 {
		fmt.Println(Bold + "config" + Reset + " - Gestiona configuración")
		fmt.Println()
		fmt.Println("Subcomandos:")
		fmt.Println("  get <key>           Obtiene un valor")
		fmt.Println("  set <key> <value>   Establece un valor")
		fmt.Println("  list                Lista toda la config")
		fmt.Println("  delete <key>        Elimina una config")
		return
	}

	subcomando := args[0]

	// Config simulada
	config := map[string]string{
		"server.host":    "localhost",
		"server.port":    "8080",
		"database.host":  "localhost",
		"database.port":  "5432",
		"database.name":  "myapp",
		"log.level":      "info",
		"log.format":     "json",
	}

	switch subcomando {
	case "get":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, Red+"Error: se requiere key"+Reset)
			os.Exit(1)
		}
		key := args[1]
		if val, ok := config[key]; ok {
			fmt.Println(val)
		} else {
			fmt.Fprintf(os.Stderr, "Key '%s' no encontrada\n", key)
			os.Exit(1)
		}

	case "set":
		if len(args) < 3 {
			fmt.Fprintln(os.Stderr, Red+"Error: se requiere key y value"+Reset)
			os.Exit(1)
		}
		key := args[1]
		value := args[2]
		fmt.Printf("%s✓ %s = %s%s\n", Green, key, value, Reset)

	case "list":
		fmt.Println()
		fmt.Println(Bold + "Configuración actual:" + Reset)
		for k, v := range config {
			fmt.Printf("  %s%s%s = %s\n", Cyan, k, Reset, v)
		}

	case "delete":
		if len(args) < 2 {
			fmt.Fprintln(os.Stderr, Red+"Error: se requiere key"+Reset)
			os.Exit(1)
		}
		key := args[1]
		fmt.Printf("%s✓ Eliminado: %s%s\n", Yellow, key, Reset)

	default:
		fmt.Fprintf(os.Stderr, "Subcomando desconocido: %s\n", subcomando)
		os.Exit(1)
	}
}

// ============================================
// MAIN - Router de comandos
// ============================================

func main() {
	if len(os.Args) < 2 {
		mostrarAyudaGlobal()
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

	case "config":
		comandoConfig(os.Args[2:])

	case "help", "-h", "--help":
		mostrarAyudaGlobal()

	default:
		fmt.Fprintf(os.Stderr, Red+"Comando desconocido: %s\n"+Reset, comando)
		fmt.Println("Usa 'programa help' para ver comandos disponibles")
		os.Exit(1)
	}

	// ============================================
	// EJERCICIOS
	// ============================================

	// TODO 1: Agrega un comando "restart" con flags:
	// -graceful (bool) - reinicio graceful
	// -timeout (int)   - timeout en segundos
	// -all             - reiniciar todos los servicios

	// TODO 2: Agrega un comando "exec" que ejecute comandos en servicios:
	// programa exec frontend "npm run build"
	// Con flags: -detach, -interactive

	// TODO 3: Agrega aliases a los comandos:
	// "s" -> "status"
	// "d" -> "deploy"
	// "l" -> "logs"

	// ============================================
	// RETO FINAL
	// ============================================

	// Implementa un CLI completo de gestión de tareas:
	//
	// programa task add "Título" [-priority=high] [-due=2024-01-20]
	// programa task list [-status=pending] [-priority=high]
	// programa task complete <id>
	// programa task delete <id>
	// programa task edit <id> [-title=...] [-priority=...]
	// programa task search "término"
	//
	// programa project create <nombre>
	// programa project list
	// programa project tasks <nombre>
	//
	// Bonus: Guarda las tareas en un archivo JSON
}
