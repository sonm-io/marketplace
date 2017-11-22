package query

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/entity/spec"
	"github.com/sonm-io/marketplace/usecase/intf"
)

type OrderBySpec interface {
	BySpecWithLimit(spec intf.Specification, limit uint64) ([]*entity.Order, error)
}

// GetOrdersHandler returns Orders by the given Specification.
type GetOrdersHandler struct {
	s OrderBySpec
}

// NewGetOrdersHandler creates a new instance of GetOrdersHandler
func NewGetOrdersHandler(s OrderBySpec) *GetOrdersHandler {
	return &GetOrdersHandler{s: s}
}

// Handle handles the given query and returns result.
// Retrieves an Order.
func (h *GetOrdersHandler) Handle(req intf.Query, result interface{}) error {

	q, ok := req.(GetOrders)
	if !ok {
		return fmt.Errorf("invalid query given")
	}

	r, ok := result.(*GetOrdersResult)
	if !ok {
		return fmt.Errorf("invalid result given")
	}

	/*
		s.compareCpuCoresBid(two) &&
		s.compareRamBytesBid(two) &&
		s.compareGpuCountBid(two) &&
		s.compareStorageBid(two) &&
		s.compareNetTrafficInBid(two) &&
		s.compareNetTrafficOutBid(two) &&
		s.compareNetworkTypeBid(two)
	*/

	orders, err := h.s.BySpecWithLimit(h.ForQuery(q), q.Limit)
	if err != nil {
		return err
	}

	var item Order
	for _, order := range orders {
		item = Order{
			ID:         order.ID,
			OrderType:  int(order.OrderType),
			Price:      order.Price,
			SupplierID: order.SupplierID,
			BuyerID:    order.BuyerID,
		}
		*r = append(*r, item)
	}

	result = r
	return err
}

//func (s *Slot) compareSupplierRating(two *Slot) bool {
//	return two.inner.GetSupplierRating() >= s.inner.GetSupplierRating()
//}

// TODO: (screwyprof) spec.ForQuery(q GetOrders) intf.Specification
func (h *GetOrdersHandler) ForQuery(q GetOrders) intf.Specification {
	var s intf.Specification
	switch entity.OrderType(q.OrderType) {
	case entity.ASK:
		s = spec.NewSupplierRatingGreaterOrEqualSpec(q.Slot.SupplierRating).
			And(spec.IsAskOrderSpec{})
	case entity.BID:
		s = spec.NewBuyerRatingGreaterOrEqualSpec(q.Slot.BuyerRating).
			And(spec.IsBidOrderSpec{})
	}

	return s
}
