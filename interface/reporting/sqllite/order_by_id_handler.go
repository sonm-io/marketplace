package sqllite

import (
	"fmt"

	"github.com/gocraft/dbr"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/interface/mapper"
	sds "github.com/sonm-io/marketplace/interface/mapper/datastruct"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// OrderRowFetcher fetches order row from storage.
type OrderRowFetcher interface {
	FetchRow(row interface{}, query string, value ...interface{}) error
}

// OrderByIDHandler returns an Order by its ID.
type OrderByIDHandler struct {
	rf OrderRowFetcher
}

// NewOrderByIDHandler creates an new instance of OrderByID.
func NewOrderByIDHandler(rf OrderRowFetcher) *OrderByIDHandler {
	return &OrderByIDHandler{
		rf: rf,
	}
}

// Handle handles the given query and returns result.
// Retrieves an Order.
func (h *OrderByIDHandler) Handle(req intf.Query, result interface{}) error {
	q, ok := req.(query.GetOrder)
	if !ok {
		return fmt.Errorf("invalid query %v given", req)
	}

	r, ok := result.(*report.GetOrderReport)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	stmt, err := GetOrderByIDStmt(q.ID, OrderTTL)
	if err != nil {
		return err
	}

	sql, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	var row sds.OrderRow
	if err := h.rf.FetchRow(&row, sql, args...); err != nil {
		if err == dbr.ErrNotFound {
			return fmt.Errorf("order %s is not found", q.ID)
		}
		return fmt.Errorf("an error occured: %v", err)
	}

	if row.Status == uint8(InActive) {
		return fmt.Errorf("order %s is inactive", q.ID)
	}

	order := ds.Order{}
	mapper.OrderFromRow(&order, &row)

	(*r).Order = order

	return nil
}
