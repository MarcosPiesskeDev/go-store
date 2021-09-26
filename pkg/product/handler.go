package product

import (
	"encoding/json"
	"errors"
	"github.com/marcos-dev88/go-store-back/pkg/http_response"
	"net/http"
	"strings"
)

type Handler interface {
	ProductHandler(rw http.ResponseWriter, req *http.Request)
	getProduct(rw http.ResponseWriter, req *http.Request)
	createProduct(rw http.ResponseWriter, req *http.Request)
	updateProduct(rw http.ResponseWriter, req *http.Request)
	deleteProduct(rw http.ResponseWriter, req *http.Request)
}

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{service: service}
}

func (h *handler) ProductHandler(rw http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		h.getProduct(rw, req)
	case "POST":
		h.createProduct(rw, req)
	case "PUT":
		h.updateProduct(rw, req)
	case "DELETE":
		h.deleteProduct(rw, req)
	default:
		http_response.ErrResponse(rw, http.StatusMethodNotAllowed, errors.New("error: method not allowed").Error())
	}
}

//Get all products
func (h *handler) getProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")
	//Get product by id
	if len(id) > 0 {
		product, er := h.service.getProduct(id)

		if er != nil {
			http_response.ErrResponse(rw, http.StatusBadGateway, er.Error())
			return
		}

		http_response.JsonResponse(rw, http.StatusOK, product)
		return
	}

	products, err := h.service.getProductList()

	if err != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, products)
}

//Create Product
func (h *handler) createProduct(rw http.ResponseWriter, req *http.Request) {
	product := NewProduct(0, 0, "", 0.0)

	if err := json.NewDecoder(req.Body).Decode(&product); err != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
		return
	}

	createdProduct, er := h.service.createProduct(*product)

	if er != nil {
		http_response.ErrResponse(rw, http.StatusBadRequest, er.Error())
		return
	}

	http_response.JsonResponse(rw, http.StatusOK, createdProduct)
}

//Update Product
func (h *handler) updateProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")

	product := NewProduct(0, 0, "", 0.0)

	if len(id) > 0 {
		err := json.NewDecoder(req.Body).Decode(&product)

		if err != nil {
			http_response.ErrResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		updatedProduct, er := h.service.updateProduct(id, *product)

		http_response.ErrorsReturnEntity(rw, er, true, updatedProduct)
		return
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("error: undefined id").Error())
}

//Delete Product
func (h *handler) deleteProduct(rw http.ResponseWriter, req *http.Request) {
	id := strings.TrimPrefix(req.URL.Path, "/product/")

	if id != "" {

		productExists, er := h.service.deleteProduct(id)

		http_response.ErrorsReturnEntity(rw, er, productExists, "message: product deleted with success")
	}

	http_response.ErrResponse(rw, http.StatusBadRequest, errors.New("error: undefined id").Error())
}
