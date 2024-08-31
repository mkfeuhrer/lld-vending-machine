package balance

import (
	"fmt"
	"vendingmachine/payment/enums"
)

// Balance Represents the balance of payments collected and change debited in the vending machine.
type Balance struct {
	PaymentInstrumentDetails map[enums.PaymentInstrument]int
}

// AddPaymentInstrument Adds a payment instrument to the balance.
func (b *Balance) AddPaymentInstrument(paymentInstrument enums.PaymentInstrument, quantity int) {
	b.PaymentInstrumentDetails[paymentInstrument] += quantity
}

// RemovePaymentInstrument Removes a payment instrument from the balance.
func (b *Balance) RemovePaymentInstrument(paymentInstrument enums.PaymentInstrument, quantity int) {
	availableQuantity, ok := b.PaymentInstrumentDetails[paymentInstrument]
	if !ok {
		return
	}

	if availableQuantity >= quantity {
		b.PaymentInstrumentDetails[paymentInstrument] -= quantity
	} else {
		// todo: error handling
		return
	}
}

// IsPaymentInstrumentAvailable Checks if a payment instrument is available in the balance.
func (b *Balance) IsPaymentInstrumentAvailable(paymentInstrument enums.PaymentInstrument) bool {
	availableQuantity, ok := b.PaymentInstrumentDetails[paymentInstrument]
	if ok && availableQuantity > 0 {
		return true
	}
	return false
}

// DisplayBalance Displays the current balance of the vending machine.
func (b *Balance) DisplayBalance() {
	fmt.Println("Current Balance:")
	fmt.Println("---------------------------------------")
	for paymentInstrument, quantity := range b.PaymentInstrumentDetails {
		fmt.Printf("%s: %d\n", paymentInstrument.String(), quantity)
	}
	fmt.Println("---------------------------------------")
}

// NewBalance Creates a new balance.
func NewBalance(paymentInstrumentDetails map[enums.PaymentInstrument]int) *Balance {
	return &Balance{
		PaymentInstrumentDetails: paymentInstrumentDetails,
	}
}
