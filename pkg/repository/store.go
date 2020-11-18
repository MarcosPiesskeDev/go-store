package repository

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/database"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
)

type StoreModel struct {
	db *sql.DB
}

func NewStoreModel(db *sql.DB) *StoreModel {
	return &StoreModel{db: db}
}

//Method get all store
func (sm StoreModel) GetAllStore() ([]entity.Store, error) {

	var stores []entity.Store
	var store = entity.Store{}
	var client entity.Client
	var clients []entity.Client
	var product = entity.Product{}
	var products []entity.Product

	rowsClient, errC := sm.db.Query("SELECT * FROM client")
	defer rowsClient.Close()

	if errC != nil {
		return nil, errC
	}

	rowsProduct, errP := sm.db.Query("SELECT * FROM product")
	defer rowsProduct.Close()

	if errP != nil {
		return nil, errP
	}

	//Searching clients
	for rowsClient.Next() {
		erc := rowsClient.Scan(&client.ID, &client.IDStore, &client.NickName, &client.Password, &client.Role, &client.FirstName, &client.LastName, &client.Cash, &client.BirthDate)
		if erc != nil {
			return nil, erc
		}

		clientAtt := entity.Client{
			ID:        client.ID,
			IDStore:   client.IDStore,
			NickName:  client.NickName,
			Password:  client.Password,
			Role:      client.Role,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Cash:      client.Cash,
			BirthDate: client.BirthDate,
		}
		clients = append(clients, clientAtt)
	}

	//Searching products
	for rowsProduct.Next() {
		erp := rowsProduct.Scan(&product.ID, &product.IDStore, &product.Name, &product.Price)

		if erp != nil {
			return nil, erp
		}

		productAtt := entity.Product{
			ID:      product.ID,
			IDStore: product.IDStore,
			Name:    product.Name,
			Price:   product.Price,
		}
		products = append(products, productAtt)
	}

	rowsStore, errS := sm.db.Query("SELECT * FROM store")
	defer rowsStore.Close()

	if errS != nil {
		return stores, errS
	}

	for rowsStore.Next() {
		ers := rowsStore.Scan(&store.ID, &store.Cnpj, &store.Name, &store.CompanyName, &store.City, &store.State)
		if ers != nil {
			return nil, ers
		}
		var clientsByStore []entity.Client
		var productsByStore []entity.Product

		//For loop to verify clients by client.id_store is same the store.id
		for _, v := range clients {
			if v.IDStore != store.ID {
				//clientsByStore = clientsByStore[:0]
				continue
			} else {
				clientsByStore = append(clientsByStore, v)
			}
		}

		for _, v := range products {
			if v.IDStore != store.ID {
				continue
			} else {
				productsByStore = append(productsByStore, v)
			}
		}

		storeAtt := entity.Store{
			ID:          store.ID,
			Cnpj:        store.Cnpj,
			Name:        store.Name,
			CompanyName: store.CompanyName,
			City:        store.City,
			State:       store.State,
			Clients:     clientsByStore,
			Products:    productsByStore,
		}
		stores = append(stores, storeAtt)

	}

	return stores, nil
}

//Method get store
func (sm StoreModel) GetStoreById(id int) (entity.Store, error) {

	var clients []entity.Client
	var client = entity.Client{}
	var products []entity.Product
	var product = entity.Product{}
	var store = entity.Store{}

	rowClient, erC := sm.db.Query("SELECT * FROM client WHERE id_store = ?", id)

	if erC != nil {
		return store, erC
	}

	rowProduct, erP := sm.db.Query("SELECT * FROM product WHERE id_store = ?", id)

	if erP != nil {
		return store, erP
	}

	for rowClient.Next() {
		errc := rowClient.Scan(&client.ID, &client.IDStore, &client.NickName, &client.Password, &client.Role, &client.FirstName, &client.LastName, &client.Cash, &client.BirthDate)
		if errc != nil {
			return store, errc
		}

		clientAtt := entity.Client{
			ID:        client.ID,
			IDStore:   client.IDStore,
			NickName:  client.NickName,
			Password:  client.Password,
			Role:      client.Role,
			FirstName: client.FirstName,
			LastName:  client.LastName,
			Cash:      client.Cash,
			BirthDate: client.BirthDate,
		}
		clients = append(clients, clientAtt)
	}

	for rowProduct.Next() {
		errp := rowProduct.Scan(&product.ID, &product.IDStore, &product.Name, &product.Price)

		if errp != nil {
			return store, errp
		}

		productAtt := entity.Product{
			ID:      product.ID,
			IDStore: product.IDStore,
			Name:    product.Name,
			Price:   product.Price,
		}
		products = append(products, productAtt)
	}

	rowStore, erS := sm.db.Query("SELECT * FROM store WHERE id = ?", id)

	if erS != nil {
		return store, erC
	}

	for rowStore.Next() {
		errs := rowStore.Scan(&store.ID, &store.Cnpj, &store.Name, &store.CompanyName, &store.City, &store.State)

		if errs != nil {
			return store, errs
		}

		store = entity.Store{
			ID:          store.ID,
			Cnpj:        store.Cnpj,
			Name:        store.Name,
			CompanyName: store.CompanyName,
			City:        store.City,
			State:       store.State,
			Clients:     clients,
			Products:    products,
		}

	}
	return store, nil
}

//Method create store
func (sm StoreModel) CreateStore(store entity.Store) error {
	rows, err := sm.db.Exec("INSERT INTO store(cnpj, name, company_name, city, state) VALUES (?,?,?,?,?)", store.Cnpj, store.Name, store.CompanyName, store.City, store.State)

	if err != nil {
		return err
	}

	var parsIdInt = int64(store.ID)

	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return nil
}

//Method update store
func (sm StoreModel) ChangeStoreById(id int, store entity.Store) (bool, error) {
	idExists := database.VerifySExists(id, "store")

	if !idExists {
		return false, nil
	}

	_, er := sm.db.Query("UPDATE store SET cnpj = ?, name = ?, company_name = ?, city = ?, state = ? WHERE id = ?", store.Cnpj, store.Name, store.CompanyName, store.City, store.State, id)

	if er != nil {
		return false, er
	}

	return true, nil
}

//Method delete store
func (sm StoreModel) DeleteStoreById(id int) (bool, error) {
	idExists := database.VerifySExists(id, "store")

	if !idExists {
		return false, nil
	}
	_, err := sm.db.Query("DELETE FROM store WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
