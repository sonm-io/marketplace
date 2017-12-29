package sqllite

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/interface/mapper"
	sds "github.com/sonm-io/marketplace/interface/mapper/datastruct"
)

// Status order Status
type Status uint8

// List of available order statuses.
const (
	InActive Status = 0
	Active          = 1
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

// Add adds the given Order to the storage.
func (s *OrderStorage) Add(o *ds.Order) error {
	if o == nil {
		return fmt.Errorf("cannot add an empty order")
	}

	row := sds.OrderRow{}
	mapper.OrderToRow(o, &row)
	row.Status = Active

	stmt := InsertOrderStmt(row)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := s.e.InsertRow(query, args...); err != nil {
		return fmt.Errorf("cannot add new order: %v", err)
	}
	return nil
}

// Cancel marks an Order with the given ID as InActive.
func (s *OrderStorage) Cancel(ID string) error {
	stmt := CancelOrderStmt(ID)
	query, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	if err := s.e.UpdateRow(query, args...); err != nil {
		return fmt.Errorf("cannot remove order: %v", err)
	}
	return nil
}
