package sqllite

// Status order Status
type Status uint8

// List of available order statuses.
const (
	Cancelled Status = 0
	Created          = 1
	Expired          = 3
)

// Engine represents Storage Engine.
type Engine interface {
	InsertRow(query string, args ...interface{}) error
	UpdateRow(query string, value ...interface{}) error
}

// OrderStorage stores and retrieves Orders.
type OrderStorage struct {
	e Engine
}

// NewOrderStorage creates an new instance of OrderStorage.
func NewOrderStorage(e Engine) *OrderStorage {
	return &OrderStorage{
		e: e,
	}
}
