package query

import ds "github.com/sonm-io/marketplace/datastruct"

// GetOrders is a query to return order info.
type GetOrders struct {
	Order ds.Order

	// paging
	Limit uint64
}

// QueryID implements Query interface.
func (q GetOrders) QueryID() string {
	return "GetOrders"
}
