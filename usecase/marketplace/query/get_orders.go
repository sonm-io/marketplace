package query

// GetOrders is a query to return order info.
type GetOrders struct {
	OrderType int
	Slot

	// paging
	Limit uint64
}

// Slot request DTO
type Slot struct {
	// Buyer’s rating. Got from Buyer’s profile for BID orders rating_supplier.
	BuyerRating int64
	// Supplier’s rating. Got from Supplier’s profile for ASK orders.
	SupplierRating int64
	// Geo represent Worker's position
	//Geo *Geo `protobuf:"bytes,3,opt,name=geo" json:"geo,omitempty"`
	// Hardware resources requirements
	//Resources *Resources `protobuf:"bytes,4,opt,name=resources" json:"resources,omitempty"`
}

// QueryID implements Query interface.
func (q GetOrders) QueryID() string {
	return "GetOrders"
}
