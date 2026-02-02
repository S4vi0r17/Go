## Lo que me falta

Estas son las sesiones que todavía te faltan leer. Para cada sesión incluyo una breve descripción, por qué es importante y una recomendación de lectura/acción. Marca la casilla cuando la completes (cambia `[ ]` por `[x]`).

- [ ] 1.7 - Manejo de Errores (`01-basics/1.7-error-handling`)
  - Objetivo: Dominar el patrón de errores en Go (no hay excepciones).
  - Qué incluye: retorno `(valor, error)`, errores como variables, wrapping (`%w`), `errors.Is` / `errors.As`, errores personalizados y manejo de múltiples errores.
  - Recomendación: abre `01-basics/1.7-error-handling/main.go`, lee los comentarios y ejecuta los ejemplos. Prueba funciones como `dividir`, `validarPuerto` y `validarConfiguracion`.
  - Tiempo estimado: 45–60 min. Prioridad: Alta.

- [ ] 1.8 - Goroutines (`01-basics/1.8-goroutines`)
  - Objetivo: Entender concurrencia básica en Go (goroutines y sincronización).
  - Qué incluye: goroutines, `sync.WaitGroup`, `Mutex`/`RWMutex`, condiciones de carrera y patrones como pools de workers.
  - Recomendación: revisa `01-basics/1.8-goroutines/main.go`, ejecuta los demos y experimenta cambiando número de goroutines y sleeps para observar comportamientos.
  - Tiempo estimado: 45–90 min. Prioridad: Crítica (recomendado para TUI / concurrencia).

- [ ] 1.9 - Channels (`01-basics/1.9-channels`)
  - Objetivo: Dominar `channels` para comunicación entre goroutines.
  - Qué incluye: channels básicos y con buffer, cerrar channels, `select`, fan-in/fan-out y pipelines.
  - Recomendación: abre `01-basics/1.9-channels/main.go`, ejecuta las demos y practica `select` y el cierre de channels.
  - Tiempo estimado: 45–90 min. Prioridad: Crítica.

- [ ] 02 - CLI básico (`02-cli-basico/`) — Módulo (varias sesiones)
  - Sub-sesiones:
    - [ ] 2.1 - `os.Args` (`02-cli-basico/2.1-args-basico`) — argumentos de línea de comandos.
    - [ ] 2.2 - `flag` (`02-cli-basico/2.2-flags`) — parsing y uso de flags.
    - [ ] 2.3 - Stdin/Stdout (`02-cli-basico/2.3-stdin-stdout`) — lectura (`bufio.Scanner`) y `stdout` vs `stderr`.
    - [ ] 2.4 - Colores ANSI (`02-cli-basico/2.4-colores-ansi`) — estilos y colores para terminal.
    - [ ] 2.5 - Subcomandos (`02-cli-basico/2.5-subcomandos`) — estructura y ayuda por comando.
    - [ ] 2.6 - Proyecto Deploy (`02-cli-basico/2.6-proyecto-deploy`) — proyecto integrador que aplica lo anterior.
  - Recomendación: recorre los `main.go` de cada subcarpeta, ejecuta los ejemplos y marca cada sub-sesión al completarla.

Sugerencias prácticas:

- Ejecuta `go run main.go` dentro de la carpeta de cada sesión para ver los demos en acción.
- Si tienes poco tiempo, prioriza `1.8` y `1.9` (concurrencia) y `1.7` (errores).
- Cuando completes una sesión, cambia la casilla a `[x]` para mantener la lista actualizada.
