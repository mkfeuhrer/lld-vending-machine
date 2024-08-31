package enums

// OrderStatus Represents the state of an order in the vending machine.
type OrderStatus string

const (
	Initiated OrderStatus = "initiated"
	Pending   OrderStatus = "pending"
	Completed OrderStatus = "completed"
	Failed    OrderStatus = "failed"
	Unknown               = "unknown"
)

func (os OrderStatus) String() string {
	return string(os)
}

func (os OrderStatus) Valid() bool {
	switch os {
	case Initiated, Pending, Completed, Failed:
		return true
	}
	return false
}
