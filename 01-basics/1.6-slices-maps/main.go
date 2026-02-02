// Ejercicio 6: Slices y Maps
// Objetivo: Dominar las colecciones principales de Go

package main

import (
	"fmt"
	"sort"
	"strings"
)

// ============================================
// PARTE 1: Arrays (estáticos)
// ============================================

func arraysBasico() {
	fmt.Println("=== Arrays ===")

	// Array estático (tamaño fijo)
	numeros := [3]int{1, 2, 3}
	fmt.Printf("Array: %v (len=%d)\n", numeros, len(numeros))

	// Otra forma de declarar
	array2 := [...]int{1, 2, 3, 4, 5} // El compilador cuenta
	fmt.Printf("Array con [...]: %v\n", array2)
}

// ============================================
// PARTE 2: Slices - Básico
// ============================================

func slicesBasico() {
	fmt.Println("\n=== Slices Básico ===")

	// Crear slice literal
	frutas := []string{"manzana", "banana", "naranja"}
	fmt.Printf("Frutas: %v\n", frutas)
	fmt.Printf("Longitud: %d, Capacidad: %d\n", len(frutas), cap(frutas))

	// Crear slice vacío con make
	numeros := make([]int, 0, 10) // len=0, cap=10
	fmt.Printf("Números vacío: %v (len=%d, cap=%d)\n", numeros, len(numeros), cap(numeros))

	// Agregar elementos con append
	numeros = append(numeros, 1)
	numeros = append(numeros, 2, 3, 4)
	numeros = append(numeros, []int{5, 6, 7}...) // Expandir slice
	fmt.Printf("Números: %v\n", numeros)

	// Acceder por índice
	fmt.Printf("Primer número: %d\n", numeros[0])
	fmt.Printf("Último número: %d\n", numeros[len(numeros)-1])

	// Modificar elemento
	frutas[0] = "pera"
	fmt.Printf("Frutas modificadas: %v\n", frutas)
}

// ============================================
// PARTE 3: Slices - Slicing (subslices)
// ============================================

func slicing() {
	fmt.Println("\n=== Slicing ===")

	numeros := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Printf("Original: %v\n", numeros)

	// [inicio:fin] - incluye inicio, excluye fin
	fmt.Printf("[2:5]: %v\n", numeros[2:5]) // [2, 3, 4]
	fmt.Printf("[:3]: %v\n", numeros[:3])   // [0, 1, 2]
	fmt.Printf("[7:]: %v\n", numeros[7:])   // [7, 8, 9]
	fmt.Printf("[3:7]: %v\n", numeros[3:7]) // [3, 4, 5, 6]

	// CUIDADO: Los subslices comparten memoria con el original
	sub := numeros[2:5]
	sub[0] = 999
	fmt.Printf("Después de modificar subslice:\n")
	fmt.Printf("  Sub: %v\n", sub)
	fmt.Printf("  Original: %v\n", numeros) // ¡También modificado!

	// Para copia independiente, usar copy
	copia := make([]int, len(numeros))
	copy(copia, numeros)
	copia[0] = 111
	fmt.Printf("Copia independiente: %v\n", copia)
	fmt.Printf("Original sin cambios: %v\n", numeros)
}

// ============================================
// PARTE 4: Slices - Operaciones comunes
// ============================================

