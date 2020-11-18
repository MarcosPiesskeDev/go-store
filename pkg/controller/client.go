package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/http_response"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/repository"
)

type ClientController struct {
	clientRepo *repository.ClientModel
}

func NewClientController(clientRepo *repository.ClientModel) *ClientController {
	return &ClientController{clientRepo: clientRepo}
}

func (cc *ClientController) InitClientMethods(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		cc.getClient(rw, req)
	case "POST":
		cc.createClient(rw, req)
	case "PUT":
		cc.updateClient(rw, req)
	case "DELETE":
		cc.deleteClient(rw, req)
	default:
		http_response.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("Method not Allowed").Error())
	}
}

//Get Client
func (cc *ClientController) getClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	idconv, _ := strconv.Atoi(id)

	//Get client by Id
	if id != "" {
		client, er := cc.clientRepo.GetClientById(idconv)

		if er != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
			return
		}

		http_response.JsonResponse(rw, http.StatusOK, client)
		return
	}

	clients, er := cc.clientRepo.GetAllClient()

	if er != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, clients)
}

//Create Client
func (cc *ClientController) createClient(rw http.ResponseWriter, req *http.Request) {
	var client entity.Client
	err := json.NewDecoder(req.Body).Decode(&client)

	if err != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	er := cc.clientRepo.CreateClient(client)

	if er != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, client)
}

//Update Client
func (cc *ClientController) updateClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	idconv, _ := strconv.Atoi(id)
	var client entity.Client

	if id != "" {
		err := json.NewDecoder(req.Body).Decode(&client) // Geting client from request
		if err != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		idExists, er := cc.clientRepo.ChangeClientById(idconv, client)

		http_response.ErrorsReturnEntity(rw, er, idExists, client)
		return
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}

//Delete a client
func (cc *ClientController) deleteClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	idconv, _ := strconv.Atoi(id)

	if id != "" {

		clientExists, er := cc.clientRepo.DeleteClientById(idconv)

		http_response.ErrorsReturnEntity(rw, er, clientExists, "Client deleted with success!")
		return
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}
