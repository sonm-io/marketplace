package storage

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/infra/storage/inmemory"
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// Engine represents Storage Engine.
type Engine interface {
	Get(ID string) (interface{}, error)
	Add(el interface{}, ID string) error
	Remove(ID string) error
	Match(q inmemory.ConcreteCriteria) ([]interface{}, error)
}

// Storage stores and retrieves Orders.
type OrderStorage struct {
	e Engine
}

// NewOrderStorage create an new instance of Storage.
func NewOrderStorage(e Engine) *OrderStorage {
	return &OrderStorage{
		e: e,
	}
}

// Store saves the given Order.
func (s *OrderStorage) Add(o *entity.Order) error {
	return s.e.Add(o, o.ID)
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

	el, err := s.e.Get(ID)
	order := el.(*entity.Order)

	slot := report.Slot{}
	if order.Slot != nil {
		slot = report.Slot{}
		slot.BuyerRating = order.Slot.BuyerRating
		slot.SupplierRating = order.Slot.SupplierRating
		if order.Slot.Resources != nil {
			res := report.Resources(*order.Slot.Resources)
			slot.Resources = &res
		}
	}

	rep := report.GetOrderReport{
		ID:         order.ID,
		Price:      order.Price,
		OrderType:  int(order.OrderType),
		SupplierID: order.SupplierID,
		BuyerID:    order.BuyerID,
		Slot:       slot,
	}

	return rep, err
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) (report.GetOrdersReport, error) {

	b := inmemory.NewBuilder()
	b.WithLimit(limit)
	b.WithSpec(spec)

	elements, err := s.e.Match(b.Build())
	if err != nil {
		return nil, err
	}

	var orders report.GetOrdersReport
	var rep report.Order
	for _, el := range elements {
		order := el.(*entity.Order)
		entityToReport(order, &rep)
		orders = append(orders, rep)
	}

	return orders, nil
}

func entityToReport(entityOrder *entity.Order, reportOrder *report.Order) {

	var slot *report.Slot
	if entityOrder.Slot != nil {
		slot = &report.Slot{}
		slot.BuyerRating = entityOrder.Slot.BuyerRating
		slot.SupplierRating = entityOrder.Slot.SupplierRating
		if entityOrder.Slot.Resources != nil {
			res := report.Resources(*entityOrder.Slot.Resources)
			slot.Resources = &res
		}
	}

	reportOrder.ID = entityOrder.ID
	reportOrder.Price = entityOrder.Price
	reportOrder.OrderType = report.OrderType(entityOrder.OrderType)
	reportOrder.SupplierID = entityOrder.SupplierID
	reportOrder.BuyerID = entityOrder.BuyerID
	reportOrder.Slot = slot
}
