package main

import (
	"MySQL/pkg/invoice"
	"MySQL/pkg/invoiceheader"
	"MySQL/pkg/invoiceitem"
	"MySQL/storage"
	"log"
)

func main() {
	storage.NewMySQLDB()

	/* Migrate */

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// if err := serviceProduct.Migrate(); err != nil {
	// 	log.Fatalf("product.Migrate: %v", err)
	// }

	// storageInvoiceHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	// serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
	// if err := serviceInvoiceHeader.Migrate(); err != nil {
	// 	log.Fatalf("invoiceheader.Migrate: %v", err)
	// }

	// storageInvoiceItem := storage.NewMySQLInvoiceItem(storage.Pool())
	// serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
	// if err := serviceInvoiceItem.Migrate(); err != nil {
	// 	log.Fatalf("invoiceitem.Migrate: %v", err)
	// }

	/* Migrate */

	/* Create product */

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// m := &product.Model{
	// 	Name:         "Pizza",
	// 	Observations: "Italian food",
	// 	Price:        12.5,
	// }
	// if err := serviceProduct.Create(m); err != nil {
	// 	log.Fatalf("product.Create: %v", err)
	// }

	// fmt.Printf("%+v\n", m)

	/* Create product */

	/* Get all products */

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)
	// products, err := serviceProduct.GetAll()
	// if err != nil {
	// 	log.Fatalf("product.GetAll: %v", err)
	// }

	// for _, p := range products {
	// 	fmt.Println(p)
	// }

	/* Get all products */

	/* Get product by ID */

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m, err := serviceProduct.GetByID(2)
	// switch {
	// case errors.Is(err, sql.ErrNoRows):
	// 	log.Fatalf("There is no product with this ID")
	// case err != nil:
	// 	log.Fatalf("product.GetByID: %v", err)
	// default:
	// 	fmt.Println(m)
	// }

	/* Get product by ID */

	/* Update product */

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// m := &product.Model{
	// 	ID:           1,
	// 	Name:         "Maki",
	// 	Observations: "Japanese food",
	// 	Price:        15.5,
	// }

	// err := serviceProduct.Update(m)
	// if err != nil {
	// 	log.Fatalf("product.Update: %v", err)
	// }

	/* Update product */

	/* Delete product */

	// storageProduct := storage.NewMySQLProduct(storage.Pool())
	// serviceProduct := product.NewService(storageProduct)

	// err := serviceProduct.Delete(1)
	// if err != nil {
	// 	log.Fatalf("product.Delete: %v", err)
	// }

	/* Delete product */

	/* Create invoice */

	storageHeader := storage.NewMySQLInvoiceHeader(storage.Pool())
	storageItem := storage.NewMySQLInvoiceItem(storage.Pool())
	storageInvoice := storage.NewMySQLInvoice(storage.Pool(), storageHeader, storageItem)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Gustavo Benites",
		},
		Items: invoiceitem.Models{
			{ProductID: 2},
			{ProductID: 3},
			// &invoiceitem.Model{ProductID: 1},
			// &invoiceitem.Model{ProductID: 2},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("invoice.Create: %v", err)
	}

	/* Create invoice */
}
