// Ejercicio 8: Goroutines
// Objetivo: Dominar concurrencia básica (CRÍTICO para TUIs)

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// ============================================
// PARTE 1: Goroutine básica
// ============================================

func tareaSimple(id int) {
	fmt.Printf("Tarea %d: iniciando\n", id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Tarea %d: completada\n", id)
}

func demoGoroutineBasica() {
	fmt.Println("=== Goroutine Básica ===")

	// Sin goroutine (secuencial)
	fmt.Println("\nSecuencial:")
	inicio := time.Now()
	tareaSimple(1)
	tareaSimple(2)
	fmt.Printf("Tiempo: %v\n", time.Since(inicio))

	// Con goroutines (paralelo)
	fmt.Println("\nParalelo:")
	inicio = time.Now()
	go tareaSimple(1)
	go tareaSimple(2)
	go tareaSimple(3)

	// PROBLEMA: El programa termina antes que las goroutines
	// Necesitamos esperar...
	time.Sleep(2 * time.Second)
	fmt.Printf("Tiempo: %v\n", time.Since(inicio))
}

// ============================================
// PARTE 2: WaitGroup para sincronizar
// ============================================

func tareaConWG(id int, wg *sync.WaitGroup) {
	defer wg.Done() // Señalar que terminó

	fmt.Printf("Tarea %d: iniciando\n", id)
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	fmt.Printf("Tarea %d: completada\n", id)
}

func demoWaitGroup() {
	fmt.Println("\n=== WaitGroup ===")

	var wg sync.WaitGroup

	for i := 1; i <= 5; i++ {
		wg.Add(1) // Agregar contador
		go tareaConWG(i, &wg)
	}

	fmt.Println("Esperando que terminen todas las tareas...")
	wg.Wait() // Bloquea hasta que el contador llegue a 0
	fmt.Println("¡Todas las tareas completadas!")
}

// ============================================
// PARTE 3: Goroutines anónimas
// ============================================

func demoGoroutineAnonima() {
	fmt.Println("\n=== Goroutines Anónimas ===")

	var wg sync.WaitGroup

	// Goroutine anónima con closure
	mensaje := "Hola desde goroutine"

	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println(mensaje)
	}()

	// CUIDADO con closures en loops
	fmt.Println("\nProblema común (closure en loop):")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// MAL: i es compartido, todas verán el último valor
		go func() {
			defer wg.Done()
			fmt.Printf("(mal) i = %d\n", i)
		}()
	}

	wg.Wait()

	fmt.Println("\nSolución (pasar como parámetro):")
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		// BIEN: pasar i como parámetro copia el valor
		go func(id int) {
			defer wg.Done()
			fmt.Printf("(bien) i = %d\n", id)
		}(i)
	}

	wg.Wait()
}

// ============================================
// PARTE 4: Caso de uso - Tareas paralelas
// ============================================

type ResultadoDeploy struct {
	Proyecto string
	Exito    bool
	Duracion time.Duration
	Error    string
}

func desplegarProyecto(nombre string) ResultadoDeploy {
	inicio := time.Now()

	// Simular deploy
	duracion := time.Duration(500+rand.Intn(1500)) * time.Millisecond
	time.Sleep(duracion)

	// Simular éxito/fallo aleatorio
	exito := rand.Float32() > 0.2 // 80% éxito

	resultado := ResultadoDeploy{
		Proyecto: nombre,
		Exito:    exito,
		Duracion: time.Since(inicio),
	}

	if !exito {
		resultado.Error = "Error de compilación"
	}

	return resultado
}

func demoDeployParalelo() {
	fmt.Println("\n=== Deploy Paralelo ===")

	proyectos := []string{"frontend", "backend", "api", "worker", "admin"}

	// Secuencial
	fmt.Println("\nDeploy secuencial:")
	inicio := time.Now()
	for _, p := range proyectos {
		r := desplegarProyecto(p)
		estado := "OK"
		if !r.Exito {
			estado = "FAIL"
		}
		fmt.Printf("  %s %s (%v)\n", estado, r.Proyecto, r.Duracion)
	}
	fmt.Printf("Total: %v\n", time.Since(inicio))

	// Paralelo
	fmt.Println("\nDeploy paralelo:")
	inicio = time.Now()

	resultados := make([]ResultadoDeploy, len(proyectos))
	var wg sync.WaitGroup

	for i, p := range proyectos {
		wg.Add(1)
		go func(idx int, proyecto string) {
			defer wg.Done()
			resultados[idx] = desplegarProyecto(proyecto)
		}(i, p)
	}

	wg.Wait()

	for _, r := range resultados {
		estado := "OK"
		if !r.Exito {
			estado = "FAIL"
		}
		fmt.Printf("  %s %s (%v)\n", estado, r.Proyecto, r.Duracion)
	}
	fmt.Printf("Total: %v\n", time.Since(inicio))
}

