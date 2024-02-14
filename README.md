# Go

<div style="display:flex;">
    <img src="https://upload.wikimedia.org/wikipedia/commons/c/ce/Robert_Griesemer.jpg" alt="Robert Griesemer" style="width:200px;">
    <img src="https://upload.wikimedia.org/wikipedia/commons/f/f8/Ken-Thompson-2019.png" alt="Ken Thompson" style="width:200px;">
    <img src="https://upload.wikimedia.org/wikipedia/commons/3/39/Rob-pike.jpg" alt="Rob Pike" style="width:200px;">
</div>

## Variables
Una variable es un nombre simbólico que se asocia a un valor y a un tipo de dato. Las variables se pueden declarar de varias formas en Go, usando la palabra reservada `var` o la sintaxis abreviada `:=`. Por ejemplo:

```go
// Declarar una variable con var, el nombre, el tipo y el valor inicial
var x int = 10

// Declarar una variable con var, el nombre y el valor inicial (el tipo se infiere)
var y = 20

// Declarar una variable con la sintaxis abreviada (el tipo se infiere)
z := 30
```

Las variables se pueden reasignar con el operador `=`. Por ejemplo:

```go
x = 40 // Cambiar el valor de x a 40
y = x + z // Cambiar el valor de y a la suma de x y z
```

Las variables también se pueden declarar en bloque, usando paréntesis. Por ejemplo:

```go
var (
  a string = "Hola"
  b bool = true
  c float64 = 3.14
)
```

## Operadores
Los operadores son símbolos que se usan para realizar operaciones sobre los valores de las variables. Los operadores disponibles en Go se pueden clasificar en las siguientes categorías:

- Operadores aritméticos: realizan operaciones matemáticas como suma, resta, multiplicación, división, etc. Por ejemplo:

```go
x := 10
y := 20
z := x + y // Suma
w := x - y // Resta
v := x * y // Multiplicación
u := x / y // División
t := x % y // Módulo (resto de la división)
```

- Operadores de comparación: comparan dos valores y devuelven un valor booleano (`true` o `false`). Por ejemplo:

```go
x := 10
y := 20
z := x == y // Igualdad
w := x != y // Desigualdad
v := x < y // Menor que
u := x > y // Mayor que
t := x <= y // Menor o igual que
s := x >= y // Mayor o igual que
```

- Operadores lógicos: realizan operaciones booleanas como conjunción, disyunción, negación, etc. Por ejemplo:

```go
x := true
y := false
z := x && y // AND (conjunción)
w := x || y // OR (disyunción)
v := !x // NOT (negación)
```

- Operadores de asignación: asignan un valor a una variable, o modifican el valor de una variable según una operación. Por ejemplo:

```go
x := 10 // Asignación simple
x += 5 // Asignación con suma (equivalente a x = x + 5)
x -= 5 // Asignación con resta (equivalente a x = x - 5)
x *= 5 // Asignación con multiplicación (equivalente a x = x * 5)
x /= 5 // Asignación con división (equivalente a x = x / 5)
x %= 5 // Asignación con módulo (equivalente a x = x % 5)
```

- Operadores de bits: realizan operaciones a nivel de bits, como desplazamiento, and, or, xor, etc. Por ejemplo:

```go
x := 10 // 00001010 en binario
y := 20 // 00010100 en binario
z := x & y // AND de bits (00000000 en binario, 0 en decimal)
w := x | y // OR de bits (00011110 en binario, 30 en decimal)
v := x ^ y // XOR de bits (00011110 en binario, 30 en decimal)
u := x << 2 // Desplazamiento a la izquierda de 2 bits (00101000 en binario, 40 en decimal)
t := x >> 2 // Desplazamiento a la derecha de 2 bits (00000010 en binario, 2 en decimal)
```

## Tipos de datos
Los tipos de datos son las categorías que definen el tipo, el tamaño y el formato de los valores que pueden almacenar las variables. Go es un lenguaje de tipado estático, lo que significa que el tipo de cada variable se debe definir explícita o implícitamente al declararla, y no se puede cambiar después. Los tipos de datos disponibles en Go se pueden dividir en las siguientes categorías:

- Tipos básicos: son los tipos más simples y comunes, como los números enteros, los números de coma flotante, los booleanos y las cadenas de caracteres. Por ejemplo:

