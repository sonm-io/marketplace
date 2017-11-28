package storage

import (
	"github.com/sonm-io/marketplace/infra/storage/inmemory"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// OrderReadStorage stores and retrieves Orders (Read side).
type OrderReadStorage struct {
	e Engine
}

// NewOrderReadStorage creates an new instance of OrderReadStorage.
func NewOrderReadStorage(e Engine) *OrderReadStorage {
	return &OrderReadStorage{
		e: e,
	}
}

// Adds the given Order to the storage.
func (s *OrderReadStorage) Add(o *report.GetOrderReport) error {
	return s.e.Add(o, o.ID)
}

// Remove removes an Order with the given ID from OrderReadStorage.
// If no orders found, an error is returned.
func (s *OrderReadStorage) Remove(ID string) error {
	return s.e.Remove(ID)
}

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderReadStorage) ByID(ID string) (report.GetOrderReport, error) {

	el, err := s.e.Get(ID)
	if err != nil {
		return report.GetOrderReport{}, err
	}
	order := el.(*report.GetOrderReport)
	return *order, nil
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
func (s *OrderReadStorage) BySpecWithLimit(spec intf.Specification, limit uint64) (report.GetOrdersReport, error) {

	b := inmemory.NewBuilder()
	b.WithLimit(limit)
	b.WithSpec(spec)

	elements, err := s.e.Match(b.Build())
	if err != nil {
		return nil, err
	}

	var orders report.GetOrdersReport
	for _, el := range elements {
		order := el.(*report.GetOrderReport)
		orders = append(orders, *order)
	}

	return orders, nil
}
