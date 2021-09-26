package store

import (
	"github.com/marcos-dev88/go-store-back/pkg/client"
	"github.com/marcos-dev88/go-store-back/pkg/product"
)

type Store struct {
	ID          int               `json:"id"`
	Cnpj        string            `json:"cnpj"`
	Name        string            `json:"name"`
	CompanyName string            `json:"company_name"`
	City        string            `json:"city"`
	State       string            `json:"state"`
	Clients     []client.Client   `json:"clients"`
	Products    []product.Product `json:"products"`
}

func NewStore(id int, cnpj, name, companyName, city, state string, clients []client.Client, products []product.Product) *Store {
	return &Store{
		ID:          id,
		Cnpj:        cnpj,
		Name:        name,
		CompanyName: companyName,
		City:        city,
		State:       state,
		Clients:     clients,
		Products:    products,
	}
}
