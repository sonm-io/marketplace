package report

// GetOrderReport is query result DTO.
type GetOrderReport struct {
	ID         string
	Price      int64
	SupplierID string
	BuyerID    string
	OrderType  int
	Slot       Slot
}
