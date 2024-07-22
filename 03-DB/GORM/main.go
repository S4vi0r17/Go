package main

import (
	"GORM/model"
	"GORM/storage"
	"fmt"
)

func main() {
	storage.NewPostgresDB()

	storage.Pool().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

	// Create
	// observations := "Pepperoni"
	// product1 := model.Product{
	// 	Name:         "Pizza",
	// 	Observations: &observations,
	// 	Price:        12.5,
	// }

	// product2 := model.Product{
	// 	Name:  "Coke",
	// 	Price: 1.5,
	// }

	// product3 := model.Product{
	// 	Name:         "Fries",
	// 	Observations: ptrString("Large"),
	// 	Price:        3.5,
	// }

	// storage.Pool().Create(&product1)
	// storage.Pool().Create(&product2)
	// storage.Pool().Create(&product3)

	// Read all
	// var products []model.Product
	// products := make([]model.Product, 0)
	// storage.Pool().Find(&products)
	// for _, product := range products {
	// 	// fmt.Printf("%d %s %s %f\n", product.ID, product.Name, *product.Observations, product.Price)
	// 	fmt.Printf("%d %s %f\n", product.ID, product.Name, product.Price)
	// }

	// Read one
	var product model.Product
	storage.Pool().First(&product, 2)
	fmt.Printf("%d %s %f\n", product.ID, product.Name, product.Price)
}

// func ptrString(s string) *string {
// 	return &s
// }
