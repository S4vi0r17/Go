package main

import (
	"github.com/S4vi0r17/Go/course"
)

func main() {
	Go := &course.Course{
		Name:    "Go desde cero",
		Slug:    "go-desde-cero",
		Skills:  []string{"backend"},
		Price:   12.34,
		IsFree:  false,
		UserIDs: []uint{12, 56, 89},
		Classes: map[uint]string{
			1: "Introducción",
			2: "Estructuras",
			3: "Maps",
		},
	}

	Go.PrintClasses()
	Go.
}
