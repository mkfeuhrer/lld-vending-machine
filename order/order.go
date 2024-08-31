package order

import (
	"vendingmachine/order/enums"
	"vendingmachine/payment"
	"vendingmachine/product"
)

// Order Represents an order in the vending machine with its properties.
type Order struct {
	ID          string
	Amount      float64
	Products    []product.Product
	Payment     *payment.Payment
	OrderStatus enums.OrderStatus
}

func NewOrder(id string, products []product.Product, payment *payment.Payment) *Order {
	return &Order{
		ID:          id,
		Products:    products,
		Payment:     payment,
		OrderStatus: enums.Initiated,
	}
}
