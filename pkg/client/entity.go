package client

import "time"

type Client struct {
	ID        int       `json:"id"`
	IdStore   int       `json:"id_store"`
	NickName  string    `json:"nick_name"`
	Password  string    `json:"password"`
	Role      string    `json:"role"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Cash      float64   `json:"cash"`
	BirthDate time.Time `json:"birth_date"`
}

func NewClient(id, idStore int, nickName, password, role, firstName, lastName string, cash float64, birthDate time.Time) *Client {
	return &Client{
		ID:        id,
		IdStore:   idStore,
		NickName:  nickName,
		Password:  password,
		Role:      role,
		FirstName: firstName,
		LastName:  lastName,
		Cash:      cash,
		BirthDate: birthDate,
	}
}
