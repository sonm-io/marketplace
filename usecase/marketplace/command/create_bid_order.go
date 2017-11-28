package command

// CreateBidOrder is a command to create a bid order.
type CreateBidOrder struct {
	// Order ID, UUIDv4
	ID string
	// Buyer's Ethereum ID
	BuyerID string
	// Order price
	Price int64
	// Slot a slot
	Slot Slot
}

// CommandID implements Command interface.
func (c CreateBidOrder) CommandID() string {
	return "CreateBidOrder"
}
