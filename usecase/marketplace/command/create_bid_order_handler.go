package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CreateBidOrderStorage adds an order to the storage.
type CreateBidOrderStorage interface {
	Store(o *entity.Order) error
}

// CreateBidOrderHandler creates new bid orders.
type CreateBidOrderHandler struct {
	s CreateBidOrderStorage
}

// NewCreateBidOrderHandler creates a new instance of CreateBidOrderHandler.
func NewCreateBidOrderHandler(s CreateBidOrderStorage) CreateBidOrderHandler {
	return CreateBidOrderHandler{s: s}
}

// Handle handles the given command.
// Creates the given bid order.
func (h CreateBidOrderHandler) Handle(cmd intf.Command) error {

	c, ok := cmd.(CreateBidOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	if entity.OrderType(c.OrderType) != entity.BID {
		return fmt.Errorf("invalid order type given: expected bid order, but got %v", c.OrderType)
	}

	order, err := entity.NewBidOrder(c.ID, c.BuyerID, c.Price, entity.Slot(c.Slot))
	if err != nil {
		return err
	}

	return h.s.Store(order)
}
