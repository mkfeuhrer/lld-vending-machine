package product

// Product Represents a product in the vending machine with its properties.
type Product struct {
	ID       string
	Name     string
	Price    float64
	Quantity int
}

func NewProduct(id, name string, price float64, quantity int) *Product {
	return &Product{
		ID:       id,
		Name:     name,
		Price:    price,
		Quantity: quantity,
	}
}
