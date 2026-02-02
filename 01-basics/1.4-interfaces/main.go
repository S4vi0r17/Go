// Ejercicio 4: Interfaces
// Objetivo: Entender interfaces (tea.Model ES una interfaz)

package main

import (
	"fmt"
	"strings"
)

// ============================================
// PARTE 1: Qué es una interfaz
// ============================================

// Una interfaz define UN CONTRATO de métodos
// Cualquier tipo que tenga esos métodos la implementa AUTOMÁTICAMENTE

// Interfaz simple
type Saludador interface {
	Saludar() string
}

// Tipos que implementan Saludador
type Persona struct {
	Nombre string
}

func (p Persona) Saludar() string {
	return "¡Hola! Soy " + p.Nombre
}

type Robot struct {
	Modelo string
}

func (r Robot) Saludar() string {
	return "BEEP BOOP. Soy " + r.Modelo
}

// Función que acepta cualquier Saludador
func imprimirSaludo(s Saludador) {
	fmt.Println(s.Saludar())
}

// ============================================
// PARTE 2: Interfaz con múltiples métodos
// ============================================

// Esta es similar a tea.Model de BubbleTea
type Componente interface {
	Iniciar() error
	Actualizar(evento string) error
	Renderizar() string
}

// Implementación: Botón
type Boton struct {
	Texto     string
	Activo    bool
	clickeado bool
}

func (b *Boton) Iniciar() error {
	b.Activo = true
	return nil
}

func (b *Boton) Actualizar(evento string) error {
	if evento == "click" {
		b.clickeado = true
	}
	return nil
}

func (b *Boton) Renderizar() string {
	estado := "[ ]"
	if b.clickeado {
		estado = "[x]"
	}
	return fmt.Sprintf("%s %s", estado, b.Texto)
}

// Implementación: Campo de texto
type CampoTexto struct {
	Etiqueta string
	Valor    string
	enfocado bool
}

func (c *CampoTexto) Iniciar() error {
	c.enfocado = false
	return nil
}

func (c *CampoTexto) Actualizar(evento string) error {
	switch evento {
	case "focus":
		c.enfocado = true
	case "blur":
		c.enfocado = false
	default:
		c.Valor += evento // Agregar texto
	}
	return nil
}

func (c *CampoTexto) Renderizar() string {
	cursor := ""
	if c.enfocado {
		cursor = "|"
	}
	return fmt.Sprintf("%s: [%s%s]", c.Etiqueta, c.Valor, cursor)
}

// Función que trabaja con cualquier Componente
func renderizarComponentes(componentes []Componente) string {
	var resultado strings.Builder
	for _, c := range componentes {
		resultado.WriteString(c.Renderizar())
		resultado.WriteString("\n")
	}
	return resultado.String()
}

// ============================================
// PARTE 3: La interfaz vacía (any)
// ============================================

// interface{} acepta CUALQUIER tipo
// En Go 1.18+ se puede usar 'any' como alias

func imprimirCualquierCosa(valor any) {
	fmt.Printf("Valor: %v (tipo: %T)\n", valor, valor)
}

// ============================================
// PARTE 4: Type assertion y type switch
// ============================================

func procesarMensaje(msg any) {
	// Type assertion: convertir interface{} a tipo concreto
	if texto, ok := msg.(string); ok {
		fmt.Println("Es un string:", texto)
		return
	}

	// Type switch: manejar múltiples tipos
	switch m := msg.(type) {
	case int:
		fmt.Println("Es un int:", m)
	case bool:
		fmt.Println("Es un bool:", m)
	case []string:
		fmt.Println("Es un slice de strings:", m)
	default:
		fmt.Printf("Tipo desconocido: %T\n", m)
	}
}

// ============================================
// PARTE 5: Ejemplo real - Similar a BubbleTea
// ============================================

// Mensaje (como tea.Msg)
type Msg interface{}

// Comando (como tea.Cmd)
type Cmd func() Msg

// Modelo (como tea.Model)
type Model interface {
	Init() Cmd
	Update(msg Msg) (Model, Cmd)
	View() string
}

