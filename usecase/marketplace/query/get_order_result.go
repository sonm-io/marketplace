package query

// GetOrderResult is query result DTO.
type GetOrderResult struct {
	ID         string
	Price      int64
	SupplierID string
	BuyerID    string
	OrderType  int
}
