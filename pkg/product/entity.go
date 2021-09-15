package product

type Product struct {
	id      int    `json:"id"`
	idStore int    `json:"id_store"`
	name    string `json:"name"`
	price   string `json:"price"`
}

func NewProduct(id, idStore int, name, price string) *Product {
	return &Product{
		id: id, idStore: idStore, name: name, price: price}
}
