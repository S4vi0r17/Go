package main

import "fmt"

// Definición de interfaz
type Ejecutable interface {
	Ejecutar() error
	Nombre() string
}

// Cualquier tipo que implemente estos métodos satisface la interfaz
type ComandoLocal struct {
	cmd string
}

func (c ComandoLocal) Ejecutar() error {
	// ejecutar comando
	return nil
}

func (c ComandoLocal) Nombre() string {
	return c.cmd
}

// Uso polimórfico
func correr(e Ejecutable) {
	fmt.Printf("Ejecutando: %s\n", e.Nombre())
	e.Ejecutar()
}

func main() {
	c := ComandoLocal{cmd: "ls"}
	correr(c)
}
