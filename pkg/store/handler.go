package store

import (
	"encoding/json"
	"errors"
	"github.com/marcos-dev88/go-store-back/pkg/http_response"
	"net/http"
	"strings"
)

type Handler interface {
	StoreHandler(rw http.ResponseWriter, req *http.Request)
	getStore(rw http.ResponseWriter, req *http.Request)
	updateStore(rw http.ResponseWriter, req *http.Request)
	deleteStore(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	service Service
}

func (h handler) StoreHandler(rw http.ResponseWriter, req *http.Request){
	switch req.Method {
	case "GET":
		h.getStore(rw, req)
	case "POST":
		h.createStore(rw, req)
	case "PUT":
		h.updateStore(rw, req)
	case "DELETE":
		h.deleteStore(rw, req)
	default:
		http_response.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("error: method not allowed").Error())
	}
}

//Get all Store
func (h *handler) getStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	//Get Store by id
	if len(id) > 0 {
		store, er := h.service.getStoreById(id)

		if er != nil {
			http_response.ErrResponse(rw, http.StatusBadGateway, er.Error())
			return
		}

		http_response.JsonResponse(rw, http.StatusOK, store)
		return
	}

	storeList, err := h.service.getStoreList()

	if err != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, storeList)
}

//Create Store
func (h *handler) createStore(rw http.ResponseWriter, req *http.Request) {
	newStore := NewStore(0, "", "", "", "", "", nil, nil)

	if err := json.NewDecoder(req.Body).Decode(&newStore); err != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	createdStore, er := h.service.createStore(*newStore)

	if er != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, createdStore)
}

//Update Store
func (h *handler) updateStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")

	newStore := NewStore(0, "", "", "", "", "", nil, nil)

	if len(id) > 0 {
		err := json.NewDecoder(req.Body).Decode(&newStore)

		if err != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		updatedStore, er := h.service.updateStore(id, *newStore)

		http_response.ErrorsReturnEntity(rw, er, true, updatedStore)
		return
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("error: undefined id").Error())
}

//Delete Store
func (h *handler) deleteStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")

	if id != "" {

		storeExists, er := h.service.deleteStore(id)

		http_response.ErrorsReturnEntity(rw, er, storeExists, "message: store deleted with success")
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("error: undefined id").Error())
}
