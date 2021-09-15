package client

import "time"

type Client struct {
	id        int       `json:"id"`
	idStore   int       `json:"id_store"`
	nickName  string    `json:"nick_name"`
	password  string    `json:"password"`
	role      string    `json:"role"`
	firstName string    `json:"first_name"`
	lastName  string    `json:"last_name"`
	cash      float64   `json:"cash"`
	birthDate time.Time `json:"birth_date"`
}

func NewClient(id, idStore int, nickName, password, role, firstName, lastName string, cash float64, birthDate time.Time) *Client {
	return &Client{
		id:        id,
		idStore:   idStore,
		nickName:  nickName,
		password:  password,
		role:      role,
		firstName: firstName,
		lastName:  lastName,
		cash:      cash,
		birthDate: birthDate,
	}
}
