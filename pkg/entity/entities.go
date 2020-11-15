package entity

import "time"

type Store struct {
	Id           int       `json:"id"`
	Cnpj         string    `json:"cnpj"`
	Name         string    `json:"name"`
	Company_name string    `json:"company_name"`
	City         string    `json:"city"`
	State        string    `json:"state"`
	Clients      []Client  `json:"clients"`
	Products     []Product `json:"products"`
}

type Client struct {
	Id         int       `json:"id"`
	Id_store   int       `json:"id_store"`
	Nick_name  string    `json:"nick_name"`
	Password   string    `json:"password"`
	Role       string    `json:"role"`
	First_name string    `json:"first_name"`
	Last_name  string    `json:"last_name"`
	Cash       float64   `json:"cash"`
	Birth_date time.Time `json:"birth_date"`
}

type Product struct {
	Id       int    `json:"id"`
	Id_store int    `json:"id_store"`
	Name     string `json:"name"`
	Price    string `json:"price"`
}
