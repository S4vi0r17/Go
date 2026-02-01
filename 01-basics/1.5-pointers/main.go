package main

import "fmt"

// Valor vs Referencia
func modificarValor(x int) {
	x = 100 // No modifica el original
}

func modificarPuntero(x *int) {
	*x = 100 // SÍ modifica el original
}

// En structs
type Contador struct {
	valor int
}

func (c *Contador) Incrementar() {
	c.valor++ // Modifica el struct original
}

func (c Contador) ObtenerValor() int {
	return c.valor // Solo lee, no necesita puntero
}

func main() {
	var x int
	var y *int

	x = 10
	y = &x

	fmt.Println("x: ", x)
	fmt.Println("y: ", y)
	fmt.Println("*y: ", *y)
	fmt.Printf("Type: %T, Value: %v, Dereferenced: %v\n", y, y, *y)

	*y = 20
	fmt.Println("\nx: ", x)
	fmt.Println("y: ", y)
	fmt.Println("*y: ", *y)
	fmt.Printf("Type: %T, Value: %v, Dereferenced: %v\n", y, y, *y)

	fmt.Println("\n\nUsando punteros en funciones")

	numero := 50

	fmt.Printf("Type: %T, Value: %v\n", numero, numero)

	modificarValor(numero)
	fmt.Println("Valor de numero no cambió: ", numero) // numero sigue siendo 50

	modificarPuntero(&numero)
	fmt.Println("Valor de numero cambió: ", numero) // numero ahora es 100

	fmt.Println("\n\nUsando punteros en structs")

	contador := Contador{valor: 0}

	fmt.Println("Valor inicial del contador: ", contador.ObtenerValor()) // Imprime 0

	contador.Incrementar()
	fmt.Println(contador.ObtenerValor()) // Imprime 1

	contadorPtr := &contador
	contadorPtr.Incrementar()
	fmt.Println("Puntero del struct: ", contadorPtr.ObtenerValor()) // Imprime 2

	fmt.Println("Struct original: ", contador.ObtenerValor()) // Imprime 2
}
