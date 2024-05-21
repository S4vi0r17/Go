package main

import (
	"fmt"

	"github.com/S4vi0r17/Go/course"
)

func main() {
	Go := course.New("Go desde cero", "go-desde-cero", []string{"backend"}, 12.34, false, map[uint]string{
		1: "Introducción",
		2: "Estructuras",
		3: "Maps",
	})

	Go.SetUserIDs([]uint{1, 2, 3, 4, 5})

	fmt.Println(Go)
	
	Go.PrintClasses()
	// Go.ChangePrice(15.5) Go.ChangePrice undefined
	fmt.Println(Go.Price())
	Go.SetName("Go desde cero 2")
	fmt.Println(Go.Name())
}
