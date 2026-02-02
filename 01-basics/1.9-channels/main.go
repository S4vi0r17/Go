// Ejercicio 9: Channels
// Objetivo: Dominar channels (comunicación entre goroutines) - CRÍTICO PARA TUIs

package main

import (
	"fmt"
	"math/rand"
	"time"
)

// ============================================
// PARTE 1: Channel básico
// ============================================

func demoChannelBasico() {
	fmt.Println("=== Channel Básico ===")

	// Crear channel
	mensajes := make(chan string)

	// Goroutine que envía
	go func() {
		time.Sleep(500 * time.Millisecond)
		mensajes <- "¡Hola desde goroutine!" // Enviar
	}()

	// Recibir (bloquea hasta que haya un mensaje)
	msg := <-mensajes
	fmt.Println("Recibido:", msg)
}

// ============================================
// PARTE 2: Channel con buffer
// ============================================

func demoChannelBuffer() {
	fmt.Println("\n=== Channel con Buffer ===")

	// Channel sin buffer - envío bloquea hasta que alguien reciba
	sinBuffer := make(chan int)

	// Channel con buffer - puede almacenar N valores sin bloquear
	conBuffer := make(chan int, 3)

	// Con buffer, podemos enviar sin receptor inmediato
	conBuffer <- 1
	conBuffer <- 2
	conBuffer <- 3
	// conBuffer <- 4 // Esto bloquearía (buffer lleno)

	fmt.Printf("Buffer tiene %d elementos\n", len(conBuffer))

	// Recibir
	fmt.Println("Recibidos:", <-conBuffer, <-conBuffer, <-conBuffer)

	// Demo sin buffer requiere goroutine
	go func() {
		sinBuffer <- 42
	}()
	fmt.Println("Sin buffer recibido:", <-sinBuffer)
}

// ============================================
// PARTE 3: Cerrar channels
// ============================================

func demoChannelClose() {
	fmt.Println("\n=== Cerrar Channels ===")

	numeros := make(chan int)

	// Productor
	go func() {
		for i := 1; i <= 5; i++ {
			numeros <- i
			time.Sleep(100 * time.Millisecond)
		}
		close(numeros) // Cerrar cuando terminamos de enviar
	}()

	// Consumidor - range itera hasta que el channel se cierre
	fmt.Print("Números: ")
	for n := range numeros {
		fmt.Printf("%d ", n)
	}
	fmt.Println()

	// Verificar si está cerrado
	canal := make(chan string, 1)
	canal <- "mensaje"
	close(canal)

	val, ok := <-canal
	fmt.Printf("Valor: %s, Abierto: %v\n", val, ok)

	val, ok = <-canal // Channel cerrado y vacío
	fmt.Printf("Valor: '%s', Abierto: %v\n", val, ok)
}

// ============================================
// PARTE 4: Select (multiplexar channels)
// ============================================

func demoSelect() {
	fmt.Println("\n=== Select ===")

	canal1 := make(chan string)
	canal2 := make(chan string)

	go func() {
		time.Sleep(100 * time.Millisecond)
		canal1 <- "mensaje de canal1"
	}()

	go func() {
		time.Sleep(200 * time.Millisecond)
		canal2 <- "mensaje de canal2"
	}()

	// Select espera en múltiples channels
	for i := 0; i < 2; i++ {
		select {
		case msg := <-canal1:
			fmt.Println("Recibido de canal1:", msg)
		case msg := <-canal2:
			fmt.Println("Recibido de canal2:", msg)
		}
	}
}

// ============================================
// PARTE 5: Select con timeout
// ============================================

func demoSelectTimeout() {
	fmt.Println("\n=== Select con Timeout ===")

	lento := make(chan string)

	go func() {
		time.Sleep(2 * time.Second) // Muy lento
		lento <- "respuesta"
	}()

	select {
	case msg := <-lento:
		fmt.Println("Recibido:", msg)
	case <-time.After(500 * time.Millisecond):
		fmt.Println("Timeout! La operación tardó demasiado")
	}
}

// ============================================
// PARTE 6: Select con default (no bloqueante)
// ============================================

func demoSelectDefault() {
	fmt.Println("\n=== Select con Default ===")

	canal := make(chan string)

	// Intento no bloqueante de recibir
	select {
	case msg := <-canal:
		fmt.Println("Recibido:", msg)
	default:
		fmt.Println("No hay mensajes disponibles")
	}

	// Intento no bloqueante de enviar
	select {
	case canal <- "mensaje":
		fmt.Println("Enviado exitosamente")
	default:
		fmt.Println("No se pudo enviar (nadie escuchando)")
	}
}

// ============================================
// PARTE 7: Patrón Pipeline
// ============================================

func demoPipeline() {
	fmt.Println("\n=== Pipeline ===")

	// Etapa 1: Generador
	generar := func(nums ...int) <-chan int {
		out := make(chan int)
		go func() {
			for _, n := range nums {
				out <- n
			}
			close(out)
		}()
		return out
	}

	// Etapa 2: Cuadrado
	cuadrado := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				out <- n * n
			}
			close(out)
		}()
		return out
	}

	// Etapa 3: Filtrar pares
	filtrarPares := func(in <-chan int) <-chan int {
		out := make(chan int)
		go func() {
			for n := range in {
				if n%2 == 0 {
					out <- n
				}
			}
			close(out)
		}()
		return out
	}

	// Conectar pipeline
	numeros := generar(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)
	cuadrados := cuadrado(numeros)
	pares := filtrarPares(cuadrados)

	fmt.Print("Resultado: ")
	for n := range pares {
		fmt.Printf("%d ", n)
	}
	fmt.Println()
}

