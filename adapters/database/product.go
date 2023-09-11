package database

import (
	"database/sql"

	"github.com/intwone/ports-and-adapters-golang/application"
	_ "github.com/mattn/go-sqlite3"
)

type ProductDatabase struct {
	database *sql.DB
}

func (pd *ProductDatabase) Get(id string) (application.ProductInterface, error) {
	var product application.Product

	stmt, err := pd.database.Prepare("SELECT id, name, price, status FROM products where id=?")

	if err != nil {
		return nil, err
	}

	err = stmt.QueryRow(id).Scan(&product.Id, &product.Name, &product.Price, &product.Status)

	if err != nil {
		return nil, err
	}

	return &product, nil
}

func NewProductDatabase(db *sql.DB) *ProductDatabase {
	return &ProductDatabase{database: db}
}

func (pd *ProductDatabase) create(product application.ProductInterface) (application.ProductInterface, error) {
	stmt, err := pd.database.Prepare(`INSERT INTO products (id, name, price, status) values (?, ?, ?, ?)`)

	if err != nil {
		return nil, err
	}

	_, err = stmt.Exec(product.GetId(), product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	err = stmt.Close()

	if err != nil {
		return nil, err
	}

	return product, err
}

func (pd *ProductDatabase) update(product application.ProductInterface) (application.ProductInterface, error) {
	_, err := pd.database.Exec("UPDATE products set name = ?, price = ?, status = ?", product.GetName(), product.GetPrice(), product.GetStatus())

	if err != nil {
		return nil, err
	}

	return product, err
}

func (pd *ProductDatabase) Save(product application.ProductInterface) (application.ProductInterface, error) {
	var rows int

	pd.database.QueryRow("SELECT id FROM products WHERE id = ?", product.GetId()).Scan(&rows)

	if rows == 0 {
		_, err := pd.create(product)

		if err != nil {
			return nil, err
		}
	} else {
		_, err := pd.update(product)

		if err != nil {
			return nil, err
		}
	}

	return product, nil
}