func operacionesSlice() {
	fmt.Println("\n=== Operaciones de Slice ===")

	// Eliminar elemento por índice
	eliminarEn := func(slice []string, i int) []string {
		return append(slice[:i], slice[i+1:]...)
	}

	opciones := []string{"Deploy", "Logs", "Status", "Config", "Exit"}
	fmt.Printf("Antes: %v\n", opciones)

	opciones = eliminarEn(opciones, 2) // Eliminar "Status"
	fmt.Printf("Después de eliminar índice 2: %v\n", opciones)

	// Insertar elemento en posición
	insertarEn := func(slice []string, i int, valor string) []string {
		slice = append(slice[:i], append([]string{valor}, slice[i:]...)...)
		return slice
	}

	opciones = insertarEn(opciones, 1, "Restart")
	fmt.Printf("Después de insertar en 1: %v\n", opciones)

	// Buscar elemento
	buscar := func(slice []string, valor string) int {
		for i, v := range slice {
			if v == valor {
				return i
			}
		}
		return -1
	}

	idx := buscar(opciones, "Config")
	fmt.Printf("Índice de 'Config': %d\n", idx)

	// Filtrar elementos
	filtrar := func(slice []string, condicion func(string) bool) []string {
		resultado := make([]string, 0)
		for _, v := range slice {
			if condicion(v) {
				resultado = append(resultado, v)
			}
		}
		return resultado
	}

	soloCortos := filtrar(opciones, func(s string) bool {
		return len(s) <= 5
	})
	fmt.Printf("Solo opciones cortas (<=5 chars): %v\n", soloCortos)
}

// ============================================
// PARTE 5: Maps - Básico
// ============================================

func mapsBasico() {
	fmt.Println("\n=== Maps Básico ===")

	// Crear map literal
	edades := map[string]int{
		"Juan":   30,
		"Ana":    25,
		"Carlos": 35,
	}
	fmt.Printf("Edades: %v\n", edades)

	// Crear map vacío con make
	estados := make(map[string]string)

	// Agregar elementos
	estados["frontend"] = "running"
	estados["backend"] = "stopped"
	estados["api"] = "running"
	fmt.Printf("Estados: %v\n", estados)

	// Acceder a valor
	fmt.Printf("Edad de Juan: %d\n", edades["Juan"])

	// Verificar si existe (IMPORTANTE)
	edad, existe := edades["Pedro"]
	if existe {
		fmt.Printf("Edad de Pedro: %d\n", edad)
	} else {
		fmt.Println("Pedro no existe en el map")
	}

	// Modificar valor
	edades["Juan"] = 31
	fmt.Printf("Nueva edad de Juan: %d\n", edades["Juan"])

	// Eliminar elemento
	delete(edades, "Carlos")
	fmt.Printf("Después de eliminar Carlos: %v\n", edades)

	// Longitud
	fmt.Printf("Cantidad de elementos: %d\n", len(estados))
}

// ============================================
// PARTE 6: Maps - Iteración
// ============================================

func mapsIteracion() {
	fmt.Println("\n=== Maps Iteración ===")

	proyectos := map[string]string{
		"frontend": "running",
		"backend":  "stopped",
		"api":      "error",
		"worker":   "running",
	}

	// Iterar con range
	fmt.Println("Todos los proyectos:")
	for nombre, estado := range proyectos {
		fmt.Printf("  %s: %s\n", nombre, estado)
	}

	// Solo claves
	fmt.Println("\nSolo nombres:")
	for nombre := range proyectos {
		fmt.Printf("  - %s\n", nombre)
	}

	// Filtrar por valor
	fmt.Println("\nProyectos running:")
	for nombre, estado := range proyectos {
		if estado == "running" {
			fmt.Printf("  - %s\n", nombre)
		}
	}

	// NOTA: El orden de iteración en maps es ALEATORIO
	// Si necesitas orden, usa slice de claves y ordena
	fmt.Println("\nOrdenado alfabéticamente:")
	claves := make([]string, 0, len(proyectos))
	for k := range proyectos {
		claves = append(claves, k)
	}
	sort.Strings(claves)
	for _, k := range claves {
		fmt.Printf("  %s: %s\n", k, proyectos[k])
	}
}

// ============================================
// PARTE 7: Maps anidados y structs
// ============================================

type Proyecto struct {
	Nombre string
	Ruta   string
	Estado string
	Puerto int
}

