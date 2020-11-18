package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/repository"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/util"
)

type StoreController struct {
	storeRepo *repository.StoreModel
}

func NewStoreController(storeRepo *repository.StoreModel) *StoreController {
	return &StoreController{storeRepo: storeRepo}
}

func (sc *StoreController) InitStoreMethods(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		sc.getStore(rw, req)
	case "POST":
		sc.createStore(rw, req)
	case "PUT":
		sc.updateStore(rw, req)
	case "DELETE":
		sc.deleteStore(rw, req)
	default:
		util.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("Method not Allowed").Error())
	}
}

//Get Store
func (sc *StoreController) getStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	idconv, _ := strconv.Atoi(id)

	//Get Store by id
	if id != "" {
		store, er := sc.storeRepo.GetStoreById(idconv)

		if er != nil {
			util.ErrResponse(rw, http.StatusBadRequest, er.Error())
			return
		}

		util.JsonResponse(rw, http.StatusOK, store)
		return
	}

	stores, er := sc.storeRepo.GetAllStore()

	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, stores)

}

//Create Store
func (sc *StoreController) createStore(rw http.ResponseWriter, req *http.Request) {
	var store entity.Store
	err := json.NewDecoder(req.Body).Decode(&store)

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	er := sc.storeRepo.CreateStore(store)
	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, store)
}

//Update Store
func (sc *StoreController) updateStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	idconv, _ := strconv.Atoi(id)
	var store entity.Store

	if id != "" {
		er := json.NewDecoder(req.Body).Decode(&store)

		idExists, er := sc.storeRepo.ChangeStoreById(idconv, store)

		util.ErrorsReturnEntity(rw, er, idExists, store)
		return
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}

//Delete Store
func (sc *StoreController) deleteStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	idconv, _ := strconv.Atoi(id)

	if id != "" {

		storeExists, er := sc.storeRepo.DeleteStoreById(idconv)

		util.ErrorsReturnEntity(rw, er, storeExists, "Store deleted with success")
		return
	}
	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())

}
