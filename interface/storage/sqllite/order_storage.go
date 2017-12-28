package sqllite

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"
)

type Status uint8

const (
	InActive Status = 0
	Active          = 1
)

// Engine represents Storage Engine.
type Engine interface {
	InsertRow(row *sds.OrderRow) error
	UpdateStatus(ID string, status uint8) error
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

// Add adds the given Order to the storage.
func (s *OrderStorage) Add(o *ds.Order) error {
	if o == nil {
		return fmt.Errorf("cannot add an empty order")
	}

	row := sds.OrderRow{}
	orderToRow(o, &row)

	if err := s.e.InsertRow(&row); err != nil {
		return fmt.Errorf("cannot add new order: %v", err)
	}
	return nil
}

// Remove removes an Order with the given ID from OrderStorage.
func (s *OrderStorage) Remove(ID string) error {
	if err := s.e.UpdateStatus(ID, uint8(InActive)); err != nil {
		return fmt.Errorf("cannot remove order: %v", err)
	}
	return nil
}

func orderToRow(order *ds.Order, row *sds.OrderRow) {
	row.ID = order.ID
	row.Type = int32(order.OrderType)
	row.BuyerID = order.BuyerID
	row.SupplierID = order.SupplierID
	row.Price = order.Price

	if order.Slot == nil {
		order.Slot = &ds.Slot{}
	}

	row.Duration = order.Slot.Duration
	row.BuyerRating = order.Slot.BuyerRating
	row.SupplierRating = order.Slot.SupplierRating

	row.CPUCores = order.Slot.Resources.CPUCores
	row.RAMBytes = order.Slot.Resources.RAMBytes
	row.GPUCount = uint64(order.Slot.Resources.GPUCount)
	row.Storage = order.Slot.Resources.Storage

	row.NetType = uint64(order.Slot.Resources.NetworkType)
	row.NetInbound = order.Slot.Resources.NetTrafficIn
	row.NetOutbound = order.Slot.Resources.NetTrafficOut

	row.Properties = sds.Properties(order.Slot.Resources.Properties)
}
