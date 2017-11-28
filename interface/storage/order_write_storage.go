package storage

import (
	"github.com/sonm-io/marketplace/entity"
)

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

// Add adds the given Order to the OrderStorage.
func (s *OrderStorage) Add(o *entity.Order) error {
	return s.e.Add(o, o.ID)
}

// Remove removes an Order with the given ID from OrderStorage.
// If no orders found, an error is returned.
func (s *OrderStorage) Remove(ID string) error {
	return s.e.Remove(ID)
}

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderStorage) ByID(ID string) (*entity.Order, error) {

	el, err := s.e.Get(ID)
	if err != nil {
		return nil, err
	}

	order := el.(*entity.Order)
	return order, nil
}
