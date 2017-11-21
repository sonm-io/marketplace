package command

// CancelOrder is a command to cancel an order.
type CancelOrder struct {
	// Order ID, UUIDv4
	ID string
}

// CommandID implements Command interface.
func (c CancelOrder) CommandID() string {
	return "CancelOrder"
}
