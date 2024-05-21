package main

import "fmt"

type Food struct {
	Name   string
	Precio float64
}

func (f *Food) ChangePrice(newPrice float64) {
	f.Precio = newPrice
}

func Methods2() {
	food := Food{"Pizza", 12.99}
	fmt.Println(food.Precio)
	food.ChangePrice(15.99)
	fmt.Println(food.Precio)
}

/*
	Tip: Si un método en una interfaz tiene un puntero receptor, solo los valores que son punteros pueden implementar la interfaz.
	Esto significa que si intentas implementar la interfaz con un valor que no es un puntero, obtendrás un error de compilación.
	Poner todo en punteros es una buena práctica, ya que los métodos con punteros pueden modificar el valor que están llamando.
*/