// ============================================
// PARTE 8: Patrón Worker Pool
// ============================================

type Tarea struct {
	ID      int
	Trabajo string
}

type Resultado struct {
	TareaID  int
	Salida   string
	Duracion time.Duration
}

func worker(id int, tareas <-chan Tarea, resultados chan<- Resultado) {
	for t := range tareas {
		inicio := time.Now()

		// Simular trabajo
		time.Sleep(time.Duration(100+rand.Intn(200)) * time.Millisecond)

		resultados <- Resultado{
			TareaID:  t.ID,
			Salida:   fmt.Sprintf("Worker %d procesó: %s", id, t.Trabajo),
			Duracion: time.Since(inicio),
		}
	}
}

func demoWorkerPool() {
	fmt.Println("\n=== Worker Pool ===")

	numWorkers := 3
	numTareas := 10

	tareas := make(chan Tarea, numTareas)
	resultados := make(chan Resultado, numTareas)

	// Iniciar workers
	for w := 1; w <= numWorkers; w++ {
		go worker(w, tareas, resultados)
	}

	// Enviar tareas
	for i := 1; i <= numTareas; i++ {
		tareas <- Tarea{
			ID:      i,
			Trabajo: fmt.Sprintf("Tarea-%d", i),
		}
	}
	close(tareas)

	// Recoger resultados
	for i := 0; i < numTareas; i++ {
		r := <-resultados
		fmt.Printf("  [%v] %s\n", r.Duracion.Round(time.Millisecond), r.Salida)
	}
}

// ============================================
// PARTE 9: Stop Channel (para TUIs)
// ============================================

func streamLogs(stopCh <-chan struct{}) <-chan string {
	logs := make(chan string)

	go func() {
		defer close(logs)
		contador := 0

		for {
			select {
			case <-stopCh:
				fmt.Println("  [Stream detenido]")
				return
			default:
				contador++
				logs <- fmt.Sprintf("Log #%d: evento aleatorio %d", contador, rand.Intn(100))
				time.Sleep(200 * time.Millisecond)
			}
		}
	}()

	return logs
}

func demoStopChannel() {
	fmt.Println("\n=== Stop Channel (patrón TUI) ===")

	stopCh := make(chan struct{})
	logs := streamLogs(stopCh)

	// Leer logs por 1 segundo
	timeout := time.After(1 * time.Second)

loop:
	for {
		select {
		case log, ok := <-logs:
			if !ok {
				break loop
			}
			fmt.Println(" ", log)
		case <-timeout:
			fmt.Println("  [1 segundo - enviando stop]")
			close(stopCh) // Señalar que pare
		}
	}

	fmt.Println("  [Fin del demo]")
}

// ============================================
// PARTE 10: Ticker para actualizaciones periódicas
// ============================================

func demoTicker() {
	fmt.Println("\n=== Ticker (actualizaciones periódicas) ===")

	ticker := time.NewTicker(300 * time.Millisecond)
	defer ticker.Stop()

	done := make(chan bool)

	go func() {
		time.Sleep(1 * time.Second)
		done <- true
	}()

	contador := 0
	for {
		select {
		case <-done:
			fmt.Println("  [Terminado]")
			return
		case t := <-ticker.C:
			contador++
			fmt.Printf("  Tick #%d a las %v\n", contador, t.Format("15:04:05.000"))
		}
	}
}

// ============================================
// MAIN
// ============================================

func main() {
	rand.Seed(time.Now().UnixNano())

	demoChannelBasico()
	demoChannelBuffer()
	demoChannelClose()
	demoSelect()
	demoSelectTimeout()
	demoSelectDefault()
	demoPipeline()
	demoWorkerPool()
	demoStopChannel()
	demoTicker()

	// ============================================
	// EJERCICIOS
	// ============================================

	fmt.Println("\n=== TU TURNO ===")

	// TODO 1: Crea un 'Fan-Out/Fan-In' donde:
	// - Un generador produce números
	// - N workers los procesan en paralelo
	// - Un colector los agrupa

	// TODO 2: Implementa un 'Rate Limiter' usando time.Ticker
	// que solo permita N operaciones por segundo

	// TODO 3: Crea un 'Broadcast' donde un emisor envía a
	// múltiples receptores

	// ============================================
	// RETO FINAL
	// ============================================

	fmt.Println("\n=== RETO FINAL ===")

	// Crea un sistema de "Log Aggregator" para TUI:
	//
	// 1. Struct 'LogEntry' con:
	//    - Timestamp time.Time
	//    - Source string (frontend, backend, etc.)
	//    - Level string (info, warn, error)
	//    - Message string
	//
	// 2. Función 'simulateLogSource(name string, logs chan<- LogEntry, stop <-chan struct{})'
	//    que genere logs aleatorios cada 100-500ms hasta recibir stop
	//
	// 3. Función 'aggregateLogs(sources []<-chan LogEntry) <-chan LogEntry'
	//    que combine logs de múltiples fuentes en un solo channel
	//
	// 4. Función 'filterLogs(logs <-chan LogEntry, level string) <-chan LogEntry'
	//    que solo deje pasar logs del nivel especificado
	//
	// 5. Struct 'LogBuffer' con:
	//    - Buffer de últimos N logs (circular)
	//    - Método Add(LogEntry)
	//    - Método GetAll() []LogEntry
	//    - Mutex para thread-safety
	//
	// 6. Main que:
	//    - Inicie 3 fuentes de logs
	//    - Las agregue en un canal
	//    - Las guarde en el buffer
	//    - Cada 500ms imprima los últimos 5 logs
	//    - Después de 3 segundos, detenga todo

	// Tu código aquí...
}
