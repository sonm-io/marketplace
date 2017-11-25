package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CreateAskOrderStorage adds an order to the storage.
type CreateAskOrderStorage interface {
	Store(o *entity.Order) error
}

// CreateAskOrderHandler creates new ask orders.
type CreateAskOrderHandler struct {
	s CreateAskOrderStorage
}

// NewCreateAskOrderHandler creates a new instance of CreateAskOrderHandler.
func NewCreateAskOrderHandler(s CreateAskOrderStorage) CreateAskOrderHandler {
	return CreateAskOrderHandler{s: s}
}

// Handle handles the given command.
// Creates the given bid order.
func (h CreateAskOrderHandler) Handle(cmd intf.Command) error {

	c, ok := cmd.(CreateAskOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	if entity.OrderType(c.OrderType) != entity.BID {
		return fmt.Errorf("invalid order type given: expected bid order, but got %v", c.OrderType)
	}

	order, err := newAskOrder(c)
	if err != nil {
		return err
	}

	return h.s.Store(order)
}

func newAskOrder(c CreateAskOrder) (*entity.Order, error) {

	res := entity.Resources(c.Slot.Resources)
	slot := entity.Slot{
		BuyerRating:    c.Slot.BuyerRating,
		SupplierRating: c.Slot.SupplierRating,
		Resources:      &res,
	}

	return entity.NewAskOrder(c.ID, c.SupplierID, c.Price, slot)
}
