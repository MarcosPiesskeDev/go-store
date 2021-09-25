package client

import (
	"encoding/json"
	"errors"
	"github.com/marcos-dev88/go-store-back/pkg/http_response"
	"net/http"
	"strings"
	"time"
)

type Handler interface {
	ClientHandler(rw http.ResponseWriter, req *http.Request)
	getClient(rw http.ResponseWriter, req *http.Request)
	createClient(rw http.ResponseWriter, req *http.Request)
	updateClient(rw http.ResponseWriter, req *http.Request)
	deleteClient(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ClientHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.getClient(rw, req)
	case "POST":
		h.createClient(rw, req)
	case "PUT":
		h.updateClient(rw, req)
	case "DELETE":
		h.deleteClient(rw, req)
	default:
		http_response.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("error: method not allowed").Error())
	}
}

//Get Client
func (h *handler) getClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")

	//Get client by Id
	if len(id) > 0 {
		client, er := h.service.getClient(id)

		if er != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
			return
		}

		http_response.JsonResponse(rw, http.StatusOK, client)
		return
	}

	clients, er := h.service.getClientList()

	if er != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, clients)
}

//Create Client
func (h *handler) createClient(rw http.ResponseWriter, req *http.Request) {
	client := *NewClient(0, 0, "", "", "", "", "", 0, time.Time{})

	if err := json.NewDecoder(req.Body).Decode(&client); err != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	createdClient, er := h.service.createClient(client)

	if er != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, createdClient)
}

//Update Client
func (h *handler) updateClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")
	client := *NewClient(0, 0, "", "", "", "", "", 0, time.Time{})

	if len(id) > 0 {
		err := json.NewDecoder(req.Body).Decode(&client) // Geting client from request
		if err != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		updatedClient, er := h.service.updateClient(id, client)

		if er != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
			return
		}

		http_response.JsonResponse(rw, http.StatusOK, updatedClient)
		return
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("error: undefined id").Error())
}

//Delete a client
func (h *handler) deleteClient(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/client/")

	if len(id) > 0 {

		if isClientDeleted := h.service.deleteClient(id); isClientDeleted {
			http_response.JsonResponse(rw, http.StatusOK, "message: client deleted with success!")
			return
		}

		http_response.ErrorsReturnEntity(rw, nil, false, "message: client deleted with success!")
		return
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("error: undefined id").Error())
}