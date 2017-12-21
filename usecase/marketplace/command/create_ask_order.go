package command

import ds "github.com/sonm-io/marketplace/datastruct"

// CreateAskOrder is a command to create an ask order.
type CreateAskOrder struct {
	// Order ID, UUIDv4
	ID string
	// Supplier's Ethereum ID
	SupplierID string
	// Buyer's Ethereum ID
	BuyerID string
	// Order price
	Price string
	// Slot a slot
	Slot ds.Slot
}

// CommandID implements Command interface.
func (c CreateAskOrder) CommandID() string {
	return "CreateAskOrder"
}
