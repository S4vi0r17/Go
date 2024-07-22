package main

import (
	"GORM/model"
	"GORM/storage"
	"fmt"
)

func main() {
	storage.NewPostgresDB()

	storage.Pool().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})

	// Read all
	// var products []model.Product
	products := make([]model.Product, 0)
	storage.Pool().Find(&products)
	for _, product := range products {
		// fmt.Printf("%d %s %s %f\n", product.ID, product.Name, *product.Observations, product.Price)
		fmt.Printf("%d %s %f\n", product.ID, product.Name, product.Price)
	}
}

// func ptrString(s string) *string {
// 	return &s
// }