```go
var x int = 10 // Un número entero de 32 o 64 bits, según la plataforma
var y float64 = 3.14 // Un número de coma flotante de 64 bits
var z bool = true // Un valor booleano (verdadero o falso)
var w string = "Go" // Una secuencia de caracteres (texto)
```

- Tipos compuestos: son los tipos que se construyen a partir de otros tipos, como los arrays, los slices, los maps, los structs y las interfaces. Por ejemplo:

```go
var x [5]int = [5]int{1, 2, 3, 4, 5} // Un array de 5 enteros
var y []int = []int{1, 2, 3} // Un slice de enteros
var z map[string]int = map[string]int{"uno": 1, "dos": 2, "tres": 3} // Un map de claves de tipo string y valores de tipo int
var w struct { // Un struct con dos campos
  nombre string
  edad int
} = struct {
  nombre: "Juan",
  edad: 25,
}
var v interface{} // Una interfaz vacía que puede contener cualquier valor
```

- Tipos derivados: son los tipos que se definen a partir de otros tipos, como los punteros, las funciones, los canales y los alias. Por ejemplo:

```go
var x int = 10 // Una variable de tipo int
var y *int = &x // Un puntero a la dirección de memoria de x
var z func(int, int) int = func(a, b int) int { // Una función que recibe dos enteros y devuelve un entero
  return a + b
}
var w chan int = make(chan int) // Un canal que puede enviar y recibir valores de tipo int
type Persona struct { // Un alias de tipo para el struct Persona
  nombre string
  edad int
}
var v Persona = Persona{ // Una variable de tipo Persona
  nombre: "Ana",
  edad: 23,
}
```

## Punteros
Un puntero es una variable que almacena la dirección de memoria de otra variable. Los punteros se pueden usar para acceder y modificar el valor de la variable apuntada, sin necesidad de hacer copias innecesarias en la memoria. Los punteros se declaran con el símbolo asterisco (*) antes del tipo de dato al que apuntan. Por ejemplo:

```go
var x int = 10 // Una variable de tipo int con el valor 10
var p *int = &x // Un puntero de tipo *int que apunta a la dirección de memoria de x
fmt.Println(*p) // Imprime el valor al que apunta p, es decir, 10
*p = 20 // Cambia el valor de la variable apuntada por p a 20
fmt.Println(x) // Imprime el valor de x, que ahora es 20
```

## Arrays
Un array es una colección ordenada y fija de elementos del mismo tipo. Los arrays se declaran con el símbolo corchete ([ ]) que indica el tamaño del array, seguido del tipo de dato de los elementos. Los elementos se pueden inicializar con una lista de valores entre llaves ({ }). Por ejemplo:

```go
var a [5]int // Un array de 5 enteros, inicializado a cero por defecto
var b [3]string = [3]string{"uno", "dos", "tres"} // Un array de 3 strings, inicializado con los valores dados
fmt.Println(a) // Imprime [0 0 0 0 0]
fmt.Println(b) // Imprime [uno dos tres]
```

## Slices
Un slice es una porción o segmento de un array, que puede variar de tamaño. Los slices se declaran con el símbolo corchete ([ ]) sin indicar el tamaño, seguido del tipo de dato de los elementos. Los elementos se pueden inicializar con una lista de valores entre llaves ({ }), o se pueden crear a partir de un array existente. Por ejemplo:

```go
var s []int // Un slice de enteros, inicialmente vacío
var t []int = []int{1, 2, 3, 4, 5} // Un slice de enteros, inicializado con los valores dados
var a [10]int = [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10} // Un array de 10 enteros
var u []int = a[2:5] // Un slice de enteros, creado a partir del array a, desde el índice 2 hasta el 4 (sin incluir el 5)
fmt.Println(s) // Imprime []
fmt.Println(t) // Imprime [1 2 3 4 5]
fmt.Println(u) // Imprime [3 4 5]
```

## Maps
Un map es una colección desordenada y dinámica de pares clave-valor, donde las claves y los valores pueden ser de cualquier tipo. Los maps se declaran con la palabra clave map, seguida del tipo de dato de las claves entre corchetes ([ ]), y el tipo de dato de los valores. Los elementos se pueden inicializar con una lista de pares clave-valor entre llaves ({ }), separados por dos puntos (:). Por ejemplo:

