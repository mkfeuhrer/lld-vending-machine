package inventory

import "vendingmachine/product"

// Inventory Represents the inventory of products in the vending machine.
type Inventory struct {
	Products []Stock
	// can be further use to track aggregate items like active product count, etc
}

type Stock struct {
	Product *product.Product
	Stock   int
}

// AddProduct Adds a product to the inventory.
func (i *Inventory) AddProduct(p *product.Product, quantity int) {
	i.Products = append(i.Products, Stock{Product: p, Stock: quantity})
}

// RemoveProduct Removes a product from the inventory.
func (i *Inventory) RemoveProduct(p *product.Product) {
	for index, stock := range i.Products {
		if stock.Product.ID == p.ID {
			i.Products = append(i.Products[:index], i.Products[index+1:]...)
			return
		}
	}
}

// IsProductAvailable Checks if a product is available in the inventory.
func (i *Inventory) IsProductAvailable(p *product.Product) bool {
	for _, stock := range i.Products {
		if stock.Product.ID == p.ID && stock.Stock > 0 {
			return true
		}
	}
	return false
}

// DecreaseProductQuantity Decreases the quantity of a product in the inventory.
func (i *Inventory) DecreaseProductQuantity(p *product.Product, quantity int) {
	for index, stock := range i.Products {
		if stock.Product.ID == p.ID {
			i.Products[index].Stock -= quantity
			return
		}
	}
}

// IncreaseProductQuantity Increases the quantity of a product in the inventory.
func (i *Inventory) IncreaseProductQuantity(p *product.Product, quantity int) {
	for index, stock := range i.Products {
		if stock.Product.ID == p.ID {
			i.Products[index].Stock += quantity
			return
		}
	}
}

// NewInventory Creates a new inventory.
func NewInventory(products []Stock) *Inventory {
	return &Inventory{
		Products: products,
	}
}
