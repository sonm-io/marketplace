package report

// GetOrderReport is query result DTO.
type GetOrderReport struct {
	ID         string
	Price      int64
	SupplierID string
	BuyerID    string
	OrderType  int
	Slot       *Slot
}

const (
	ANY = 0
	BID = 1
	ASK = 2
)
