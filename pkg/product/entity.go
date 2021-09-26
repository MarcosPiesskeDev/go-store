package product

type Product struct {
	ID      int    `json:"id"`
	IdStore int    `json:"id_store"`
	Name    string `json:"name"`
	Price   float64 `json:"price"`
}

func NewProduct(id, idStore int, name string, price float64) *Product {
	return &Product{
		ID: id, IdStore: idStore, Name: name, Price: price}
}
