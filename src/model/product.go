package model

import (
	"database/sql"
	"go-store-back/src/entity"
	"go-store-back/src/util"
	"log"
	"reflect"
)

type ProductModel struct {
	Db *sql.DB
}

//Method get all product
func (pm ProductModel) GetAllProduct() ([]entity.Product, error) {

	rows, err := pm.Db.Query("SELECT * FROM product")
	var product = entity.Product{}
	var products []entity.Product

	if err != nil {
		return products, err
	}

	for rows.Next() {
		er := rows.Scan(&product.Id, &product.Id_store, &product.Name, &product.Price)

		if er != nil {
			return products, er
		}

		productAtt := entity.Product{
			Id:       product.Id,
			Id_store: product.Id_store,
			Name:     product.Name,
			Price:    product.Price,
		}
		products = append(products, productAtt)
	}

	return products, nil
}

//Method get product by id
func (pm ProductModel) GetProductById(id int) (entity.Product, error) {
	var product = entity.Product{}

	rows, err := pm.Db.Query("SELECT * FROM product WHERE id = ?", id)
	defer rows.Close()

	if err != nil {
		return product, err
	}

	for rows.Next() {
		er := rows.Scan(&product.Id, &product.Id_store, &product.Name, &product.Price)

		if er != nil {
			return product, er
		}

		product = entity.Product{
			Id:       product.Id,
			Id_store: product.Id_store,
			Name:     product.Name,
			Price:    product.Price,
		}
	}
	return product, nil
}

//Method create product
func (pm ProductModel) CreateProduct(product entity.Product) error {
	rows, err := pm.Db.Exec("INSERT INTO product (id_store, name, price) VALUES (?,?,?)", product.Id_store, product.Name, product.Price)

	if err != nil {
		return err
	}
	var parsIdInt = int64(product.Id)
	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return nil
}

//Method update product
func (pm ProductModel) ChangeProductById(id int, product entity.Product) (bool, error) {
	idExists := util.VerifySExists(id, "product")

	if !idExists {
		return false, nil
	}

	_, err := pm.Db.Query("UPDATE product SET  id_store = ?, name = ?, price = ? WHERE id = ?", product.Id_store, product.Name, product.Price, id)

	if err != nil {
		return false, err
	}

	return true, nil
}

//Method delete product
func (pm ProductModel) DeleteProductById(id int) (bool, error) {
	idExists := util.VerifySExists(id, "product")

	if !idExists {
		return false, nil
	}

	_, err := pm.Db.Query("DELETE FROM product WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
