package command

// CancelOrder is a command to update orders' ttl.
type TouchOrders struct {
	IDs []string
}

// CommandID implements Command interface.
func (c TouchOrders) CommandID() string {
	return "TouchOrders"
}
