package controller

import (
	"encoding/json"
	"errors"
	"go-store-back/src/config"
	"go-store-back/src/entity"
	"go-store-back/src/model"
	"go-store-back/src/util"
	"strconv"
	"strings"

	"net/http"
)

func InitClientMethods(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getClient(rw, req)
	case "POST":
		createClient(rw, req)
	case "PUT":
		updateClient(rw, req)
	case "DELETE":
		deleteClient(rw, req)
	default:
		util.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("Method not Allowed").Error())
	}
}

//Get Client
func getClient(rw http.ResponseWriter, req *http.Request) {
	db, err := config.GetDb()
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	idconv, _ := strconv.Atoi(id)

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	clientM := model.ClientModel{
		Db: db,
	}

	//Get client by Id
	if id != "" {
		client, er := clientM.GetClientById(idconv)

		if er != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		util.JsonResponse(rw, http.StatusOK, client)
		return
	}

	clients, er := clientM.GetAllClient()

	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, clients)
}

//Create Client
func createClient(rw http.ResponseWriter, req *http.Request) {
	var client entity.Client
	err := json.NewDecoder(req.Body).Decode(&client)

	db, err := config.GetDb()

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	clientM := model.ClientModel{
		Db: db,
	}

	er := clientM.CreateClient(client)

	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, client)
}

//Update Client
func updateClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	idconv, _ := strconv.Atoi(id)
	var client entity.Client

	if id != "" {
		err := json.NewDecoder(req.Body).Decode(&client) // Geting client from request
		db, err := config.GetDb()
		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		clientM := model.ClientModel{
			Db: db,
		}

		idExists, er := clientM.ChangeClientById(idconv, client)

		util.ErrorsReturnEntity(rw, er, idExists, client)
		return
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}

//Delete a client
func deleteClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	idconv, _ := strconv.Atoi(id)

	if id != "" {
		db, err := config.GetDb()

		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		clientM := model.ClientModel{
			Db: db,
		}

		clientExists, er := clientM.DeleteClientById(idconv)

		util.ErrorsReturnEntity(rw, er, clientExists, "Client deleted with success!")
		return
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}