```go
var m map[string]int // Un map de claves de tipo string y valores de tipo int, inicialmente vacío
var n map[string]int = map[string]int{"uno": 1, "dos": 2, "tres": 3} // Un map de claves de tipo string y valores de tipo int, inicializado con los pares dados
fmt.Println(m) // Imprime map[]
fmt.Println(n) // Imprime map[uno:1 dos:2 tres:3]
```

## Structs
Un struct es una estructura de datos que agrupa campos relacionados, que pueden ser de diferentes tipos. Los structs se declaran con la palabra clave type, seguida del nombre del struct y la palabra clave struct. Los campos se definen entre llaves ({ }), con el nombre y el tipo de cada campo. Por ejemplo:

```go
type Persona struct { // Un struct llamado Persona
  nombre string // Un campo de tipo string
  edad int // Un campo de tipo int
}
var p Persona // Una variable de tipo Persona, inicializada con valores cero por defecto
var q Persona = Persona{nombre: "Juan", edad: 25} // Una variable de tipo Persona, inicializada con los valores dados
fmt.Println(p) // Imprime { 0}
fmt.Println(q) // Imprime {Juan 25}
```

## If
La instrucción `if` se usa para ejecutar un bloque de código si se cumple una condición. La sintaxis es la siguiente:

```go
if condicion {
  // código a ejecutar si la condición es verdadera
}
```

La condición debe ser una expresión booleana, es decir, que devuelva `true` o `false`. Los paréntesis alrededor de la condición son opcionales, pero las llaves son obligatorias. Por ejemplo:

```go
x := 10
if x > 0 {
  fmt.Println("x es positivo")
}
```

También se puede usar una declaración corta antes de la condición, que se ejecutará antes de evaluarla. La variable declarada en la declaración corta solo estará disponible dentro del bloque del `if`. Por ejemplo:

```go
if x := 10; x > 0 {
  fmt.Println("x es positivo")
}
```

Además, se puede usar la instrucción `else` para ejecutar otro bloque de código si la condición no se cumple. Por ejemplo:

```go
x := -10
if x > 0 {
  fmt.Println("x es positivo")
} else {
  fmt.Println("x es negativo o cero")
}
```

Finalmente, se puede usar la instrucción `else if` para encadenar varias condiciones. Por ejemplo:

```go
x := 0
if x > 0 {
  fmt.Println("x es positivo")
} else if x < 0 {
  fmt.Println("x es negativo")
} else {
  fmt.Println("x es cero")
}
```

## Switch
La instrucción `switch` se usa para comparar una expresión con varios casos posibles, y ejecutar el bloque de código correspondiente al primer caso que coincida. La sintaxis es la siguiente:

```go
switch expresion {
case valor1:
  // código a ejecutar si la expresión es igual a valor1
case valor2:
  // código a ejecutar si la expresión es igual a valor2
...
default:
  // código a ejecutar si la expresión no coincide con ningún caso
}
```

La expresión puede ser de cualquier tipo, y los casos deben ser valores constantes o únicos del mismo tipo. No es necesario usar la palabra clave `break` al final de cada caso, ya que Go lo hace automáticamente. Por ejemplo:

```go
x := 10
switch x {
case 1:
  fmt.Println("x es uno")
case 2:
  fmt.Println("x es dos")
case 3:
  fmt.Println("x es tres")
default:
  fmt.Println("x es otro número")
}
```

También se puede usar una declaración corta antes de la expresión, al igual que en el `if`. Por ejemplo:

```go
switch x := 10; x {
case 1:
  fmt.Println("x es uno")
case 2:
  fmt.Println("x es dos")
case 3:
  fmt.Println("x es tres")
default:
  fmt.Println("x es otro número")
}
```

Además, se puede omitir la expresión y usar condiciones booleanas en los casos, lo que equivale a una cadena de `else if`. Por ejemplo:

```go
x := 10
switch {
case x > 0:
  fmt.Println("x es positivo")
case x < 0:
  fmt.Println("x es negativo")
default:
  fmt.Println("x es cero")
}
```

## For
La instrucción `for` se usa para crear bucles, es decir, repetir un bloque de código mientras se cumpla una condición. La sintaxis es la siguiente:

```go
for condicion {
  // código a ejecutar mientras la condición sea verdadera
}
```

