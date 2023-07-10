package main

import "fmt"

func main() {

	//var numeros[5] int
	//numeros := [5]int{100, 200, 300, 400, 500}
	//numeros := [...]int{100, 200, 300, 400, 500}

	/*
	numeros[0] = 100
	numeros[1] = 200
	numeros[2] = 300
	numeros[3] = 400
	numeros[4] = 500
	*/

	//monedas := [...]string{"PEN", "USD", "EUR", "JPY", "CNY"}
	monedas := [...]string{ 0: "PEN", 1: "USD", 2: "EUR", 3: "JPY", 4: "CNY"}

	for i := 0; i < len(monedas); i++ {
		fmt.Println(monedas[i])
	}

	fmt.Println(monedas)

	submonedas := monedas[0:3] // Slice

	fmt.Println(submonedas)
}