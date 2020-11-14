package controller

import (
	"encoding/json"
	"errors"
	"go-store-back/src/config"
	"go-store-back/src/entity"
	"go-store-back/src/model"
	"go-store-back/src/util"
	"net/http"
	"strconv"
	"strings"
)

func InitProductMethods(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		getProduct(rw, req)
	case "POST":
		createProduct(rw, req)
	case "PUT":
		updateProduct(rw, req)
	case "DELETE":
		deleteProduct(rw, req)
	default:
		util.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("Method not Allowed").Error())
	}
}

//Get all products
func getProduct(rw http.ResponseWriter, req *http.Request) {
	db, err := config.GetDb()
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	idconv, _ := strconv.Atoi(id)

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	productM := model.ProductModel{
		Db: db,
	}

	//Get product by id
	if id != "" {
		product, er := productM.GetProductById(idconv)

		if er != nil {
			util.ErrResponse(rw, http.StatusBadGateway, err.Error())
			return
		}

		util.JsonResponse(rw, http.StatusOK, product)
		return
	}

	products, er := productM.GetAllProduct()

	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, products)
}

//Create Product
func createProduct(rw http.ResponseWriter, req *http.Request) {
	var product entity.Product

	err := json.NewDecoder(req.Body).Decode(&product)

	db, err := config.GetDb()

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	productM := model.ProductModel{
		Db: db,
	}

	er := productM.CreateProduct(product)
	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, product)
}

//Update Product
func updateProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	idconv, _ := strconv.Atoi(id)
	var product entity.Product

	if id != "" {
		err := json.NewDecoder(req.Body).Decode(&product)
		db, err := config.GetDb()

		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		productM := model.ProductModel{
			Db: db,
		}

		idExists, er := productM.ChangeProductById(idconv, product)

		util.ErrorsReturnEntity(rw, er, idExists, product)
		return
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}

//Delete Product
func deleteProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	idconv, _ := strconv.Atoi(id)

	db, err := config.GetDb()
	if id != "" {
		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		productM := model.ProductModel{
			Db: db,
		}

		productExists, er := productM.DeleteProductById(idconv)

		util.ErrorsReturnEntity(rw, er, productExists, "Product deleted with success")
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}
