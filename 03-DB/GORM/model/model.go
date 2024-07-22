package model

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name         string  `gorm:"size:255;not null"`
	Observations string  `gorm:"type:text"`
	Price        float64 `gorm:"not null"`
	InvoiceItems []InvoiceItem
}

type InvoiceHeader struct {
	gorm.Model
	Client       string `gorm:"size:255;not null"`
	InvoiceItems []InvoiceItem
}

type InvoiceItem struct {
	gorm.Model
	InvoiceHeaderID uint
	ProductID       uint
}
