package storage

import (
	"PostgreSQL/pkg/invoiceitem"
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items (
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT invoice_items_id_pk PRIMARY KEY (id),
		CONSTRAINT invoice_items_invoice_header_id_fk FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE CASCADE ON DELETE CASCADE,
		CONSTRAINT invoice_items_product_id_fk FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE CASCADE ON DELETE CASCADE
	)`
)

type PsqlInvoiceItem struct {
	db *sql.DB
}

func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

func (p *PsqlInvoiceItem) Migrate() error {
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

func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerId uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(`INSERT INTO invoice_items (invoice_header_id, product_id) VALUES ($1, $2) RETURNING id, created_at`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, m := range ms {
		err = stmt.QueryRow(headerId, m.ProductID).Scan(&m.ID, &m.CreatedAt)
		if err != nil {
			return err
		}
	}

	return nil
}