// ============================================
// PARTE 5: Mutex para datos compartidos
// ============================================

// Sin mutex (DATA RACE - peligroso)
type ContadorInseguro struct {
	valor int
}

func (c *ContadorInseguro) Incrementar() {
	c.valor++
}

// Con mutex (seguro)
type ContadorSeguro struct {
	valor int
	mu    sync.Mutex
}

func (c *ContadorSeguro) Incrementar() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.valor++
}

func (c *ContadorSeguro) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valor
}

func demoMutex() {
	fmt.Println("\n=== Mutex ===")

	// Probar contador inseguro
	fmt.Println("\nContador inseguro (puede dar resultados incorrectos):")
	inseguro := &ContadorInseguro{}
	var wg sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			inseguro.Incrementar()
		}()
	}
	wg.Wait()
	fmt.Printf("Esperado: 1000, Obtenido: %d\n", inseguro.valor)

	// Probar contador seguro
	fmt.Println("\nContador seguro (siempre correcto):")
	seguro := &ContadorSeguro{}

	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			seguro.Incrementar()
		}()
	}
	wg.Wait()
	fmt.Printf("Esperado: 1000, Obtenido: %d\n", seguro.Valor())
}

// ============================================
// PARTE 6: RWMutex para lecturas concurrentes
// ============================================

type Cache struct {
	datos map[string]string
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		datos: make(map[string]string),
	}
}

func (c *Cache) Set(key, value string) {
	c.mu.Lock() // Lock exclusivo para escribir
	defer c.mu.Unlock()
	c.datos[key] = value
}

func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock() // Lock de lectura (múltiples lectores OK)
	defer c.mu.RUnlock()
	val, ok := c.datos[key]
	return val, ok
}

func demoRWMutex() {
	fmt.Println("\n=== RWMutex ===")

	cache := NewCache()
	var wg sync.WaitGroup

	// Escritores
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id)
			cache.Set(key, fmt.Sprintf("valor%d", id))
			fmt.Printf("Escrito: %s\n", key)
		}(i)
	}

	// Lectores (pueden correr en paralelo)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			key := fmt.Sprintf("key%d", id%5)
			if val, ok := cache.Get(key); ok {
				fmt.Printf("Leído: %s = %s\n", key, val)
			}
		}(i)
	}

	wg.Wait()
}

// ============================================
// MAIN
// ============================================

func main() {
	rand.Seed(time.Now().UnixNano())

	demoGoroutineBasica()
	demoWaitGroup()
	demoGoroutineAnonima()
	demoDeployParalelo()
	demoMutex()
	demoRWMutex()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea una función 'buscarEnParalelo' que busque un string
	// en múltiples archivos simultáneamente y retorne el primero encontrado

	// TODO 2: Crea un 'Pool de Workers' con N goroutines que procesen
	// tareas de un slice

	// TODO 3: Implementa un 'rate limiter' que solo permita N operaciones
	// por segundo usando time.Ticker

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un "Health Checker" que:
	//
	// 1. Tenga una lista de servicios (nombre, url)
	//
	// 2. Función 'checkService(url string) (bool, time.Duration)'
	//    que simule una verificación de salud
	//
	// 3. Función 'checkAllServices(services []Service) []HealthResult'
	//    que verifique todos los servicios EN PARALELO
	//
	// 4. Struct 'HealthResult' con:
	//    - Servicio string
	//    - Saludable bool
	//    - Latencia time.Duration
	//    - Error string
	//
	// 5. Usa WaitGroup y Mutex para:
	//    - Esperar a que terminen todas las verificaciones
	//    - Almacenar resultados de forma segura
	//
	// 6. Implementa un timeout: si un servicio tarda más de 2s,
	//    márcalo como no saludable

	// Tu código aquí...
}
