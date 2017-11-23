package storage

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// Engine represents Storage Engine.
type Engine interface {
	ByID(ID string, result interface{}) error
	Store(o *entity.Order) error
	Remove(ID string) error
	Match(q inmemory.ConcreteCriteria, result interface{}) error
}

// OrderStorage stores and retrieves Orders.
type OrderStorage struct {
	e Engine
}

// NewOrderStorage create an new instance of OrderStorage.
func NewOrderStorage(e Engine) *OrderStorage {
	return &OrderStorage{
		e: e,
	}
}

// Store saves the given Order.
func (s *OrderStorage) Store(o *entity.Order) error {
	return s.e.Store(o)
}

// Remove deletes an Order with the given ID from Storage.
// If no orders found, an error is returned.
func (s *OrderStorage) Remove(ID string) error {
	return s.e.Remove(ID)
}

//
// READ SIDE
//

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderStorage) ByID(ID string) (report.GetOrderReport, error) {
	var order report.GetOrderReport
	err := s.e.ByID(ID, &order)

	return order, err
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) (report.GetOrdersReport, error) {

	b := inmemory.NewBuilder()
	b.WithLimit(limit)
	b.WithSpec(spec)

	var orders report.GetOrdersReport
	err := s.e.Match(b.Build(), &orders)
	if err != nil {
		return nil, err
	}
	return orders, nil
}
