// Ejercicio 5: Punteros
// Objetivo: Entender cuándo usar punteros (crítico para modificar estado)

package main

import "fmt"

// ============================================
// PARTE 1: Valor vs Referencia
// ============================================

// Por valor: se copia el dato
func duplicarValor(x int) {
	x = x * 2
	fmt.Printf("  Dentro de función: x = %d\n", x)
}

// Por referencia: se pasa la dirección de memoria
func duplicarPuntero(x *int) {
	*x = *x * 2
	fmt.Printf("  Dentro de función: *x = %d\n", *x)
}

// ============================================
// PARTE 2: Structs y punteros
// ============================================

type Contador struct {
	valor  int
	nombre string
}

// Método con receiver por valor (NO modifica)
func (c Contador) IncrementarValor() {
	c.valor++ // Modifica una COPIA
}

// Método con receiver por puntero (SÍ modifica)
func (c *Contador) IncrementarPuntero() {
	c.valor++ // Modifica el ORIGINAL
}

// Método que solo lee (puede ser valor o puntero)
func (c Contador) ObtenerValor() int {
	return c.valor
}

func (c *Contador) Reset() {
	c.valor = 0
}

// ============================================
// PARTE 3: Cuándo usar punteros
// ============================================

// Struct grande - usar puntero para evitar copias
type ConfiguracionGrande struct {
	Servidor   string
	Puerto     int
	Usuario    string
	Password   string
	BaseDatos  string
	SSL        bool
	Timeout    int
	MaxConn    int
	MinConn    int
	PoolSize   int
	LogLevel   string
	LogPath    string
	CacheTTL   int
	RetryCount int
	// ... muchos más campos
}

// Pasar por puntero es más eficiente
func procesarConfig(cfg *ConfiguracionGrande) {
	fmt.Printf("Procesando config para: %s\n", cfg.Servidor)
}

// ============================================
// PARTE 4: nil y verificación
// ============================================

type Conexion struct {
	host      string
	conectado bool
}

func (c *Conexion) Conectar() error {
	if c == nil {
		return fmt.Errorf("conexión es nil")
	}
	c.conectado = true
	return nil
}

func (c *Conexion) EstaConectado() bool {
	if c == nil {
		return false
	}
	return c.conectado
}

// ============================================
// PARTE 5: Punteros en slices y maps
// ============================================

type Usuario struct {
	ID     int
	Nombre string
	Activo bool
}

// Slice de punteros permite modificar elementos
type BaseDatosUsuarios struct {
	usuarios []*Usuario
}

func (db *BaseDatosUsuarios) Agregar(u *Usuario) {
	db.usuarios = append(db.usuarios, u)
}

func (db *BaseDatosUsuarios) BuscarPorID(id int) *Usuario {
	for _, u := range db.usuarios {
		if u.ID == id {
			return u // Retorna puntero al original
		}
	}
	return nil
}

func (db *BaseDatosUsuarios) DesactivarTodos() {
	for _, u := range db.usuarios {
		u.Activo = false // Modifica el original
	}
}

// ============================================
// PARTE 6: new() vs &{}
// ============================================

func crearContadores() {
	// Forma 1: new() - retorna puntero a zero value
	c1 := new(Contador)
	fmt.Printf("new(): %+v\n", c1)

	// Forma 2: &{} - retorna puntero a struct literal
	c2 := &Contador{valor: 10, nombre: "contador2"}
	fmt.Printf("&{}: %+v\n", c2)

	// Forma 3: Declarar y luego tomar dirección
	c3 := Contador{valor: 20, nombre: "contador3"}
	c3Ptr := &c3
	fmt.Printf("&var: %+v\n", c3Ptr)
}

// ============================================
// MAIN
// ============================================

