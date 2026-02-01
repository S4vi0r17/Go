package main

import "fmt"

// Definición
type Config struct {
	Host    string
	Puerto  int
	Usuario string
	Activo  bool
}

// Structs anidados
type App struct {
	Nombre string
	Config Config
	Estado string
}

func main() {
	config := Config{
		Host:    "localhost",
		Puerto:  22,
		Usuario: "admin",
		Activo:  true,
	}

	app := App{
		Nombre: "MiApp",
		Config: config,
		Estado: "Activo",
	}

	fmt.Println(app.Nombre)
	fmt.Println(app.Config.Host)
	fmt.Println(app.Estado)

	fmt.Printf("\nServer running on %s:%d\n", config.Host, config.Puerto)
}
