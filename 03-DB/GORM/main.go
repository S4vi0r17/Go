package main

import (
	"GORM/model"
	"GORM/storage"
)

func main() {
	storage.NewPostgresDB()

	storage.Pool().AutoMigrate(&model.Product{}, &model.InvoiceHeader{}, &model.InvoiceItem{})
}