func main() {
	fmt.Println("=== Conceptos Básicos de Punteros ===")

	var x int
	var y *int

	x = 10
	y = &x // y apunta a x

	fmt.Println("x: ", x)
	fmt.Println("y (dirección): ", y)
	fmt.Println("*y (valor): ", *y)
	fmt.Printf("Type: %T, Value: %v, Dereferenced: %v\n", y, y, *y)

	*y = 20 // Modificar a través del puntero
	fmt.Println("\nDespués de *y = 20:")
	fmt.Println("x: ", x)
	fmt.Println("*y: ", *y)

	fmt.Println("\n=== Valor vs Referencia ===")

	numero := 10
	fmt.Printf("Antes: numero = %d\n", numero)

	fmt.Println("\nPasando por valor:")
	duplicarValor(numero)
	fmt.Printf("Después: numero = %d (sin cambios)\n", numero)

	fmt.Println("\nPasando por puntero:")
	duplicarPuntero(&numero)
	fmt.Printf("Después: numero = %d (modificado)\n", numero)

	fmt.Println("\n=== Structs y Métodos ===")

	contador := Contador{valor: 0, nombre: "mi-contador"}
	fmt.Printf("Inicial: %d\n", contador.ObtenerValor())

	contador.IncrementarValor() // No modifica
	fmt.Printf("Después de IncrementarValor: %d\n", contador.ObtenerValor())

	contador.IncrementarPuntero() // Sí modifica
	fmt.Printf("Después de IncrementarPuntero: %d\n", contador.ObtenerValor())

	contador.IncrementarPuntero()
	contador.IncrementarPuntero()
	fmt.Printf("Después de 2 más: %d\n", contador.ObtenerValor())

	contador.Reset()
	fmt.Printf("Después de Reset: %d\n", contador.ObtenerValor())

	fmt.Println("\n=== nil Safety ===")

	var conexion *Conexion // nil por defecto
	fmt.Printf("Conexión nil está conectada: %v\n", conexion.EstaConectado())

	err := conexion.Conectar()
	if err != nil {
		fmt.Println("Error:", err)
	}

	// Crear conexión real
	conexion = &Conexion{host: "localhost"}
	conexion.Conectar()
	fmt.Printf("Conexión real está conectada: %v\n", conexion.EstaConectado())

	fmt.Println("\n=== Slice de Punteros ===")

	db := &BaseDatosUsuarios{}

	db.Agregar(&Usuario{ID: 1, Nombre: "Juan", Activo: true})
	db.Agregar(&Usuario{ID: 2, Nombre: "Ana", Activo: true})
	db.Agregar(&Usuario{ID: 3, Nombre: "Carlos", Activo: true})

	// Buscar y modificar
	usuario := db.BuscarPorID(2)
	if usuario != nil {
		fmt.Printf("Encontrado: %+v\n", usuario)
		usuario.Nombre = "Ana María" // Modifica el original
	}

	// Verificar que se modificó
	usuario2 := db.BuscarPorID(2)
	fmt.Printf("Después de modificar: %+v\n", usuario2)

	// Desactivar todos
	fmt.Println("\nAntes de desactivar:")
	for _, u := range db.usuarios {
		fmt.Printf("  %s: Activo=%v\n", u.Nombre, u.Activo)
	}

	db.DesactivarTodos()

	fmt.Println("\nDespués de desactivar:")
	for _, u := range db.usuarios {
		fmt.Printf("  %s: Activo=%v\n", u.Nombre, u.Activo)
	}

	fmt.Println("\n=== new() vs &{} ===")
	crearContadores()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea una función 'intercambiar(a, b *int)' que intercambie los valores

	// TODO 2: Crea un struct 'Pila' con un slice interno y métodos:
	// - Push(valor int) - agrega al final
	// - Pop() (int, bool) - quita y retorna el último
	// - Peek() (int, bool) - retorna el último sin quitar
	// Decide cuáles métodos necesitan puntero

	// TODO 3: ¿Por qué este código no funciona como esperas?
	// numeros := []int{1, 2, 3}
	// for _, n := range numeros {
	//     n = n * 2  // ¿Modifica el slice?
	// }
	// Arréglalo para que duplique cada número

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un "Registro de Estado" para aplicaciones:
	//
	// 1. Struct 'EstadoApp' con:
	//    - Nombre string
	//    - Estado string (running, stopped, error)
	//    - UltimoError *string (nil si no hay error)
	//    - Metricas *Metricas
	//
	// 2. Struct 'Metricas' con:
	//    - CPU float64
	//    - Memoria int64
	//    - Requests int
	//
	// 3. Struct 'Registro' con:
	//    - apps map[string]*EstadoApp
	//
	// 4. Métodos del Registro:
	//    - Registrar(nombre string) - crea nueva app en estado "stopped"
	//    - ObtenerApp(nombre string) *EstadoApp
	//    - IniciarApp(nombre string) error
	//    - DetenerApp(nombre string) error
	//    - ReportarError(nombre string, err string)
	//    - ActualizarMetricas(nombre string, cpu float64, mem int64, req int)
	//    - Listar() - imprime estado de todas las apps

	// Tu código aquí...
}
