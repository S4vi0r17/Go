package main

import (
	"GORM/model"
	"GORM/storage"
)

func main() {
	storage.NewPostgresDB()

	storage.Pool().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

	// Create
	observations := "Pepperoni"
	product1 := model.Product{
		Name:         "Pizza",
		Observations: &observations,
		Price:        12.5,
	}

	product2 := model.Product{
		Name:  "Coke",
		Price: 1.5,
	}

	product3 := model.Product{
		Name:         "Fries",
		Observations: ptrString("Large"),
		Price:        3.5,
	}

	storage.Pool().Create(&product1)
	storage.Pool().Create(&product2)
	storage.Pool().Create(&product3)
}

func ptrString(s string) *string {
	return &s
}
