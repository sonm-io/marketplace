package sqllite

import (
	"database/sql"
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// OrderRowFetcher fetches order row from storage.
type OrderRowFetcher interface {
	FetchRow(ID string, row *sds.OrderRow) error
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

	var row sds.OrderRow
	if err := h.rf.FetchRow(q.ID, &row); err != nil {
		if err == sql.ErrNoRows {
			return fmt.Errorf("order %s is not found", q.ID)
		}
		return fmt.Errorf("an error occured: %v", err)
	}

	order := ds.Order{}
	orderFromRow(&order, &row)

	if row.Status == uint8(InActive) {
		return fmt.Errorf("order %s is inactive", q.ID)
	}

	return nil
}
