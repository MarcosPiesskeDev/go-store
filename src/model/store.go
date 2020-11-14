package model

import (
	"database/sql"
	"go-store-back/src/entity"
	"go-store-back/src/util"

	"log"
	"reflect"
)

type StoreModel struct {
	Db *sql.DB
}

//Method get all store
func (sm StoreModel) GetAllStore() ([]entity.Store, error) {

	var stores []entity.Store
	var store = entity.Store{}
	var client entity.Client
	var clients []entity.Client
	var product = entity.Product{}
	var products []entity.Product

	rowsClient, errC := sm.Db.Query("SELECT * FROM client")
	defer rowsClient.Close()

	if errC != nil {
		return nil, errC
	}

	rowsProduct, errP := sm.Db.Query("SELECT * FROM product")
	defer rowsProduct.Close()

	if errP != nil {
		return nil, errP
	}

	//Searching clients
	for rowsClient.Next() {
		erc := rowsClient.Scan(&client.Id, &client.Id_store, &client.Nick_name, &client.Password, &client.Role, &client.First_name, &client.Last_name, &client.Cash, &client.Birth_date)
		if erc != nil {
			return nil, erc
		}

		clientAtt := entity.Client{
			Id:         client.Id,
			Id_store:   client.Id_store,
			Nick_name:  client.Nick_name,
			Password:   client.Password,
			Role:       client.Role,
			First_name: client.First_name,
			Last_name:  client.Last_name,
			Cash:       client.Cash,
			Birth_date: client.Birth_date,
		}
		clients = append(clients, clientAtt)
	}

	//Searching products
	for rowsProduct.Next() {
		erp := rowsProduct.Scan(&product.Id, &product.Id_store, &product.Name, &product.Price)

		if erp != nil {
			return nil, erp
		}

		productAtt := entity.Product{
			Id:       product.Id,
			Id_store: product.Id_store,
			Name:     product.Name,
			Price:    product.Price,
		}
		products = append(products, productAtt)
	}

	rowsStore, errS := sm.Db.Query("SELECT * FROM store")
	defer rowsStore.Close()

	if errS != nil {
		return stores, errS
	}

	for rowsStore.Next() {
		ers := rowsStore.Scan(&store.Id, &store.Cnpj, &store.Name, &store.Company_name, &store.City, &store.State)
		if ers != nil {
			return nil, ers
		}
		var clientsByStore []entity.Client
		var productsByStore []entity.Product

		//For loop to verify clients by client.id_store is same the store.id
		for _, v := range clients {
			if v.Id_store != store.Id {
				//clientsByStore = clientsByStore[:0]
				continue
			} else {
				clientsByStore = append(clientsByStore, v)
			}
		}

		for _, v := range products {
			if v.Id_store != store.Id {
				continue
			} else {
				productsByStore = append(productsByStore, v)
			}
		}

		storeAtt := entity.Store{
			Id:           store.Id,
			Cnpj:         store.Cnpj,
			Name:         store.Name,
			Company_name: store.Company_name,
			City:         store.City,
			State:        store.State,
			Clients:      clientsByStore,
			Products:     productsByStore,
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

	rowClient, erC := sm.Db.Query("SELECT * FROM client WHERE id_store = ?", id)

	if erC != nil {
		return store, erC
	}

	rowProduct, erP := sm.Db.Query("SELECT * FROM product WHERE id_store = ?", id)

	if erP != nil {
		return store, erP
	}

	for rowClient.Next() {
		errc := rowClient.Scan(&client.Id, &client.Id_store, &client.Nick_name, &client.Password, &client.Role, &client.First_name, &client.Last_name, &client.Cash, &client.Birth_date)

		if errc != nil {
			return store, errc
		}

		clientAtt := entity.Client{
			Id:         client.Id,
			Id_store:   client.Id_store,
			Nick_name:  client.Nick_name,
			Password:   client.Password,
			Role:       client.Role,
			First_name: client.First_name,
			Last_name:  client.Last_name,
			Cash:       client.Cash,
			Birth_date: client.Birth_date,
		}
		clients = append(clients, clientAtt)
	}

	for rowProduct.Next() {
		errp := rowProduct.Scan(&product.Id, &product.Id_store, &product.Name, &product.Price)

		if errp != nil {
			return store, errp
		}

		productAtt := entity.Product{
			Id:       product.Id,
			Id_store: product.Id_store,
			Name:     product.Name,
			Price:    product.Price,
		}
		products = append(products, productAtt)
	}

	rowStore, erS := sm.Db.Query("SELECT * FROM store WHERE id = ?", id)

	if erS != nil {
		return store, erC
	}

	for rowStore.Next() {
		errs := rowStore.Scan(&store.Id, &store.Cnpj, &store.Name, &store.Company_name, &store.City, &store.State)

		if errs != nil {
			return store, errs
		}

		store = entity.Store{
			Id:           store.Id,
			Cnpj:         store.Cnpj,
			Name:         store.Name,
			Company_name: store.Company_name,
			City:         store.City,
			State:        store.State,
			Clients:      clients,
			Products:     products,
		}

	}
	return store, nil
}

//Method create store
func (sm StoreModel) CreateStore(store entity.Store) error {
	rows, err := sm.Db.Exec("INSERT INTO store(cnpj, name, company_name, city, state) VALUES (?,?,?,?,?)", store.Cnpj, store.Name, store.Company_name, store.City, store.State)

	if err != nil {
		return err
	}

	var parsIdInt = int64(store.Id)

	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return nil
}

//Method update store
func (sm StoreModel) ChangeStoreById(id int, store entity.Store) (bool, error) {
	idExists := util.VerifySExists(id, "store")

	if !idExists {
		return false, nil
	}

	_, er := sm.Db.Query("UPDATE store SET cnpj = ?, name = ?, company_name = ?, city = ?, state = ? WHERE id = ?", store.Cnpj, store.Name, store.Company_name, store.City, store.State, id)

	if er != nil {
		return false, er
	}

	return true, nil
}

//Method delete store
func (sm StoreModel) DeleteStoreById(id int) (bool, error) {
	idExists := util.VerifySExists(id, "store")

	if !idExists {
		return false, nil
	}
	_, err := sm.Db.Query("DELETE FROM store WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
