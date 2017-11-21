package command

// CreateOrder is a command to create an order.
type CreateOrder struct {
	// Order ID, UUIDv4
	ID string
	// Buyer's Ethereum ID
	BuyerID string
	// Supplier's Ethereum ID
	SupplierID string
	// Order price
	Price int64
	// Order type
	OrderType int
}

// CommandID implements Command interface.
func (c CreateOrder) CommandID() string {
	return "CreateOrder"
}