// Mensajes concretos
type KeyMsg struct {
	Tecla string
}

type TickMsg struct{}

type QuitMsg struct{}

// Implementación de Model: Contador
type ContadorModel struct {
	valor int
}

func (m ContadorModel) Init() Cmd {
	return nil
}

func (m ContadorModel) Update(msg Msg) (Model, Cmd) {
	switch msg := msg.(type) {
	case KeyMsg:
		switch msg.Tecla {
		case "up", "k":
			m.valor++
		case "down", "j":
			m.valor--
		case "q":
			return m, func() Msg { return QuitMsg{} }
		}
	}
	return m, nil
}

func (m ContadorModel) View() string {
	return fmt.Sprintf("Contador: %d\n\nup/down: cambiar valor\nq: salir", m.valor)
}

// ============================================
// MAIN
// ============================================

func main() {
	fmt.Println("=== Interfaz Simple ===")

	persona := Persona{Nombre: "Juan"}
	robot := Robot{Modelo: "R2-D2"}

	// Ambos implementan Saludador
	imprimirSaludo(persona)
	imprimirSaludo(robot)

	// Slice de interfaces
	saludadores := []Saludador{persona, robot}
	for _, s := range saludadores {
		fmt.Println("-", s.Saludar())
	}

	fmt.Println("\n=== Interfaz Componente ===")

	boton := &Boton{Texto: "Aceptar"}
	campo := &CampoTexto{Etiqueta: "Nombre"}

	// Iniciar componentes
	boton.Iniciar()
	campo.Iniciar()

	// Simular eventos
	campo.Actualizar("focus")
	campo.Actualizar("J")
	campo.Actualizar("u")
	campo.Actualizar("a")
	campo.Actualizar("n")
	boton.Actualizar("click")

	// Renderizar usando la interfaz
	componentes := []Componente{campo, boton}
	fmt.Println(renderizarComponentes(componentes))

	fmt.Println("=== Interfaz Vacía (any) ===")

	imprimirCualquierCosa("hola")
	imprimirCualquierCosa(42)
	imprimirCualquierCosa(true)
	imprimirCualquierCosa([]string{"a", "b", "c"})

	fmt.Println("\n=== Type Switch ===")

	procesarMensaje("texto")
	procesarMensaje(100)
	procesarMensaje(true)
	procesarMensaje(3.14)

	fmt.Println("\n=== Patrón BubbleTea ===")

	var modelo Model = ContadorModel{valor: 0}

	// Simular el loop de BubbleTea
	fmt.Println("Estado inicial:")
	fmt.Println(modelo.View())

	// Simular teclas
	eventos := []Msg{
		KeyMsg{Tecla: "up"},
		KeyMsg{Tecla: "up"},
		KeyMsg{Tecla: "up"},
		KeyMsg{Tecla: "down"},
	}

	for _, evento := range eventos {
		modelo, _ = modelo.Update(evento)
	}

	fmt.Println("\nDespués de eventos:")
	fmt.Println(modelo.View())

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea una interfaz 'Ejecutable' con método:
	// - Ejecutar() (string, error)

	// TODO 2: Crea dos tipos que implementen Ejecutable:
	// - ComandoLocal{cmd string}
	// - ComandoRemoto{cmd string, host string}

	// TODO 3: Crea una función 'ejecutarTodos(comandos []Ejecutable)'
	// que ejecute todos los comandos e imprima resultados

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea tu propio mini-framework de UI:
	//
	// 1. Interfaz 'Vista' con métodos:
	//    - Nombre() string
	//    - Renderizar() string
	//    - ManejarInput(tecla string) (siguienteVista string)
	//
	// 2. Implementa 3 vistas:
	//    - VistaMenu (muestra opciones, retorna nombre de vista según selección)
	//    - VistaProyectos (lista proyectos)
	//    - VistaSalir (despedida)
	//
	// 3. Crea un 'Router' que:
	//    - Tenga un map[string]Vista
	//    - Tenga un método Navegar(nombreVista string)
	//    - Renderice la vista actual

	// Tu código aquí...
}
