package command

import ds "github.com/sonm-io/marketplace/datastruct"

// CreateBidOrder is a command to create a bid order.
type CreateBidOrder struct {
	// Order ID, UUIDv4
	ID string
	// Buyer's Ethereum ID
	BuyerID string
	// Supplier's Ethereum ID
	SupplierID string
	// Order price
	Price string
	// Slot a slot
	Slot ds.Slot
}

// CommandID implements Command interface.
func (c CreateBidOrder) CommandID() string {
	return "CreateBidOrder"
}
