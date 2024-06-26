package storage

import (
	"MySQL/pkg/invoiceitem"
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE CASCADE ON DELETE CASCADE,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE CASCADE ON DELETE CASCADE
	)`
)

type MySQLInvoiceItem struct {
	db *sql.DB
}

func NewMySQLInvoiceItem(db *sql.DB) *MySQLInvoiceItem {
	return &MySQLInvoiceItem{db}
}

func (p *MySQLInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("InvoiceItem migration executed successfully")
	return nil
}

func (p *MySQLInvoiceItem) CreateTx(tx *sql.Tx, headerId uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(`INSERT INTO invoice_items (invoice_header_id, product_id) VALUES (?, ?)`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, m := range ms {
		result, err := stmt.Exec(headerId, m.ProductID)
		if err != nil {
			return err
		}

		id, _ := result.LastInsertId()
		m.ID = uint(id)
	}

	return nil
}
