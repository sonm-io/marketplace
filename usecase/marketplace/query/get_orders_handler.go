package query

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/spec"
)

// OrderBySpecStorage fetches reports by the given criteria
type OrderBySpecStorage interface {
	BySpecWithLimit(spec intf.Specification, limit uint64) ([]ds.Order, error)
}

// GetOrdersHandler returns Orders by the given CompositeSpecification.
type GetOrdersHandler struct {
	s OrderBySpecStorage
}

// NewGetOrdersHandler creates a new instance of GetOrdersHandler.
func NewGetOrdersHandler(s OrderBySpecStorage) *GetOrdersHandler {
	return &GetOrdersHandler{s: s}
}

// Handle handles the given query and returns result.
// Retrieves Orders by the given Spec.
func (h *GetOrdersHandler) Handle(req intf.Query, result interface{}) error {

	q, ok := req.(GetOrders)
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

	orders, err := h.s.BySpecWithLimit(s, q.Limit)
	if err != nil {
		return err
	}

	for _, order := range orders {
		*r = append(*r, report.GetOrderReport{Order: order})
	}

	return err
}
