package storage

import (
	"database/sql"
	"fmt"

	"github.com/jcblastor/rest_go/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observations VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products(
	name, observations, price, created_at
	) VALUES(
		$1, $2, $3, $4
	) RETURNING id`
	psqlGetAllProduct = `SELECT id, name, observations, price, created_at, updated_at FROM products`
)

type PsqlProduct struct {
	db *sql.DB
}

func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db: db}
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

	fmt.Println("Migración de producto ejecutada correctamente")

	return nil
}

func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	err = stmt.QueryRow(m.Name, stringToNull(m.Observations), m.Price, m.CreatedAt).Scan(&m.Id)
	if err != nil {
		return err
	}

	fmt.Println("Se creo el producto correctamente")
	return nil
}

func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := make(product.Models, 0)
	for rows.Next() {
		p := &product.Model{}

		obsNull := sql.NullString{}
		updateNull := sql.NullTime{}

		err := rows.Scan(
			&p.Id,
			&p.Name,
			&obsNull,
			&p.Price,
			&p.CreatedAt,
			&updateNull,
		)
		if err != nil {
			return nil, err
		}

		p.Observations = obsNull.String
		p.UpdatedAt = updateNull.Time

		products = append(products, p)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return products, nil
}
