package sqllite

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// OrderRowFetcher fetches order rows from storage.
type OrderRowsFetcher interface {
	FetchRows(rows interface{}, query string, value ...interface{}) error
}

// MatchOrdersHandler returns Orders by the given CompositeSpecification.
type MatchOrdersHandler struct {
	rf OrderRowsFetcher
}

// NewMatchOrdersHandler creates a new instance of MatchOrdersHandler.
func NewMatchOrdersHandler(rf OrderRowsFetcher) *MatchOrdersHandler {
	return &MatchOrdersHandler{rf: rf}
}

// Handle handles the given query and returns result.
// Retrieves Orders by the given Spec.
//
// if req.Limit is > 0, then only the given number of Orders will be returned.
func (h *MatchOrdersHandler) Handle(req intf.Query, result interface{}) error {
	q, ok := req.(query.GetOrders)
	if !ok {
		return fmt.Errorf("invalid query %v given", req)
	}

	r, ok := result.(*report.GetOrdersReport)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	stmt, err := MatchOrdersStmt(q.Order, q.Limit)
	if err != nil {
		return err
	}

	sql, args, err := ToSQL(stmt)
	if err != nil {
		return err
	}

	var rows sds.OrderRows
	if err := h.rf.FetchRows(&rows, sql, args...); err != nil {
		return err
	}

	var order ds.Order
	for idx := range rows {
		order = ds.Order{}
		orderFromRow(&order, &rows[idx])
		*r = append(*r, report.GetOrderReport{Order: order})
	}

	return nil
}
