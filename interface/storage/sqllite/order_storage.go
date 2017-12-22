package sqllite

import (
	"database/sql"
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
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
	FetchRow(ID string, row *sds.OrderRow) error
	FetchAll() (sds.OrderRows, error)
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

// ByID Fetches an Order by its ID.
// If ID is not found, an error is returned.
func (s *OrderStorage) ByID(ID string) (ds.Order, error) {
	var row sds.OrderRow
	if err := s.e.FetchRow(ID, &row); err != nil {
		if err == sql.ErrNoRows {
			return ds.Order{}, fmt.Errorf("order %s is not found", ID)
		}
		return ds.Order{}, fmt.Errorf("an error occured: %v", err)
	}

	order := ds.Order{}
	orderFromRow(&order, &row)

	if row.Status == uint8(InActive) {
		return ds.Order{}, fmt.Errorf("order %s is inactive", ID)
	}

	return order, nil
}

// BySpecWithLimit fetches Orders that satisfy the given Spec.
// if limit is > 0, then only the given number of Orders will be returned.
// WARNING: At nonce all the Orders will be loaded in memory and after that filtered.
// TODO: (screwyprof) Add filtering without loading all the records.
func (s *OrderStorage) BySpecWithLimit(spec intf.Specification, limit uint64) ([]ds.Order, error) {
	rows, err := s.e.FetchAll()
	if err != nil {
		return nil, err
	}

	var (
		order  ds.Order
		orders []ds.Order
	)
	for idx := range rows {
		if limit > 0 && uint64(len(orders)) >= limit {
			break
		}

		order = ds.Order{}
		orderFromRow(&order, &rows[idx])

		if spec.IsSatisfiedBy(&order) {
			orders = append(orders, order)
		}
	}

	return orders, nil
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

func orderFromRow(order *ds.Order, row *sds.OrderRow) {
	if order == nil {
		return
	}

	order.ID = row.ID
	order.OrderType = ds.OrderType(row.Type)
	order.Price = row.Price

	order.BuyerID = row.BuyerID
	order.SupplierID = row.SupplierID

	if order.Slot == nil {
		order.Slot = &ds.Slot{}
	}

	slot := &ds.Slot{
		BuyerRating:    row.BuyerRating,
		SupplierRating: row.SupplierRating,

		Resources: ds.Resources{
			CPUCores: row.CPUCores,
			RAMBytes: row.RAMBytes,
			GPUCount: ds.GPUCount(row.GPUCount),
			Storage:  row.Storage,

			NetworkType:   ds.NetworkType(row.NetType),
			NetTrafficIn:  row.NetInbound,
			NetTrafficOut: row.NetOutbound,

			Properties: row.Properties,
		},
	}

	order.Slot = slot
}
