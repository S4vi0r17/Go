package main

import (
	"GORM/model"
	"GORM/storage"
	// "fmt"
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
	// var product model.Product
	// storage.Pool().First(&product, 2)
	// fmt.Printf("%d %s %f\n", product.ID, product.Name, product.Price)

	// Update all fields
	// var product model.Product
	// storage.Pool().First(&product, 3)
	// product.Name = "Fries"
	// product.Price = 4.5
	// storage.Pool().Save(&product)

	// Update only one field
	// var product model.Product
	// storage.Pool().Model(&product).Where("id = ?", 3).Update("price", 5.5)
	// myProduct := model.Product{}
	// myProduct.ID = 3
	// storage.Pool().Model(&myProduct).Updates(model.Product{Price: 5.5, Name: "Fries"})

	// Delete: soft delete
	// var product model.Product
	// // myProduct := model.Product{}
	// // myProduct.ID = 3
	// // storage.Pool().Delete(&myProduct)
	// storage.Pool().Delete(&product, 3)

	// Delete: hard delete
	// var product model.Product
	// storage.Pool().Unscoped().Delete(&product, 3)
	// // myProduct := model.Product{}
	// // myProduct.ID = 3
	// // storage.Pool().Unscoped().Delete(&myProduct)

	// Transaction
	/*
	// Old way
		storage.Pool().Model(&model.InvoiceItem{}).AddForeignKey("product_id", "products(id)", "RESTRICT", "RESTRICT")
		storage.Pool().Model(&model.InvoiceItem{}).AddForeignKey("invoice_header_id", "invoice_headers(id)", "RESTRICT", "RESTRICT")
	*/

	// New way
	invoice := model.InvoiceHeader{
		Client: "John Doe",
		InvoiceItems: []model.InvoiceItem{
			{ProductID: 1},
			{ProductID: 2},
		},
	}

	storage.Pool().Create(&invoice)
}

// func ptrString(s string) *string {
// 	return &s
// }
