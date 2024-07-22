package main

import "GORM/storage"

func main() {
	storage.NewPostgresDB()
}
