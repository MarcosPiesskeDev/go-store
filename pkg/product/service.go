package product

import (
	"errors"
	"strconv"
)

type Service interface {
	getProduct(productId string) (Product, error)
	getProductList() ([]Product, error)
	createProduct(product Product) (Product, error)
	updateProduct(productId string, product Product) (Product, error)
	deleteProduct(productId string) (bool, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository: repository}
}

func (s *service) getProduct(productId string) (Product, error) {
	idconv, _ := strconv.Atoi(productId)

	newProduct := NewProduct(0, 0, "", "")

	if len(productId) == 0 {
		return *newProduct, errors.New("we got an empty id")
	}

	product, err := s.repository.getProduct(idconv)

	if err != nil {
		return *newProduct, err
	}

	return product, nil
}

func (s *service) getProductList() ([]Product, error) {
	productList, err := s.repository.getProductList()

	if err != nil {
		return nil, err
	}

	return productList, nil
}

func (s *service) createProduct(product Product) (Product, error) {

	createdProduct, err := s.repository.createProduct(product)

	if err != nil {
		return createdProduct, err
	}

	return createdProduct, nil
}

func (s *service) updateProduct(productId string, product Product) (Product, error) {
	idconv, _ := strconv.Atoi(productId)

	newProduct := NewProduct(0, 0, "", "")

	if len(productId) == 0 {
		return *newProduct, errors.New("we got an empty id")
	}

	updatedProduct, err := s.repository.updateProduct(idconv, product)

	if err != nil {
		return *newProduct, err
	}

	return updatedProduct, nil
}

func (s *service) deleteProduct(productId string) (bool, error){
	idconv, _ := strconv.Atoi(productId)

	if len(productId) == 0 {
		return false, errors.New("we got an empty id")
	}

	isDeleted, err := s.repository.deleteProduct(idconv)

	if err != nil {
		return false, err
	}

	return isDeleted, nil
}
