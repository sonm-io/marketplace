package report

// Order represents an order.
type Order struct {
	// Order ID, UUIDv4
	ID string
	// Buyer's EtherumID (ASK)
	BuyerID string
	// Supplier's is EtherumID (BID)
	SupplierID string
	// Order price
	Price int64
	// Order type (Bid or Ask)
	OrderType OrderType
	// Slot describes resource requirements
	Slot *Slot
}

type OrderType int32

const (
	ANY OrderType = 0
	BID OrderType = 1
	ASK OrderType = 2
)
