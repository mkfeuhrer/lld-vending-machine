package app

import (
	"errors"
	"vendingmachine/balance"
	"vendingmachine/inventory"
	"vendingmachine/order"
	"vendingmachine/payment/enums"
	"vendingmachine/product"
)

// App Represents the vending machine application.
type App struct {
	Inventory *inventory.Inventory
	Orders    []*order.Order
	Balance   *balance.Balance
}

// NewApp Creates a new vending machine application.
func NewApp(inventory *inventory.Inventory, balance *balance.Balance) *App {
	return &App{
		Inventory: inventory,
		Orders:    make([]*order.Order, 0),
		Balance:   balance,
	}
}

// AddOrder Adds an order to the vending machine application.
func (a *App) AddOrder(order *order.Order) error {
	productNotAvailable := false
	// check inventory
	for _, p := range order.Products {
		if !a.Inventory.IsProductAvailable(&p) {
			productNotAvailable = true
		}
	}

	if productNotAvailable {
		return errors.New("some products are missing in inventory")
	}

	// todo: validate payment amount

	for _, p := range order.Products {
		for _, txn := range order.Payment.Transactions {
			a.Balance.AddPaymentInstrument(txn.PaymentInstrument, p.Quantity*p.Price)
		}
		a.Inventory.DecreaseProductQuantity(&p)
	}

	// todo: handle change return

	a.Orders = append(a.Orders, order)
	return nil
}

// RemoveOrder Removes an order from the vending machine application.
func (a *App) RemoveOrder(order *order.Order) {
	// todo: implement me
}

// AddBalance Adds a payment instrument quantity to the balance of the vending machine application.
func (a *App) AddBalance(paymentInstrument enums.PaymentInstrument, quantity int) {
	a.Balance.AddPaymentInstrument(paymentInstrument, quantity)
}

// DeductBalance Updates a payment instrument quantity in the balance of the vending machine application.
func (a *App) DeductBalance(paymentInstrument enums.PaymentInstrument, quantity int) {
	a.Balance.RemovePaymentInstrument(paymentInstrument, quantity)
}

// AddProduct Adds a product to the inventory.
func (a *App) AddProduct(product *product.Product, quantity int) {
	a.Inventory.AddProduct(product, quantity)
}

// RemoveProduct Removes a product from the inventory.
func (a *App) RemoveProduct(product *product.Product) {
	a.Inventory.RemoveProduct(product)
}

// DecreaseProductQuantity Decreases the quantity of a product in the inventory.
func (a *App) DecreaseProductQuantity(product *product.Product) {
	a.Inventory.DecreaseProductQuantity(product)
}

// IncreaseProductQuantity Increases the quantity of a product in the inventory.
func (a *App) IncreaseProductQuantity(product *product.Product) {
	a.Inventory.IncreaseProductQuantity(product)
}
