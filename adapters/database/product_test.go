package database_test

import (
	"database/sql"
	"log"
	"testing"

	"github.com/intwone/ports-and-adapters-golang/adapters/database"
	"github.com/intwone/ports-and-adapters-golang/application"
	"github.com/stretchr/testify/require"
)

var Database *sql.DB

func setup() {
	db, _ := sql.Open("sqlite3", ":memory:")
	createTable(db)
	createProduct(db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products ("id" string, "name": string, "price" int, "status": string)`

	stmt, err := db.Prepare(table)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products values ("1", "product1", 1000, "enabled")`

	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()
}

func TestProductDatabase_Get(t *testing.T) {
	setup()
	defer Database.Close()

	productDatabase := database.NewProductDatabase(Database)

	product, err := productDatabase.Get("1")

	require.Nil(t, err)
	require.Equal(t, "product1", product.GetName())
	require.Equal(t, 1000, product.GetPrice())
	require.Equal(t, "enabled", product.GetStatus())
}

func TestProductDatabase_Save(t *testing.T) {
	setup()
	defer Database.Close()

	productDatabase := database.NewProductDatabase(Database)

	product := application.NewProduct()
	product.Name = "product1"
	product.Price = 1000

	result, err := productDatabase.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Name, result.GetName())
	require.Equal(t, product.Price, result.GetPrice())
	require.Equal(t, product.Status, result.GetStatus())

	product.Status = string(application.Disabled)

	result, err = productDatabase.Save(product)

	require.Nil(t, err)
	require.Equal(t, product.Status, result.GetStatus())
}
