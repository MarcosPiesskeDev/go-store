package store

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/marcos-dev88/go-store-back/pkg/client"
	"github.com/marcos-dev88/go-store-back/pkg/database"
	"github.com/marcos-dev88/go-store-back/pkg/product"
	"log"
	"reflect"
)

type Repository interface {
	getStoreList(clients []client.Client, products []product.Product) ([]Store, error)
	getStoreById(storeId int, clients []client.Client, products []product.Product) (Store, error)
	createStore(store Store) (Store, error)
	updateStore(storeId int, store Store) (Store, error)
	deleteStore(storeId int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) getStoreList(clients []client.Client, products []product.Product) ([]Store, error) {
	var storeList []Store
	newStore := NewStore(0, "", "", "", "", "", clients, products)

	rows, err := r.db.Query(
		"SELECT * FROM store " +
			"INNER JOIN product p ON p.id_store = store.id " +
			"INNER JOIN client c c.id_store",
	)

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		err := rows.Scan(
			&newStore.id,
			&newStore.cnpj,
			&newStore.name,
			&newStore.companyName,
			&newStore.city,
			&newStore.state,
		)

		if err != nil {
			return nil, err
		}

		storeAtt := NewStore(
			newStore.id,
			newStore.cnpj,
			newStore.name,
			newStore.companyName,
			newStore.city,
			newStore.state,
			newStore.clients,
			newStore.products,
		)

		storeList = append(storeList, *storeAtt)
	}

	return storeList, nil
}

func (r *repository) getStoreById(storeId int, clients []client.Client, products []product.Product) (Store, error) {
	newStore := NewStore(0, "", "", "", "", "", clients, products)

	rows, err := r.db.Query(
		"SELECT * FROM store "+
			"INNER JOIN product p ON p.id_store = store.id "+
			"INNER JOIN client c c.id_store = store.id "+
			"WHERE id = ?", storeId,
	)

	if err != nil {
		return *newStore, err
	}

	for rows.Next() {
		err := rows.Scan(
			&newStore.id,
			&newStore.cnpj,
			&newStore.name,
			&newStore.companyName,
			&newStore.city,
			&newStore.state,
		)

		if err != nil {
			return *newStore, err
		}

		newStore = NewStore(
			newStore.id,
			newStore.cnpj,
			newStore.name,
			newStore.companyName,
			newStore.city,
			newStore.state,
			newStore.clients,
			newStore.products,
		)
	}

	return *newStore, nil
}

func (r *repository) createStore(store Store) (Store, error) {
	newStore := NewStore(0, "", "", "", "", "", nil, nil)

	rows, err := r.db.Exec(
		"INSERT INTO store(cnpj, name, company_name, city, state) VALUES (?,?,?,?,?)",
			store.id,
			store.cnpj,
			store.companyName,
			store.city,
			store.state,
		)

	if err != nil {
		return *newStore, nil
	}

	parsIdInt := int64(store.id)
	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return store, nil
}

func (r *repository) updateStore(storeId int, store Store) (Store, error) {
	newStore := NewStore(0, "", "", "", "", "", nil, nil)
	idExists := database.VerifySExists(storeId, "client") // IT NEEDS TO BE REFACTORED

	if !idExists {
		return *newStore, errors.New("doesn't a client with this id")
	}

	_, err := r.db.Query("UPDATE store SET (cnpj = ?, name = ?, company_name = ?, city = ?, state = ? WHERE id = ?",
		store.cnpj,
		store.name,
		store.companyName,
		store.city,
		store.state,
	)

	if err != nil {
		return *newStore, err
	}

	return store, nil
}

func (r *repository) deleteStore(storeId int) (bool, error) {
	idExists := database.VerifySExists(storeId, "client")

	if !idExists {
		return false, errors.New(fmt.Sprintf("there is no store with this id -> %v", storeId))
	}

	_, err := r.db.Query("DELETE FROM store WHERE id = ?", storeId)

	if err != nil {
		return false, err
	}

	return true, nil
}
