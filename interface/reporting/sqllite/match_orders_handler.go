package sqllite

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	sds "github.com/sonm-io/marketplace/infra/storage/sqllite/datastruct"

	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/spec"
)

// OrderRowFetcher fetches order rows from storage.
type OrderRowsFetcher interface {
	FetchAll() (sds.OrderRows, error)
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
// WARNING: At nonce all the Orders will be loaded in memory and after that filtered.
func (h *MatchOrdersHandler) Handle(req intf.Query, result interface{}) error {
	q, ok := req.(query.GetOrders)
	if !ok {
		return fmt.Errorf("invalid query %v given", req)
	}

	r, ok := result.(*report.GetOrdersReport)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	s, err := spec.MatchOrders(q.Order)
	if err != nil {
		return err
	}

	// spec is empty, nothing to return
	if s == nil {
		return nil
	}

	rows, err := h.rf.FetchAll()
	if err != nil {
		return err
	}

	var (
		order  ds.Order
		orders []ds.Order
	)

	for idx := range rows {
		if q.Limit > 0 && uint64(len(orders)) >= q.Limit {
			break
		}

		order = ds.Order{}
		orderFromRow(&order, &rows[idx])

		if s.IsSatisfiedBy(&order) {
			orders = append(orders, order)
		}
	}

	for _, order := range orders {
		*r = append(*r, report.GetOrderReport{Order: order})
	}

	return err
}
