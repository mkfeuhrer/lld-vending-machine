package enums

// PaymentInstrument Represents a payment instrument in the vending machine.
type PaymentInstrument string

const (
	Coin  PaymentInstrument = "coin"
	Notes                   = "notes"
	None                    = ""
)

func (p PaymentInstrument) String() string {
	switch p {
	case Coin:
		return "Coin"
	case Notes:
		return "Notes"
	default:
		return None
	}
}

func (p PaymentInstrument) Valid() bool {
	switch p {
	case Coin, Notes:
		return true
	}
	return false
}
