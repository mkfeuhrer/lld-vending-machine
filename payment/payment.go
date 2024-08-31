package payment

import "vendingmachine/payment/enums"

type Service interface {
	Debit()
	Credit()
}

// Payment Represents a payment in the vending machine with its properties.
type Payment struct {
	TotalAmount  float64
	Transactions []Transactions
}

type Transactions struct {
	Amount            float64
	PaymentInstrument enums.PaymentInstrument
}

func NewPayment(amount float64, transactions []Transactions) *Payment {
	return &Payment{
		TotalAmount:  amount,
		Transactions: transactions,
	}
}