func mapsAvanzados() {
	fmt.Println("\n=== Maps Avanzados ===")

	// Map de structs
	proyectos := map[string]Proyecto{
		"frontend": {Nombre: "Frontend", Ruta: "/var/www/frontend", Estado: "running", Puerto: 3000},
		"backend":  {Nombre: "Backend", Ruta: "/var/www/backend", Estado: "running", Puerto: 8080},
	}

	// Acceder a struct en map
	fe := proyectos["frontend"]
	fmt.Printf("Frontend: %s en puerto %d\n", fe.Ruta, fe.Puerto)

	// Modificar struct en map (necesitas reasignar)
	be := proyectos["backend"]
	be.Estado = "stopped"
	proyectos["backend"] = be // Reasignar

	// O usar punteros
	proyectosPtrs := map[string]*Proyecto{
		"api": {Nombre: "API", Ruta: "/var/www/api", Estado: "running", Puerto: 9000},
	}
	proyectosPtrs["api"].Estado = "error" // Modifica directamente

	// Map anidado
	config := map[string]map[string]string{
		"desarrollo": {
			"host": "localhost",
			"port": "3000",
		},
		"produccion": {
			"host": "server.com",
			"port": "80",
		},
	}

	fmt.Printf("\nHost de producción: %s\n", config["produccion"]["host"])

	// Agregar a map anidado (cuidado con nil)
	config["staging"] = make(map[string]string) // Inicializar primero
	config["staging"]["host"] = "staging.com"
	config["staging"]["port"] = "8080"
	fmt.Printf("Config staging: %v\n", config["staging"])
}

// ============================================
// PARTE 8: Patrones comunes en TUI
// ============================================

func patronesTUI() {
	fmt.Println("\n=== Patrones TUI ===")

	// Menú con índice y opciones
	type MenuItem struct {
		Label  string
		Action string
	}

	menu := []MenuItem{
		{Label: "Deploy proyecto", Action: "deploy"},
		{Label: "Ver logs", Action: "logs"},
		{Label: "Reiniciar servicio", Action: "restart"},
		{Label: "Salir", Action: "quit"},
	}

	cursor := 0

	// Renderizar menú
	renderMenu := func(items []MenuItem, selected int) string {
		var sb strings.Builder
		for i, item := range items {
			if i == selected {
				sb.WriteString("> " + item.Label + " <\n")
			} else {
				sb.WriteString("  " + item.Label + "\n")
			}
		}
		return sb.String()
	}

	fmt.Println("Menú renderizado:")
	fmt.Print(renderMenu(menu, cursor))

	// Mover cursor
	cursor = 2
	fmt.Println("\nDespués de mover a índice 2:")
	fmt.Print(renderMenu(menu, cursor))

	// Obtener acción seleccionada
	accion := menu[cursor].Action
	fmt.Printf("Acción seleccionada: %s\n", accion)
}

// ============================================
// MAIN
// ============================================

func main() {
	arraysBasico()
	slicesBasico()
	slicing()
	operacionesSlice()
	mapsBasico()
	mapsIteracion()
	mapsAvanzados()
	patronesTUI()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea una función 'reverso(slice []int) []int' que retorne
	// el slice invertido sin modificar el original

	// TODO 2: Crea una función 'contarPalabras(texto string) map[string]int'
	// que cuente las ocurrencias de cada palabra

	// TODO 3: Crea una función 'unicos(slice []string) []string' que
	// retorne solo los elementos únicos (sin duplicados)

	// TODO 4: Crea una función 'agrupar(items []Proyecto) map[string][]Proyecto'
	// que agrupe proyectos por su Estado

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un sistema de caché simple:
	//
	// 1. Struct 'CacheEntry' con:
	//    - Valor string
	//    - Accesos int
	//
	// 2. Struct 'Cache' con:
	//    - datos map[string]*CacheEntry
	//    - maxSize int
	//
	// 3. Métodos:
	//    - Set(key, value string) - agrega o actualiza
	//    - Get(key string) (string, bool) - obtiene y aumenta accesos
	//    - Delete(key string)
	//    - Keys() []string - retorna todas las claves ordenadas
	//    - MasAccedidos(n int) []string - retorna las n claves más accedidas
	//    - Limpiar() - elimina entradas con 0 accesos
	//
	// Bonus: Implementa un límite de tamaño que elimine la entrada
	// menos accedida cuando se supere maxSize

	// Tu código aquí...
}
