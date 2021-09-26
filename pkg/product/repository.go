package product

import (
	"database/sql"
	"errors"
	"log"
	"reflect"
)

type Repository interface {
	GetProduct(productId int) (Product, error)
	GetProductList() ([]Product, error)
	GetProductListByStoreId(storeId int) ([]Product, error)
	CreateProduct(product Product) (Product, error)
	UpdateProduct(productId int, product Product) (Product, error)
	DeleteProduct(productId int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r repository) CreateProduct(product Product) (Product, error) {
	rows, err := r.db.Exec("INSERT INTO product (id_store, name, price) VALUES (?,?,?)", product.IdStore, product.Name, product.Price)

	defer func(db *sql.DB) {
		err := db.Close()
		if err != nil {
			log.Fatalf("we got an error: %v", err)
		}
	}(r.db)

	if err != nil {
		return product, err
	}

	var parsIdInt = int64(product.ID)

	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		return product, errors.New("we got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return product, nil
}

func (r repository) GetProduct(productId int) (Product, error) {

	product := NewProduct(0, 0, "", 0.0)

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
		er := rows.Scan(&product.ID, &product.IdStore, &product.Name, &product.Price)

		if er != nil {
			return *product, er
		}

		product = NewProduct(product.ID, product.IdStore, product.Name, product.Price)
	}

	return *product, nil
}

func (r repository) GetProductListByStoreId(storeId int) ([]Product, error){
	rows, err := r.db.Query("SELECT * FROM product WHERE id_store = ?", storeId)

	if err != nil {
		return nil, err
	}
	var product = *NewProduct(0, 0, "", 0.0)
	var productList []Product

	for rows.Next() {
		err := rows.Scan(&product.ID, &product.IdStore, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		productAtt := *NewProduct(product.ID, product.IdStore, product.Name, product.Price)
		productList = append(productList, productAtt)
	}

	return productList, nil
}

func (r repository) GetProductList() ([]Product, error) {
	product := NewProduct(0, 0, "", 0.0)

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
		err := rows.Scan(&product.ID, &product.IdStore, &product.Name, &product.Price)

		if err != nil {
			return nil, err
		}

		product = NewProduct(product.ID, product.IdStore, product.Name, product.Price)

		productList = append(productList, *product)
	}

	return productList, nil
}

func (r repository) UpdateProduct(productId int, product Product) (Product, error) {

	_, err := r.db.Query("UPDATE product SET  id_store = ?, name = ?, price = ? WHERE id = ?", product.ID, product.Name, product.Price, productId)

	if err != nil {
		return product, err
	}

	return product, nil
}

func (r repository) DeleteProduct(productId int) (bool, error) {

	_, err := r.db.Query("DELETE FROM product WHERE id = ?", productId)

	if err != nil {
		return false, err
	}

	return true, nil

}
