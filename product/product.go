package product

// Product Represents a product in the vending machine with its properties.
type Product struct {
	ID       string
	Name     string
	Price    int
	Quantity int
}

func NewProduct(id, name string, price int, quantity int) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}
