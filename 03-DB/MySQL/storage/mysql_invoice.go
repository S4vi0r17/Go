package storage

import (
	"MySQL/pkg/invoice"
	"MySQL/pkg/invoiceheader"
	"MySQL/pkg/invoiceitem"
	"database/sql"
	"fmt"
)

type MySQLInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItem   invoiceitem.Storage
}

func NewMySQLInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *MySQLInvoice {
	return &MySQLInvoice{db, h, i}
}

func (p *MySQLInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("storageHeader.CreateTx: %v", err)
	}

	fmt.Println("Invoice header created with ID:", m.Header.ID)

	if err := p.storageItem.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return fmt.Errorf("storageItem.CreateTx: %v", err)
	}

	fmt.Println("Invoice items created successfully", len(m.Items))

	return tx.Commit()
}
