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
