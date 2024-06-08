package storage

import (
	"PostgreSQL/pkg/product"
	"database/sql"
	"fmt"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

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
	psqlCreateProduct  = `INSERT INTO products (name, observations, price, created_at) VALUES ($1, $2, $3, $4) RETURNING id`
	psqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	psqlGetProductByID = psqlGetAllProducts + ` WHERE id = $1`
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

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProducts)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var products product.Models
	for rows.Next() {
		m, err := scanProductRow(rows)
		if err != nil {
			return nil, err
		}

		products = append(products, m)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}


func scanProductRow(s scanner) (*product.Model, error) {
	m := &product.Model{}

	observationsNull := sql.NullString{}
	updatedAtNull := sql.NullTime{}

	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationsNull,
		&m.Price,
		&m.CreatedAt,
		&updatedAtNull,
	)

	if err != nil {
		return &product.Model{}, err
	}

	m.Observations = observationsNull.String
	m.UpdatedAt = updatedAtNull.Time

	return m, nil
}
