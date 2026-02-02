// Ejercicio 3: Structs
// Objetivo: Dominar la definición y uso de structs (fundamental para BubbleTea)

package main

import (
	"fmt"
	"time"
)

// ============================================
// PARTE 1: Struct básico
// ============================================

// Definición de struct
type Persona struct {
	Nombre string
	Edad   int
	Email  string
}

// ============================================
// PARTE 2: Structs anidados
// ============================================

type Direccion struct {
	Calle  string
	Ciudad string
	Pais   string
	Codigo string
}

type Empresa struct {
	Nombre    string
	Direccion Direccion // Struct anidado
	Empleados int
}

// ============================================
// PARTE 3: Struct con campos privados/públicos
// ============================================

// En Go, la visibilidad se controla con mayúsculas/minúsculas:
// - Mayúscula = Público (exportado)
// - Minúscula = Privado (solo en el mismo package)

type Configuracion struct {
	Host     string // Público
	Puerto   int    // Público
	password string // Privado (minúscula)
}

// Constructor (patrón común en Go)
func NuevaConfiguracion(host string, puerto int, password string) *Configuracion {
	return &Configuracion{
		Host:     host,
		Puerto:   puerto,
		password: password,
	}
}

// Getter para campo privado
func (c *Configuracion) Password() string {
	return c.password
}

// ============================================
// PARTE 4: Structs con tags (para JSON, YAML)
// ============================================

type ProyectoConfig struct {
	Nombre       string `yaml:"nombre" json:"nombre"`
	Ruta         string `yaml:"ruta" json:"ruta"`
	Tipo         string `yaml:"tipo" json:"tipo"`
	Rama         string `yaml:"rama" json:"rama"`
	ComandoBuild string `yaml:"comando_build" json:"comando_build"`
}

// ============================================
// PARTE 5: Composición (embebido)
// ============================================

// Struct base
type ConexionBase struct {
	Host      string
	Puerto    int
	Conectado bool
}

func (c *ConexionBase) Conectar() {
	c.Conectado = true
	fmt.Printf("Conectado a %s:%d\n", c.Host, c.Puerto)
}

func (c *ConexionBase) Desconectar() {
	c.Conectado = false
	fmt.Println("Desconectado")
}

// Struct que embebe otro (composición)
type ConexionSSH struct {
	ConexionBase // Embebido (hereda campos y métodos)
	Usuario      string
	ClavePath    string
}

// Puede tener sus propios métodos adicionales
func (c *ConexionSSH) EjecutarComando(cmd string) string {
	if !c.Conectado { // Accede al campo del struct embebido
		return "Error: no conectado"
	}
	return fmt.Sprintf("Ejecutando '%s' como %s", cmd, c.Usuario)
}

// ============================================
// PARTE 6: Struct para estado de UI (como BubbleTea)
// ============================================

// Este es el patrón que usa BubbleTea
type ModeloUI struct {
	// Estado de navegación
	vistaActual int
	cursor      int

	// Datos
	opciones      []string
	opcionElegida string

	// Estado de operaciones
	cargando bool
	error    error

	// Dimensiones
	ancho int
	alto  int
}

// Constantes para vistas
const (
	VistaMenu = iota
	VistaProyectos
	VistaResultado
)

func NuevoModeloUI() ModeloUI {
	return ModeloUI{
		vistaActual: VistaMenu,
		cursor:      0,
		opciones:    []string{"Deploy", "Logs", "Status", "Salir"},
		cargando:    false,
	}
}

func (m ModeloUI) OpcionActual() string {
	if m.cursor >= 0 && m.cursor < len(m.opciones) {
		return m.opciones[m.cursor]
	}
	return ""
}

func (m *ModeloUI) MoverArriba() {
	if m.cursor > 0 {
		m.cursor--
	}
}

func (m *ModeloUI) MoverAbajo() {
	if m.cursor < len(m.opciones)-1 {
		m.cursor++
	}
}

func (m *ModeloUI) Seleccionar() {
	m.opcionElegida = m.OpcionActual()
}

// ============================================
// MAIN
// ============================================

