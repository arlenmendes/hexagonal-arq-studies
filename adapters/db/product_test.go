package db_test

import (
	"database/sql"
	"github/arlenmendes/hexagonal-arq-studies/adapters/db"
	"github/arlenmendes/hexagonal-arq-studies/application"
	"log"
	"testing"

	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	tableCreateQuery := `CREATE TABLE products (
												"id" string,
												"name" string,
												"price" float,
												"status" string
											);`

	stmt, err := db.Prepare(tableCreateQuery)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	productInsertQuery := `INSERT INTO products (id, name, price, status) VALUES ("p123", "Product Test", 0, "disabled");`

	stmt, err := db.Prepare(productInsertQuery)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDb_Get(t *testing.T) {
	setUp()

	defer Db.Close()

	createProduct(Db)

	productDb := db.NewProductDb(Db)

	product, err := productDb.Get("p123")

	require.Nil(t, err)

	require.Equal(t, "Product Test", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())
}

func TestProductDb_Save(t *testing.T) {
	setUp()
	defer Db.Close()
	productDb := db.NewProductDb(Db)

	product := application.NewProduct()
	product.Name = "Product Test"
	product.Price = 25

	productResult, err := productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())

	product.Status = "enabled"

	productResult, err = productDb.Save(product)
	require.Nil(t, err)
	require.Equal(t, product.Name, productResult.GetName())
	require.Equal(t, product.Price, productResult.GetPrice())
	require.Equal(t, product.Status, productResult.GetStatus())
}
