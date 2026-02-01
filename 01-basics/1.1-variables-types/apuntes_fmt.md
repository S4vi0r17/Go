# Apuntes: `fmt` — verbs y cómo escapar `%`

Ubicación: `Go/01-basics/1.2-variables-types` (junto a tu `main.go`).

Pequeño resumen en formato preguntas y respuestas con ejemplos. Para ejecutar los ejemplos localmente:

1. Añade `FmtApuntesExamples()` dentro de tu `main()` o llama a esa función desde donde quieras.
2. Ejecuta `go run .` en esta carpeta.

---

### P: ¿Qué hace `%T`? ¿Y otros verbs comunes?

R: `%T` imprime el tipo concreto de la variable (por ejemplo `int`, `string`, `*int`).
Algunos verbs útiles:

- `%v` valor por defecto
- `%+v` como `%v`, pero en structs incluye nombres de campo
- `%#v` sintaxis Go
- `%T` tipo
- `%%` literal `%`
- `%t` booleano
- `%d`, `%b`, `%o`, `%x`, `%X` enteros (decimal, binario, octal, hex)
- `%c`, `%q`, `%U` runes/caracteres/Unicode
- `%f`, `%e`, `%g`, `%b` floats
- `%s`, `%q`, `%x` cadenas / bytes
- `%p` punteros

Ejemplo (ver tipo de variables):

```go
package main

import "fmt"

func main() {
	nombre := "servidor"
	puerto := 8080
	activo := true

	fmt.Printf("Tipos: nombre=%T, puerto=%T, activo=%T\n", nombre, puerto, activo)
	// Salida esperada: string int bool
}
```

---

### P: ¿Cómo imprimo `%T` literalmente (sin interpretarlo)?

R: Opciones:

1. Escapar `%` con `%%`: `fmt.Printf("%%T\n")` → imprime `%T`.
2. Pasarlo como argumento: `fmt.Printf("%s\n", "%T")`.
3. Usar `fmt.Print`/`fmt.Println` (no interpretan verbs): `fmt.Print("%T\n")`.

Ejemplo (variantes):

```go
package main

import "fmt"

func main() {
	// 1) Escapando % con %%
	fmt.Printf("Escapando %% para mostrar '%%T': %%T\n")

	// 2) Pasándolo como argumento
	fmt.Printf("Pasando como argumento: %s\n", "%T")

	// 3) Usando fmt.Print (no interpreta formato)
	fmt.Print("Usando fmt.Print: ")
	fmt.Print("%T\n")
}
```

Nota: si tu formato viene de una cadena dinámica, puedes sanearla con `strings.ReplaceAll(fmtStr, "%", "%%")` para que no se interpreten verbs.

---

### P: ¿Cómo controlo ancho y precisión para floats?

R: Usa `%.2f` para dos decimales, `%6.2f` para ancho mínimo 6 y 2 decimales, `%g` para formato compacto.
Ejemplo:

```go
package main

import "fmt"

func main() {
	pi := 3.1415926535
	fmt.Printf("Float (%%f): %f\n", pi)
	fmt.Printf("Float (precisión 2): %.2f\n", pi)
	fmt.Printf("Float (%%g compacto): %g\n", pi)
	fmt.Printf("Formato ancho y precisión (%%6.2f): %6.2f\n", pi)
}
```

---

### P: ¿Qué pasa si uso un verb incorrecto o me faltan argumentos?

R: `fmt` inserta mensajes especiales que empiezan por `%!` describiendo el error (por ejemplo `%!d(string=valor)` o `%!d(MISSING)`).
Ejemplo:

```go
package main

import "fmt"

func main() {
	// Ejemplo de error de formato (salida mostrará %!verb...):
	fmt.Printf("%d\n", "esto-no-es-un-int") // genera: %!d(string=esto-no-es-un-int)
}
```

---

Resumen rápido (cheat‑sheet):

- `%T` → tipo
- `%%` → `%` literal
- `%v / %+v / %#v` → por defecto / con nombres / Go-syntax
- `%d %b %o %x` → enteros (dec, bin, oct, hex)
- `%.2f / %6.2f / %g` → floats y precisión
- Flags: `+` (siempre signo), `-` (izquierda), `0` (relleno), `#` (formato alternativo), espacio (coloca espacio para signo)
