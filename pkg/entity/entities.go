package entity

import (
	"time"
)

type Store struct {
	ID          int       `json:"id"`
	Cnpj        string    `json:"cnpj"`
	Name        string    `json:"name"`
	CompanyName string    `json:"company_name"`
	City        string    `json:"city"`
	State       string    `json:"state"`
	Clients     []Client  `json:"clients"`
	Products    []Product `json:"products"`
}

type Client struct {
	ID        int       `json:"id"`
	IDStore   int       `json:"id_store"`
	NickName  string    `json:"nick_name"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Cash      float64   `json:"cash"`
	BirthDate time.Time `json:"birth_date"`
}

type Product struct {
	ID      int    `json:"id"`
	IDStore int    `json:"id_store"`
	Name    string `json:"name"`
	Price   string `json:"price"`
}
