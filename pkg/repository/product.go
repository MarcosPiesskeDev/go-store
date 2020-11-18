package repository

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/database"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
)

type ProductModel struct {
	db *sql.DB
}

func NewProductModel(db *sql.DB) *ProductModel {
	return &ProductModel{db: db}
}

//Method get all product
func (pm ProductModel) GetAllProduct() ([]entity.Product, error) {

	rows, err := pm.db.Query("SELECT * FROM product")
	var product = entity.Product{}
	var products []entity.Product

	if err != nil {
		return products, err
	}

	for rows.Next() {
		er := rows.Scan(&product.ID, &product.IDStore, &product.Name, &product.Price)

		if er != nil {
			return products, er
		}

		productAtt := entity.Product{
			ID:      product.ID,
			IDStore: product.IDStore,
			Name:    product.Name,
			Price:   product.Price,
		}
		products = append(products, productAtt)
	}

	return products, nil
}

//Method get product by id
func (pm ProductModel) GetProductById(id int) (entity.Product, error) {
	var product = entity.Product{}

	rows, err := pm.db.Query("SELECT * FROM product WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		return product, err
	}

	for rows.Next() {
		er := rows.Scan(&product.ID, &product.IDStore, &product.Name, &product.Price)

		if er != nil {
			return product, er
		}

		product = entity.Product{
			ID:      product.ID,
			IDStore: product.IDStore,
			Name:    product.Name,
			Price:   product.Price,
		}
	}
	return product, nil
}

//Method create product
func (pm ProductModel) CreateProduct(product entity.Product) error {
	rows, err := pm.db.Exec("INSERT INTO product (id_store, name, price) VALUES (?,?,?)", product.IDStore, product.Name, product.Price)

	if err != nil {
		return err
	}
	var parsIdInt = int64(product.ID)
	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return nil
}

//Method update product
func (pm ProductModel) ChangeProductById(id int, product entity.Product) (bool, error) {
	idExists := database.VerifySExists(id, "product")

	if !idExists {
		return false, nil
	}

	_, err := pm.db.Query("UPDATE product SET  id_store = ?, name = ?, price = ? WHERE id = ?", product.IDStore, product.Name, product.Price, id)

	if err != nil {
		return false, err
	}

	return true, nil
}

//Method delete product
func (pm ProductModel) DeleteProductById(id int) (bool, error) {
	idExists := database.VerifySExists(id, "product")

	if !idExists {
		return false, nil
	}

	_, err := pm.db.Query("DELETE FROM product WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
