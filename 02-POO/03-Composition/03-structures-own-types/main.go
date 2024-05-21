package main

import (
	"03-Composition/customer"
	"03-Composition/invoice"
	"03-Composition/invoiceItem"
	"fmt"
)

func main() {
	invoice := invoice.New(
		"Peru",
		"Lima",
		customer.New("Eder Benites", "Av. Los Alamos", "987654321"),
		invoiceItem.NewItems(
			invoiceItem.New(1, "Laptop", 1000.0),
			invoiceItem.New(2, "Mouse", 10.0),
			invoiceItem.New(3, "Keyboard", 20.0),
		),
	)
	invoice.SetTotal()
	fmt.Printf("%+v", invoice)
}
