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

type ProductController struct {
	productRepo *repository.ProductModel
}

func NewProductController(productRepo *repository.ProductModel) *ProductController {
	return &ProductController{productRepo: productRepo}
}

func (pc *ProductController) InitProductMethods(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		pc.getProduct(rw, req)
	case "POST":
		pc.createProduct(rw, req)
	case "PUT":
		pc.updateProduct(rw, req)
	case "DELETE":
		pc.deleteProduct(rw, req)
	default:
		util.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("Method not Allowed").Error())
	}
}

//Get all products
func (pc *ProductController) getProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	idconv, _ := strconv.Atoi(id)

	//Get product by id
	if id != "" {
		product, er := pc.productRepo.GetProductById(idconv)

		if er != nil {
			util.ErrResponse(rw, http.StatusBadGateway, er.Error())
			return
		}

		util.JsonResponse(rw, http.StatusOK, product)
		return
	}

	products, err := pc.productRepo.GetAllProduct()

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, products)
}

//Create Product
func (pc *ProductController) createProduct(rw http.ResponseWriter, req *http.Request) {
	var product entity.Product

	err := json.NewDecoder(req.Body).Decode(&product)

	if err != nil {
		util.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	er := pc.productRepo.CreateProduct(product)

	if er != nil {
		util.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	util.JsonResponse(rw, http.StatusOK, product)
}

//Update Product
func (pc *ProductController) updateProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	idconv, _ := strconv.Atoi(id)
	var product entity.Product

	if id != "" {
		err := json.NewDecoder(req.Body).Decode(&product)

		if err != nil {
			util.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		idExists, er := pc.productRepo.ChangeProductById(idconv, product)

		util.ErrorsReturnEntity(rw, er, idExists, product)
		return
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}

//Delete Product
func (pc *ProductController) deleteProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	idconv, _ := strconv.Atoi(id)

	if id != "" {

		productExists, er := pc.productRepo.DeleteProductById(idconv)

		util.ErrorsReturnEntity(rw, er, productExists, "Product deleted with success")
	}

	util.ErrResponse(rw, http.StatusBadRequest, errors.New("Undefined id").Error())
}