La condición debe ser una expresión booleana, y el bucle se detendrá cuando la condición sea falsa. Por ejemplo:

```go
x := 1
for x < 10 {
  fmt.Println(x)
  x++
}
```

También se puede usar una declaración de inicialización y una declaración de post-ejecución antes y después de la condición, separadas por punto y coma (;). Por ejemplo:

```go
for x := 1; x < 10; x++ {
  fmt.Println(x)
}
```

Además, se puede omitir la condición y usar solo las declaraciones de inicialización y post-ejecución, lo que crea un bucle infinito que solo se puede detener con una instrucción `break` o `return`. Por ejemplo:

```go
for x := 1; ; x++ {
  fmt.Println(x)
  if x == 10 {
    break // salir del bucle
  }
}
```

Finalmente, se puede omitir todo y crear un bucle infinito con solo la palabra clave `for`. Por ejemplo:

```go
for {
  fmt.Println("Hola")
}
```

Los ciclos for son una forma de repetir un bloque de código mientras se cumple una condición o se recorre una colección. En Go, hay diferentes formas de usar los ciclos for, dependiendo de la situación. Te daré un breve resumen de las principales formas de usar los ciclos for en Go, con algunos ejemplos.

## Ciclo for clásico
Esta es la forma más común de usar un ciclo for en Go. Se usa cuando se quiere iterar sobre un rango de valores, usando una variable de control que se inicializa, se verifica y se modifica en cada iteración. La sintaxis es la siguiente:

```go
for inicializacion; condicion; postcondicion {
  // código a ejecutar en cada iteración
}
```

Por ejemplo, si queremos imprimir los números del 1 al 10, podemos usar el siguiente ciclo for:

```go
for i := 1; i <= 10; i++ {
  fmt.Println(i)
}
```

## Ciclo for continuo (while)
Esta forma de usar un ciclo for en Go se parece a un ciclo while en otros lenguajes. Se usa cuando se quiere repetir un bloque de código mientras se cumple una condición, sin necesidad de usar una variable de control. La sintaxis es la siguiente:

```go
for condicion {
  // código a ejecutar mientras la condición sea verdadera
}
```

Por ejemplo, si queremos imprimir los números pares menores que 20, podemos usar el siguiente ciclo for:

```go
x := 0
for x < 20 {
  fmt.Println(x)
  x += 2
}
```

## Ciclo for forever
Esta forma de usar un ciclo for en Go se usa cuando se quiere repetir un bloque de código indefinidamente, hasta que se use una instrucción break o return para salir del ciclo. Se omite la condición y cualquier otra parte del ciclo for. La sintaxis es la siguiente:

```go
for {
  // código a ejecutar infinitamente
}
```

Por ejemplo, si queremos imprimir un mensaje cada segundo, podemos usar el siguiente ciclo for:

```go
import "time"
for {
  fmt.Println("Hola")
  time.Sleep(time.Second)
}
```

## Ciclo for con range-slice
Esta forma de usar un ciclo for en Go se usa cuando se quiere iterar sobre los elementos de un slice, que es una porción o segmento de un array. Se usa la palabra clave range, que devuelve el índice y el valor de cada elemento. La sintaxis es la siguiente:

```go
for indice, valor := range slice {
  // código a ejecutar con el índice y el valor de cada elemento
}
```

Por ejemplo, si queremos imprimir los elementos de un slice de strings, podemos usar el siguiente ciclo for:

```go
nombres := []string{"Ana", "Juan", "Pedro", "Luisa"}
for i, nombre := range nombres {
  fmt.Printf("El nombre en el índice %d es %s\n", i, nombre)
}
```

## Ciclo for con range-map
Esta forma de usar un ciclo for en Go se usa cuando se quiere iterar sobre los elementos de un map, que es una colección desordenada y dinámica de pares clave-valor. Se usa la palabra clave range, que devuelve la clave y el valor de cada elemento. La sintaxis es la siguiente:

```go
for clave, valor := range map {
  // código a ejecutar con la clave y el valor de cada elemento
}
```

Por ejemplo, si queremos imprimir los elementos de un map de claves de tipo string y valores de tipo int, podemos usar el siguiente ciclo for:

```go
valores := map[string]int{"A": 4, "E": 3, "I": 1, "O": 0}
for k, v := range valores {
  fmt.Printf("La clave %s tiene el valor %d\n", k, v)
}
```

## Ciclo for con range-string
Esta forma de usar un ciclo for en Go se usa cuando se quiere iterar sobre los caracteres de un string, que es una secuencia de bytes que representan texto. Se usa la palabra clave range, que devuelve el índice y el valor de cada byte. La sintaxis es la siguiente:

```go
for indice, valor := range string {
  // código a ejecutar con el índice y el valor de cada byte
}
```

Por ejemplo, si queremos imprimir los caracteres de un string, podemos usar el siguiente ciclo for:

```go
texto := "Hola"
for i, b := range texto {
  fmt.Printf("El byte en el índice %d es %d\n", i, b)
}
```

Es importante tener en cuenta que los strings en Go se codifican en UTF-8, lo que significa que cada carácter puede ocupar más de un byte. Si queremos iterar sobre los caracteres como unidades de texto, debemos convertir el string a un slice de runes, que son valores enteros que representan los puntos de código Unicode. Por ejemplo:

```go
texto := "Hola"
for i, r := range []rune(texto) {
  fmt.Printf("El rune en el índice %d es %c\n", i, r)
}
```
## Funciones

En Go, las funciones se declaran con la palabra clave `func`, seguida del nombre de la función, los parámetros entre paréntesis y los valores de retorno entre paréntesis. El cuerpo de la función se escribe entre llaves. Por ejemplo:

```go
// Una función que recibe dos enteros y devuelve su suma
func sumar(a int, b int) int {
  return a + b
}
```

Para invocar una función, se usa el nombre de la función seguido de los argumentos entre paréntesis. Por ejemplo:

```go
// Invocar la función sumar con los argumentos 3 y 5
resultado := sumar(3, 5)
fmt.Println(resultado) // Imprime 8
```

Las funciones pueden tener nombres cortos para los parámetros y los valores de retorno, o se pueden omitir si no se usan. Por ejemplo:

```go
// Una función que no recibe ni devuelve nada
func saludar() {
  fmt.Println("Hola")
}

// Una función que devuelve dos valores, sin nombres
func dividir(a, b int) (int, int) {
  return a / b, a % b
}
```

Las funciones pueden ser anónimas, es decir, que no tienen un nombre asignado. Las funciones anónimas se pueden asignar a variables o pasar como argumentos a otras funciones. Por ejemplo:

```go
// Una función anónima que se asigna a una variable
cuadrado := func(x int) int {
  return x * x
}

// Una función anónima que se pasa como argumento a otra función
sort.Slice(numeros, func(i, j int) bool {
  return numeros[i] < numeros[j]
})
```

Las funciones son valores de primera clase en Go, lo que significa que se pueden tratar como cualquier otro valor. Las funciones pueden ser parte de un tipo compuesto, como un slice, un map o un struct. Por ejemplo:

```go
// Un slice de funciones
operaciones := []func(int, int) int{sumar, restar, multiplicar, dividir}

// Un map de funciones
calculadora := map[string]func(int, int) int{"+": sumar, "-": restar, "*": multiplicar, "/": dividir}

// Un struct con un campo de tipo función
type Persona struct {
  nombre string
  saludo func()
}
```

Las funciones pueden ser recursivas, es decir, que se pueden llamar a sí mismas dentro de su cuerpo. Esto se puede usar para resolver problemas que se pueden dividir en subproblemas más pequeños. Por ejemplo:

```go
// Una función recursiva que calcula el factorial de un número
func factorial(n int) int {
  if n == 0 {
    return 1
  }
  return n * factorial(n-1)
}
```

## Funciones que devuelven múltiples valores
En Go, una función puede devolver más de un valor, separándolos con comas entre paréntesis. Esto es útil cuando se quiere devolver varios resultados relacionados o un valor y un posible error. Por ejemplo:

```go
// Una función que devuelve dos valores, el cociente y el resto de una división
func dividir(a, b int) (int, int) {
  return a / b, a % b
}

// Una función que devuelve un valor y un error, si lo hay
func abrir(archivo string) (*os.File, error) {
  f, err := os.Open(archivo)
  return f, err
}
```

Para recibir los valores devueltos por una función, se pueden usar variables separadas por comas, o un solo identificador si solo se quiere recibir uno de los valores. Por ejemplo:

