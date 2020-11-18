package repository

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/database"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
)

type ClientModel struct {
	db *sql.DB
}

func NewClientModel(db *sql.DB) *ClientModel {
	return &ClientModel{db: db}
}

//Method get all client
func (cm ClientModel) GetAllClient() ([]entity.Client, error) {
	rows, err := cm.db.Query("SELECT * FROM client")
	var clients []entity.Client
	var client = entity.Client{}

	if err != nil {
		return clients, err
	}

	for rows.Next() {

		er := rows.Scan(&client.ID, &client.IDStore, &client.NickName, &client.Password, &client.Role, &client.FirstName, &client.LastName, &client.Cash, &client.BirthDate)

		if er != nil {
			return clients, er
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
	return clients, nil
}

//Method get client by id
func (cm ClientModel) GetClientById(id int) (entity.Client, error) {
	rows, err := cm.db.Query("SELECT * FROM client WHERE id = ?", id)
	var client = entity.Client{}

	if err != nil {
		return client, err
	}

	for rows.Next() {

		er := rows.Scan(&client.ID, &client.IDStore, &client.NickName, &client.Password, &client.Role, &client.FirstName, &client.LastName, &client.Cash, &client.BirthDate)

		if er != nil {
			return client, er
		}

		client = entity.Client{
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
	}
	return client, nil
}

//Method create client
func (cm ClientModel) CreateClient(client entity.Client) error {
	dbStr := client.BirthDate
	dbDateFormated := dbStr.Format("01-02-2006")
	rows, err := cm.db.Exec("INSERT INTO client (id_store, nick_name, password, role, first_name, last_name, cash, birth_date) VALUES (?,?,?,?,?,?,?,?)", client.IDStore, client.NickName, client.Password, client.Role, client.FirstName, client.LastName, client.Cash, dbDateFormated)
	if err != nil {
		return err
	}

	var parsIdInt = int64(client.ID)
	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return nil
}

//Method update client
func (cm ClientModel) ChangeClientById(id int, client entity.Client) (bool, error) {
	idExists := database.VerifySExists(id, "client")
	dbStr := client.BirthDate
	dbDateFormated := dbStr.Format("01-02-2006")

	if !idExists {
		return false, nil
	}

	_, er := cm.db.Query("UPDATE client SET id_store = ?, nick_name = ?, password = ?, role = ?, first_name = ?, last_name = ?, cash = ?, birth_date = ? WHERE id = ?", client.ID, client.IDStore, client.NickName, client.Password, client.Role, client.FirstName, client.LastName, client.Cash, dbDateFormated)

	if er != nil {
		return false, er
	}

	return true, nil
}

//Method delete client
func (cm ClientModel) DeleteClientById(id int) (bool, error) {
	idExists := database.VerifySExists(id, "client")

	if !idExists {
		return false, nil
	}

	_, err := cm.db.Query("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