func main() {
	fmt.Println("=== Struct Básico ===")

	// Inicialización con campos nombrados (recomendado)
	persona1 := Persona{
		Nombre: "Juan",
		Edad:   30,
		Email:  "juan@email.com",
	}
	fmt.Printf("Persona: %+v\n", persona1)

	// Inicialización posicional (no recomendado)
	persona2 := Persona{"Ana", 25, "ana@email.com"}
	fmt.Printf("Persona: %+v\n", persona2)

	// Acceso a campos
	fmt.Printf("Nombre: %s, Edad: %d\n", persona1.Nombre, persona1.Edad)

	// Modificar campos
	persona1.Edad = 31
	fmt.Printf("Nueva edad: %d\n", persona1.Edad)

	fmt.Println("\n=== Structs Anidados ===")

	empresa := Empresa{
		Nombre: "TechCorp",
		Direccion: Direccion{
			Calle:  "Calle Principal 123",
			Ciudad: "Madrid",
			Pais:   "España",
			Codigo: "28001",
		},
		Empleados: 50,
	}
	fmt.Printf("Empresa: %s\n", empresa.Nombre)
	fmt.Printf("Ciudad: %s\n", empresa.Direccion.Ciudad)

	fmt.Println("\n=== Constructor y Campos Privados ===")

	config := NuevaConfiguracion("localhost", 22, "secreto123")
	fmt.Printf("Host: %s\n", config.Host)
	fmt.Printf("Puerto: %d\n", config.Puerto)
	// fmt.Println(config.password) // ERROR: campo privado
	fmt.Printf("Password (via getter): %s\n", config.Password())

	fmt.Println("\n=== Composición (Embebido) ===")

	ssh := ConexionSSH{
		ConexionBase: ConexionBase{
			Host:   "servidor.com",
			Puerto: 22,
		},
		Usuario:   "admin",
		ClavePath: "~/.ssh/id_rsa",
	}

	// Accede a campos del struct embebido directamente
	fmt.Printf("Host: %s\n", ssh.Host) // No necesita ssh.ConexionBase.Host

	// Usa métodos del struct embebido
	ssh.Conectar() // Método de ConexionBase

	// Usa métodos propios
	resultado := ssh.EjecutarComando("ls -la")
	fmt.Println(resultado)

	ssh.Desconectar()

	fmt.Println("\n=== Modelo de UI ===")

	modelo := NuevoModeloUI()
	fmt.Printf("Vista actual: %d\n", modelo.vistaActual)
	fmt.Printf("Opción actual: %s\n", modelo.OpcionActual())

	modelo.MoverAbajo()
	fmt.Printf("Después de bajar: %s\n", modelo.OpcionActual())

	modelo.MoverAbajo()
	modelo.Seleccionar()
	fmt.Printf("Seleccionado: %s\n", modelo.opcionElegida)

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea un struct 'Tarea' con campos:
	// - ID (int)
	// - Titulo (string)
	// - Completada (bool)
	// - FechaCreacion (time.Time)

	// TODO 2: Crea un método 'Completar' que marque la tarea como completada

	// TODO 3: Crea un método 'Resumen' que retorne un string descriptivo

	// TODO 4: Crea un slice de tareas y practiva iterarlas

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea la estructura de datos para un "Deploy Manager":
	//
	// 1. Struct 'Proyecto' con:
	//    - Nombre, Ruta, Tipo, Rama
	//
	// 2. Struct 'Servidor' con:
	//    - Host, Puerto, Usuario
	//
	// 3. Struct 'DeployConfig' que contenga:
	//    - Servidor (embebido o anidado)
	//    - Proyectos (slice de Proyecto)
	//    - UltimoDespliegue (time.Time)
	//
	// 4. Métodos:
	//    - AgregarProyecto(p Proyecto)
	//    - ListarProyectos() []string
	//    - BuscarProyecto(nombre string) (*Proyecto, bool)
	//
	// Ejemplo de uso:
	// dm := DeployConfig{...}
	// dm.AgregarProyecto(Proyecto{Nombre: "frontend", ...})
	// proyectos := dm.ListarProyectos()

	// Tu código aquí...

	// Placeholder para usar time
	_ = time.Now()
}
