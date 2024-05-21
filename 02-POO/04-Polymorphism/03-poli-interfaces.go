package main

import (
	"fmt"
)

type PayMethod interface {
	Pay()
}

type Paypal struct{}

func (p Paypal) Pay() {
	fmt.Println("Paypal payment")
}

type CreditCard struct{}

func (c CreditCard) Pay() {
	fmt.Println("Credit card payment")
}

type Cash struct{}

func (c Cash) Pay() {
	fmt.Println("Cash payment")
}

func Factory(method string) PayMethod {
	switch method {
	case "Paypal":
		return Paypal{}
	case "Predit":
		return CreditCard{}
	case "Cash":
		return Cash{}
	default:
		return nil
	}
}

func main() {
	// paypal := Factory("paypal")
	// credit := Factory("credit")
	// cash := Factory("cash")

	// paypal.Pay()
	// credit.Pay()
	// cash.Pay()

	var method string
	fmt.Println("Enter payment method:")
	fmt.Println("\"Paypal\", \"Credit\", \"Cash\"")
	fmt.Scanln(&method)

	payment := Factory(method)
	if payment != nil {
		payment.Pay()
	} else {
		fmt.Println("Invalid payment method")
	}
}
