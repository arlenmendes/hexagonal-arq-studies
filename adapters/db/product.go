package db

import (
	"database/sql"
	"github/arlenmendes/hexagonal-arq-studies/application"

	_ "github.com/mattn/go-sqlite3"
)

type ProducDb struct {
	db *sql.DB
}

func NewProductDb(db *sql.DB) *ProducDb {
	return &ProducDb{db}
}

func (p *ProducDb) Get(id string) (application.ProductInterface, error) {

	var product application.Product

	stmt, err := p.db.Prepare("select id, name, price, status from products where id = ?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *ProducDb) Save(product application.ProductInterface) (application.ProductInterface, error) {

	var rows int

	p.db.QueryRow("select count(*) from products where id = ?", product.GetID()).Scan(&rows)

	if rows == 0 {
		return p.create(product)
	}

	return p.update(product)
}

func (p *ProducDb) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("insert into products (id, name, price, status) values (?, ?, ?, ?)")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetID(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (p *ProducDb) update(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := p.db.Prepare("update products set name = ?, price = ?, status = ? where id = ?")

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetName(), product.GetPrice(), product.GetStatus(), product.GetID())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, nil
}
