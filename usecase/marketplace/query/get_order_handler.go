package query

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

type OrderByID interface {
	ByID(id string) (*entity.Order, error)
}

// GetOrderHandler returns an Order.
type GetOrderHandler struct {
	s OrderByID
}

// NewGetOrderHandler creates a new instance of GetOrderHandler
func NewGetOrderHandler(s OrderByID) *GetOrderHandler {
	return &GetOrderHandler{s: s}
}

// Handle handles the given query and returns result.
// Retrieves an Order.
func (h *GetOrderHandler) Handle(req intf.Query, result interface{}) error {

	q, ok := req.(GetOrder)
	if !ok {
		return fmt.Errorf("invalid query given")
	}

	r, ok := result.(*GetOrderResult)
	if !ok {
		return fmt.Errorf("invalid result given")
	}

	order, err := h.s.ByID(q.ID)
	if err != nil {
		return err
	}

	r.ID = order.ID
	r.Price = order.Price
	r.SupplierID = order.SupplierID
	r.BuyerID = order.BuyerID
	r.OrderType = int(order.OrderType)

	return nil
}
