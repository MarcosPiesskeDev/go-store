package product

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
)

type Repository interface {
	createProduct(product Product) (Product, error)
	getProduct(productId int) (Product, error)
	getProductList() ([]Product, error)
	updateProduct(productId int, product Product) (Product, error)
	deleteProduct(productId int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r repository) createProduct(product Product) (Product, error) {
	rows, err := r.db.Exec("INSERT INTO product (id_store, name, price) VALUES (?,?,?)", product.idStore, product.name, product.price)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("we got an error: %v", err)
		}
	}(r.db)

	if err != nil {
		return product, err
	}

	var parsIdInt = int64(product.id)

	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		return product, errors.New("we got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return product, nil
}

func (r repository) getProduct(productId int) (Product, error) {

	product := NewProduct(0, 0, "", "")

	rows, err := r.db.Query("SELECT * FROM product WHERE id = ?", productId)

	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Fatalf("we got an error: %v", err)
		}
	}(rows)

	if err != nil {
		return *product, err
	}

	for rows.Next() {
		er := rows.Scan(&product.id, &product.idStore, &product.name, &product.price)

		if er != nil {
			return *product, er
		}

		product = NewProduct(product.id, product.idStore, product.name, product.price)
	}

	return *product, nil
}

func (r repository) getProductList() ([]Product, error) {
	product := NewProduct(0, 0, "", "")

	rows, err := r.db.Query("SELECT * FROM product")

	if err != nil {
		return nil, err
	}

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("we got an error: %v", err)
		}
	}(r.db)

	var productList []Product

	for rows.Next() {
		err := rows.Scan(&product.id, &product.idStore, &product.name, &product.price)

		if err != nil {
			return nil, err
		}

		product = NewProduct(product.id, product.idStore, product.name, product.price)

		productList = append(productList, *product)
	}

	return productList, nil
}

func (r repository) updateProduct(productId int, product Product) (Product, error) {

	_, err := r.db.Query("UPDATE product SET  id_store = ?, name = ?, price = ? WHERE id = ?", product.id, product.name, product.price, productId)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r repository) deleteProduct(productId int) (bool, error) {

	_, err := r.db.Query("DELETE FROM product WHERE id = ?", productId)

	if err != nil {
		return false, err
	}

	return true, nil

}
