package storage

import (
	"MySQL/pkg/product"
	"database/sql"
	"fmt"
)

type scanner interface {
	Scan(dest ...interface{}) error
}

const (
	mysqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products (
		id INT AUTO_INCREMENT NOT NULL PRIMARY KEY,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price NUMERIC(10,2) DEFAULT 0,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP
	)`
	mysqlCreateProduct  = `INSERT INTO products (name, observations, price, created_at) VALUES (?, ?, ?, ?)`
	mysqlGetAllProducts = `SELECT id, name, observations, price, created_at, updated_at FROM products`
	mysqlGetProductByID = mysqlGetAllProducts + ` WHERE id = ?`
	mysqlUpdateProduct  = `UPDATE products SET name = ?, observations = ?, price = ?, updated_at = ? WHERE id = ?`
	mysqlDeleteProduct  = `DELETE FROM products WHERE id = ?`
)

type MySQLProduct struct {
	db *sql.DB
}

func NewMySQLProduct(db *sql.DB) *MySQLProduct {
	return &MySQLProduct{db}
}

func (p *MySQLProduct) Migrate() error {
	stmt, err := p.db.Prepare(mysqlMigrateProduct)
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

func (p *MySQLProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		m.CreatedAt,
	)
	if err != nil {
		return err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return err
	}

	m.ID = uint(id)

	fmt.Println("Product created successfully with ID:", m.ID)
	return nil
}

func (p *MySQLProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(mysqlGetAllProducts)
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

func (p *MySQLProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(mysqlGetProductByID)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()

	return scanProductRow(stmt.QueryRow(id))
}

func (p *MySQLProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(mysqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		m.Name,
		stringToNull(m.Observations),
		m.Price,
		timeToNull(m.UpdatedAt),
		m.ID,
	)

	if err != nil {
		return err
	}

	//res: Resolve
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("doens't exist a product with ID: %d", m.ID)
	}

	fmt.Println("Product updated successfully")

	return nil
}

func (p *MySQLProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(mysqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return fmt.Errorf("doens't exist a product with ID: %d", id)
	}

	fmt.Println("Product deleted successfully")

	return nil
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
