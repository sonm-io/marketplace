package query

// GetOrdersResult is query result DTO.
type GetOrdersResult []Order


type Order struct {
	ID         string
	Price      int64
	SupplierID string
	BuyerID    string
	OrderType  int
}