package dataservice

import (
	"INVENTORY/model"
	"database/sql"
)

// create
func CreateTableIfNotExists(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS products1 (
		id INT PRIMARY KEY,
		name VARCHAR(100),
		quantity INT
	);`
	_, err := db.Exec(query)
	return err
}

// Insert product
func CreateProduct(db *sql.DB, Product model.Product) error {
	_, err := db.Exec("INSERT INTO products1 (id, name, quantity) VALUES (?, ?, ?)",
		Product.Id, Product.Name, Product.Quantity)
	return err
}

// update
func UpdateProduct(db *sql.DB, Product model.Product) error {
	_, err := db.Exec(`Update Products1 SET name = ?, quantity = ? where id = ?`, Product.Name, Product.Quantity, Product.Id)
	if err != nil {
		return err
	}
	return nil
}
