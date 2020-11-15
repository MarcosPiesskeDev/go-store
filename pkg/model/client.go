package model

import (
	"database/sql"
	"log"
	"reflect"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/util"
)

type ClientModel struct {
	Db *sql.DB
}

//Method get all client
func (cm ClientModel) GetAllClient() ([]entity.Client, error) {
	rows, err := cm.Db.Query("SELECT * FROM client")
	var clients []entity.Client
	var client = entity.Client{}

	if err != nil {
		return clients, err
	}

	for rows.Next() {

		er := rows.Scan(&client.Id, &client.Id_store, &client.Nick_name, &client.Password, &client.Role, &client.First_name, &client.Last_name, &client.Cash, &client.Birth_date)

		if er != nil {
			return clients, er
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
	return clients, nil
}

//Method get client by id
func (cm ClientModel) GetClientById(id int) (entity.Client, error) {
	rows, err := cm.Db.Query("SELECT * FROM client WHERE id = ?", id)
	var client = entity.Client{}

	if err != nil {
		return client, err
	}

	for rows.Next() {

		er := rows.Scan(&client.Id, &client.Id_store, &client.Nick_name, &client.Password, &client.Role, &client.First_name, &client.Last_name, &client.Cash, &client.Birth_date)

		if er != nil {
			return client, er
		}

		client = entity.Client{
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
	}
	return client, nil
}

//Method create client
func (cm ClientModel) CreateClient(client entity.Client) error {
	dbStr := client.Birth_date
	dbDateFormated := dbStr.Format("01-02-2006")
	rows, err := cm.Db.Exec("INSERT INTO client (id_store, nick_name, password, role, first_name, last_name, cash, birth_date) VALUES (?,?,?,?,?,?,?,?)", client.Id_store, client.Nick_name, client.Password, client.Role, client.First_name, client.Last_name, client.Cash, dbDateFormated)
	if err != nil {
		return err
	}

	var parsIdInt = int64(client.Id)
	if reflect.TypeOf(parsIdInt).Kind() == reflect.String {
		log.Println("We got an id with a type string and we need with type int")
	}

	parsIdInt, _ = rows.LastInsertId()

	return nil
}

//Method update client
func (cm ClientModel) ChangeClientById(id int, client entity.Client) (bool, error) {
	idExists := util.VerifySExists(id, "client")
	dbStr := client.Birth_date
	dbDateFormated := dbStr.Format("01-02-2006")

	if !idExists {
		return false, nil
	}

	_, er := cm.Db.Query("UPDATE client SET id_store = ?, nick_name = ?, password = ?, role = ?, first_name = ?, last_name = ?, cash = ?, birth_date = ? WHERE id = ?", client.Id_store, client.Nick_name, client.Password, client.Role, client.First_name, client.Last_name, client.Cash, dbDateFormated, id)

	if er != nil {
		return false, er
	}

	return true, nil
}

//Method delete client
func (cm ClientModel) DeleteClientById(id int) (bool, error) {
	idExists := util.VerifySExists(id, "client")

	if !idExists {
		return false, nil
	}

	_, err := cm.Db.Query("DELETE FROM client WHERE id = ?", id)

	if err != nil {
		return false, err
	}

	return true, nil
}
