package product

type Product struct {
	id      int    `json:"id"`
	idStore int    `json:"id_store"`
	name    string `json:"name"`
	price   float64 `json:"price"`
}

func NewProduct(id, idStore int, name string, price float64) *Product {
	return &Product{
		id: id, idStore: idStore, name: name, price: price}
}
