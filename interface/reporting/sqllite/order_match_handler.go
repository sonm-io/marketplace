package sqllite

import (
	"database/sql"
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
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