```go
// Recibir los dos valores devueltos por la función dividir
cociente, resto := dividir(10, 3)
fmt.Println(cociente, resto) // Imprime 3 1

// Recibir solo el valor de error devuelto por la función abrir
_, err := abrir("archivo.txt")
if err != nil {
  fmt.Println(err) // Imprime el error, si lo hay
}
```

## Funciones con errores
En Go, los errores son valores que representan una situación anómala o fallida. Los errores se pueden crear con la función `errors.New`, que recibe un mensaje de texto y devuelve un valor de tipo `error`. Por ejemplo:

```go
// Crear un error con un mensaje
err := errors.New("algo salió mal")
fmt.Println(err) // Imprime algo salió mal
```

Los errores se pueden devolver como uno de los valores de una función, para indicar que la función no pudo completar su tarea correctamente. Por convención, el valor de error se devuelve como el último valor de la función. Por ejemplo:

```go
// Una función que devuelve un valor y un error, si lo hay
func raizCuadrada(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("no se puede calcular la raíz cuadrada de un número negativo")
  }
  return math.Sqrt(x), nil
}
```

Para manejar los errores devueltos por una función, se puede usar una instrucción `if` para verificar si el valor de error es distinto de `nil`, que significa que hay un error. Por ejemplo:

```go
// Invocar la función raizCuadrada con un argumento válido
r, err := raizCuadrada(25)
if err != nil {
  fmt.Println(err) // No se ejecuta, porque no hay error
} else {
  fmt.Println(r) // Imprime 5
}

// Invocar la función raizCuadrada con un argumento inválido
r, err = raizCuadrada(-25)
if err != nil {
  fmt.Println(err) // Imprime no se puede calcular la raíz cuadrada de un número negativo
} else {
  fmt.Println(r) // No se ejecuta, porque hay error
}
```

## Funciones que retornan y reciben funciones
En Go, las funciones son valores de primera clase, lo que significa que se pueden tratar como cualquier otro valor. Esto implica que las funciones pueden ser devueltas por otras funciones, o recibidas como argumentos de otras funciones. Por ejemplo:

```go
// Una función que devuelve otra función
func crearSaludo(nombre string) func() string {
  return func() string {
    return "Hola, " + nombre
  }
}

// Una función que recibe otra función como argumento
func aplicarFuncion(f func(int) int, x int) int {
  return f(x)
}
```

Para usar las funciones que retornan o reciben funciones, se puede asignar la función devuelta a una variable, o invocar la función directamente. Por ejemplo:

```go
// Asignar la función devuelta por crearSaludo a una variable
saludar := crearSaludo("Juan")
fmt.Println(saludar()) // Imprime Hola, Juan

// Invocar la función devuelta por crearSaludo directamente
fmt.Println(crearSaludo("Ana")()) // Imprime Hola, Ana

// Asignar una función anónima a una variable
cuadrado := func(x int) int {
  return x * x
}

// Pasar la función cuadrado como argumento a la función aplicarFuncion
fmt.Println(aplicarFuncion(cuadrado, 5)) // Imprime 25

// Pasar una función anónima directamente como argumento a la función aplicarFuncion
fmt.Println(aplicarFuncion(func(x int) int {
  return x * 2
}, 5)) // Imprime 10
```

## Función variádica
Una función variádica es una función que puede recibir un número variable de argumentos del mismo tipo. Para declarar una función variádica, se usa el símbolo de puntos suspensivos (...) antes del tipo del último parámetro. Por ejemplo:

```go
// Una función variádica que recibe un número variable de enteros
func sumar(numeros ...int) int {
  suma := 0
  for _, n := range numeros {
    suma += n
  }
  return suma
}
```

Para invocar una función variádica, se pueden pasar los argumentos separados por comas, o se puede pasar un slice de valores con el símbolo de puntos suspensivos al final. Por ejemplo:

```go
// Invocar la función sumar con argumentos separados por comas
fmt.Println(sumar(1, 2, 3)) // Imprime 6

// Invocar la función sumar con un slice de valores
numeros := []int{4, 5, 6}
fmt.Println(sumar(numeros...)) // Imprime 15
```

