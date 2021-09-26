package client

import (
	"database/sql"
	"errors"
	"github.com/marcos-dev88/go-store-back/pkg/database"
	"log"
	"reflect"
	"time"
)

type Repository interface {
	GetAllClient() ([]Client, error)
	GetClientById(clientId int) (Client, error)
	GetClientsByStoreId(storeId int) ([]Client, error)
	CreateClient(client Client) (Client, error)
	UpdateClient(clientId int, client Client) (Client, error)
	DeleteClientById(id int) (bool, error)
}

type repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *repository {
	return &repository{db: db}
}

func (r *repository) GetAllClient() ([]Client, error) {
	rows, err := r.db.Query("SELECT * FROM client")
	var clients []Client
	var client = *NewClient(0, 0, "", "", "", "", "", 0, time.Time{})

	if err != nil {
		return clients, err
	}

	for rows.Next() {
		er := rows.Scan(
			&client.ID,
			&client.IdStore,
			&client.NickName,
			&client.Password,
			&client.Role,
			&client.FirstName,
			&client.LastName,
			&client.Cash,
			&client.BirthDate,
			)

		if er != nil {
			return clients, er
		}

		clientAtt := *NewClient(
			client.ID,
			client.IdStore,
			client.NickName,
			client.Password,
			client.Role,
			client.FirstName,
			client.LastName,
			client.Cash,
			client.BirthDate,
		)

		clients = append(clients, clientAtt)
	}
	return clients, nil
}

//Method get client by id
func (r *repository) GetClientById(id int) (Client, error) {
	rows, err := r.db.Query("SELECT * FROM client WHERE id = ?", id)
	var client = *NewClient(0, 0, "", "", "", "", "", 0, time.Time{})

	if err != nil {
		return client, err
	}

	for rows.Next() {

		er := rows.Scan(
			&client.ID,
			&client.IdStore,
			&client.NickName,
			&client.Password,
			&client.Role,
			&client.FirstName,
			&client.LastName,
			&client.Cash,
			&client.BirthDate,
		)

		if er != nil {
			return client, er
		}

		client = *NewClient(
			client.ID,
			client.IdStore,
			client.NickName,
			client.Password,
			client.Role,
			client.FirstName,
			client.LastName,
			client.Cash,
			client.BirthDate,
		)
	}
	return client, nil
}

func (r *repository) GetClientsByStoreId(storeId int) ([]Client, error) {
	rows, err := r.db.Query("SELECT * FROM client WHERE id_store = ?", storeId)
	var clients []Client
	var client = *NewClient(0, 0, "", "", "", "", "", 0, time.Time{})

	if err != nil {
		return clients, err
	}

	for rows.Next() {
		er := rows.Scan(
			&client.ID,
			&client.IdStore,
			&client.NickName,
			&client.Password,
			&client.Role,
			&client.FirstName,
			&client.LastName,
			&client.Cash,
			&client.BirthDate,
		)

		if er != nil {
			return clients, er
		}

		clientAtt := *NewClient(
			client.ID,
			client.IdStore,
			client.NickName,
			client.Password,
			client.Role,
			client.FirstName,
			client.LastName,
			client.Cash,
			client.BirthDate,
		)

		clients = append(clients, clientAtt)
	}
	return clients, nil
}

//Method create client
func (r *repository) CreateClient(client Client) (Client, error) {
	dbStr := client.BirthDate
	dbDateFormated := dbStr.Format("01-02-2006")
	rows, err := r.db.Exec("INSERT INTO client (id_store, nick_name, password, role, first_name, last_name, cash, birth_date) VALUES (?,?,?,?,?,?,?,?)",
		client.ID,
		client.IdStore,
		client.NickName,
		client.Password,
		client.Role,
		client.FirstName,
		client.LastName,
		client.Cash,
		dbDateFormated,
	)

	if err != nil {
		return Client{}, err
	}

	var parsIdInt = int64(client.ID)
	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return client, nil
}

//Method update client
func (r *repository) UpdateClient(id int, client Client) (Client, error) {
	idExists := database.VerifySExists(id, "client") // IT NEEDS TO BE REFACTORED
	dbStr := client.BirthDate
	dbDateFormated := dbStr.Format("01-02-2006")

	if !idExists {
		return Client{}, errors.New("doesn't a client with this id")
	}

	_, er := r.db.Query("UPDATE client SET id_store = ?, nick_name = ?, password = ?, role = ?, first_name = ?, last_name = ?, cash = ?, birth_date = ? WHERE id = ?", client.ID, client.IdStore, client.NickName, client.Password, client.Role, client.FirstName, client.LastName, client.Cash, dbDateFormated)

	if er != nil {
		return Client{}, er
	}

	return client, nil
}

//Method delete client
func (r repository) DeleteClientById(id int) (bool, error) {
	idExists := database.VerifySExists(id, "client")

	if !idExists {
		return false, nil
	}

	_, err := r.db.Query("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
