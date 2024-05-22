package storage

import (
	"PostgreSQL/pkg/product"
	"database/sql"
	"fmt"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price NUMERIC(10,2) DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products (name, observations, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
)

type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Product migration executed successfully")
	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)

	if err != nil {
		return err
	}

	fmt.Println("Product created successfully")
	return nil
}
