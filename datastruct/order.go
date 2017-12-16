package datastruct

// Order represents an order.
type Order struct {
	// Order ID, UUIDv4
	ID string
	// Buyer's EthereumID (Ask).
	BuyerID string
	// Supplier's is EthereumID (Bid).
	SupplierID string
	// Order price
	Price string
	// Order type (Bid or Ask)
	OrderType OrderType
	// Slot describes resource requirements
	Slot *Slot
}

// OrderType defines an order type.
type OrderType int32

// List of available order types.
const (
	Any OrderType = 0
	Bid           = 1
	Ask           = 2
)
