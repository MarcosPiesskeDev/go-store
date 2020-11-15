package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/MarcosPiesskeDev/go-store-back/pkg/config"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/entity"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/model"
	"github.com/MarcosPiesskeDev/go-store-back/pkg/util"
)

func InitStoreMethods(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getStore(rw, req)
	case "POST":
		createStore(rw, req)
	case "PUT":
		updateStore(rw, req)
	case "DELETE":
		deleteStore(rw, req)
	default:
		util.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("Method not Allowed").Error())
	}
}

//Get Store
func getStore(rw http.ResponseWriter, req *http.Request) {
	db, err := config.GetDb()
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	idconv, _ := strconv.Atoi(id)

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	storeM := model.StoreModel{
		Db: db,
	}

	//Get Store by id
	if id != "" {
		store, er := storeM.GetStoreById(idconv)

		if er != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		util.JsonResponse(rw, http.StatusOK, store)
		return
	}

	stores, er := storeM.GetAllStore()

	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, stores)

}

//Create Store
func createStore(rw http.ResponseWriter, req *http.Request) {
	var store entity.Store
	err := json.NewDecoder(req.Body).Decode(&store)

	db, err := config.GetDb()

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	storeM := model.StoreModel{
		Db: db,
	}

	er := storeM.CreateStore(store)
	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, store)
}

//Update Store
func updateStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	idconv, _ := strconv.Atoi(id)
	var store entity.Store

	if id != "" {
		err := json.NewDecoder(req.Body).Decode(&store)
		db, err := config.GetDb()

		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		storeM := model.StoreModel{
			Db: db,
		}

		idExists, er := storeM.ChangeStoreById(idconv, store)

		util.ErrorsReturnEntity(rw, er, idExists, store)
		return
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}

//Delete Store
func deleteStore(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/store/")
	idconv, _ := strconv.Atoi(id)

	if id != "" {
		db, err := config.GetDb()

		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		storeM := model.StoreModel{
			Db: db,
		}

		storeExists, er := storeM.DeleteStoreById(idconv)

		util.ErrorsReturnEntity(rw, er, storeExists, "Store deleted with success")
		return
	}
	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())

}