## Funciones anónimas
Las funciones anónimas son funciones que no tienen un nombre asignado. Las funciones anónimas se pueden asignar a variables o pasar como argumentos a otras funciones. Las funciones anónimas se declaran con la palabra clave `func`, seguida de los parámetros y los valores de retorno entre paréntesis, y el cuerpo de la función entre llaves. Por ejemplo:

```go
// Una función anónima que se asigna a una variable
cuadrado := func(x int) int {
  return x * x
}

// Una función anónima que se pasa como argumento a otra función
sort.Slice(numeros, func(i, j int) bool {
  return numeros[i] < numeros[j]
})
```

Para invocar una función anónima, se usa el nombre de la variable que la contiene, o se usa la palabra clave `func` seguida de los argumentos entre paréntesis. Por ejemplo:

```go
// Invocar la función anónima asignada a la variable cuadrado
fmt.Println(cuadrado(5)) // Imprime 25

// Invocar una función anónima directamente
func(x int) {
  fmt.Println(x * 2)
}(5) // Imprime 10
```

---

## Go Modules

Go es un lenguaje de programación de alto nivel, compilado, concurrente y con tipado estático, diseñado por Google. ¹² Su sintaxis es similar a C, pero con características como la seguridad de memoria, la recolección de basura, el tipado estructural y la concurrencia estilo CSP. ² Go tiene un sistema de gestión de paquetes y dependencias integrado, que se basa en los siguientes conceptos:

- **Módulos**: son colecciones de paquetes relacionados que se versionan juntos. Un módulo se define por un archivo `go.mod` que especifica su nombre, versión y dependencias. Los módulos pueden ser publicados en repositorios remotos o locales. ³
- **Paquetes**: son unidades de código fuente que se pueden importar y usar en otros paquetes. Un paquete se compone de uno o más archivos `.go` que pertenecen al mismo directorio. Cada paquete tiene un nombre, que se usa para referenciarlo desde otros paquetes. Los paquetes pueden ser estándar (incluidos en la biblioteca de Go) o de terceros (desarrollados por la comunidad). ⁴
- **Dependencias**: son los paquetes que un paquete o un módulo necesita para funcionar correctamente. Las dependencias se declaran en el archivo `go.mod` de un módulo, usando la directiva `require`. Cada dependencia tiene un nombre (el import path del paquete) y una versión (un número semántico o un identificador de commit). ³

Para trabajar con paquetes y dependencias en Go, se utiliza la herramienta `go` desde la línea de comandos. Algunos de los comandos más importantes son:

- `go mod init`: crea un nuevo módulo e inicializa el archivo `go.mod` con el nombre y la versión del módulo. ⁵
- `go get`: descarga e instala un paquete o un módulo, y actualiza el archivo `go.mod` con la dependencia correspondiente. También se puede usar para actualizar o cambiar la versión de una dependencia existente, usando el flag `-u` o especificando la versión deseada. 
- `go mod tidy`: elimina las dependencias que no se usan en el código fuente, y añade las que faltan. También actualiza el archivo `go.sum`, que contiene los hashes criptográficos de las dependencias para garantizar su integridad. 

A continuación se muestran algunos ejemplos de cómo usar estos comandos:

```go
// Crear un nuevo módulo llamado example.com/hello
$ go mod init example.com/hello
go: creating new go.mod: module example.com/hello

// Importar y usar el paquete rsc.io/quote en el archivo main.go
package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	fmt.Println(quote.Hello())
}

// Descargar e instalar el paquete rsc.io/quote y sus dependencias
$ go get rsc.io/quote
go: downloading rsc.io/quote v1.5.2
go: downloading rsc.io/sampler v1.3.0
go: downloading golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c

// Ver el contenido del archivo go.mod
module example.com/hello

go 1.16

require rsc.io/quote v1.5.2

// Actualizar el paquete rsc.io/quote a la versión v1.5.3
$ go get rsc.io/quote@v1.5.3
go: downloading rsc.io/quote v1.5.3
go: downloading rsc.io/sampler v1.99.99
go: downloading golang.org/x/text v0.0.0-20170915090833-1cbadb444a80

// Ver el contenido del archivo go.mod
module example.com/hello

go 1.16

require rsc.io/quote v1.5.3

// Eliminar las dependencias que no se usan y actualizar el archivo go.sum
$ go mod tidy
go: downloading github.com/golang/protobuf v1.2.0
```
