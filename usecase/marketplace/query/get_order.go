package query

// GetOrder is a query to return order info.
type GetOrder struct{
	ID string
}

// QueryID implements Query interface.
func (q GetOrder) QueryID() string {
	return "GetOrder"
}