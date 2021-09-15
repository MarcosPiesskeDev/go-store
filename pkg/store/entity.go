package store

import (
	"github.com/marcos-dev88/go-store-back/pkg/client"
	"github.com/marcos-dev88/go-store-back/pkg/product"
)

type Store struct {
	id          int               `json:"id"`
	cnpj        string            `json:"cnpj"`
	name        string            `json:"name"`
	companyName string            `json:"company_name"`
	city        string            `json:"city"`
	state       string            `json:"state"`
	clients     []client.Client   `json:"clients"`
	products    []product.Product `json:"products"`
}

func NewStore(id int, cnpj, name, companyName, city, state string, clients []client.Client, products []product.Product) *Store {
	return &Store{
		id:          id,
		cnpj:        cnpj,
		name:        name,
		companyName: companyName,
		city:        city,
		state:       state,
		clients:     clients,
		products:    products,
	}
}
