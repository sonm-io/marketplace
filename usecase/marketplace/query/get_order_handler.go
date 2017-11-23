package query

import (
	"fmt"

	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// OrderByIDStorage fetches an Order by the given ID.
type OrderByIDStorage interface {
	ByID(id string) (report.GetOrderReport, error)
}

// GetOrderHandler returns an Order.
type GetOrderHandler struct {
	s OrderByIDStorage
}

// NewGetOrderHandler creates a new instance of GetOrderHandler.
func NewGetOrderHandler(s OrderByIDStorage) *GetOrderHandler {
	return &GetOrderHandler{s: s}
}

// Handle handles the given query and returns result.
// Retrieves an Order.
func (h *GetOrderHandler) Handle(req intf.Query, result interface{}) error {

	q, ok := req.(GetOrder)
	if !ok {
		return fmt.Errorf("invalid query %v given", req)
	}

	r, ok := result.(*report.GetOrderReport)
	if !ok {
		return fmt.Errorf("invalid result %v given", result)
	}

	order, err := h.s.ByID(q.ID)
	if err != nil {
		return err
	}

	*r = order

	return nil
}
