package main

import (
	// "PostgreSQL/pkg/invoiceheader"
	// "PostgreSQL/pkg/invoiceitem"
	"PostgreSQL/pkg/product"
	"PostgreSQL/storage"
	"log"
)

func main() {
	storage.NewPostgresDB()

	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	// if err := serviceInvoiceHeader.Migrate(); err != nil {
	// 	log.Fatalf("invoiceheader.Migrate: %v", err)
	// }

	// storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
	// serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	// if err := serviceInvoiceItem.Migrate(); err != nil {
	// 	log.Fatalf("invoiceitem.Migrate: %v", err)
	// }

	// // Create product
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// m := &product.Model{
	// 	Name:         "Taco",
	// 	Observations: "Mexican food",
	// 	Price:        10.5,
	// }
	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// fmt.Printf("%+v\n", m)

	// // Get all products
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// products, err := serviceProduct.GetAll()
	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }

	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	// // Get product by ID
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m, err := serviceProduct.GetByID(5)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	log.Fatalf("There is no product with this ID")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	// // Update product
	// storageProduct := storage.NewPsqlProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m := &product.Model{
	// 	ID:           3,
	// 	Name:         "Maki",
	// 	Observations: "Japanese food",
	// 	Price:        15.5,
	// }

	// err := serviceProduct.Update(m)
	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	// Delete product
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	err := serviceProduct.Delete(4)
	if err != nil {
		log.Fatalf("product.Delete: %v", err)
	}
}
