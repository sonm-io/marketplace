package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CreateBidOrderStorage adds an order to the storage.
type CreateBidOrderStorage interface {
	Add(o *entity.Order) error
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

	order, err := newBidOrder(c)
	if err != nil {
		return err
	}

	return h.s.Add(order)
}

func newBidOrder(c CreateBidOrder) (*entity.Order, error) {

	res := entity.Resources(c.Slot.Resources)
	slot := entity.Slot{
		BuyerRating:    c.Slot.BuyerRating,
		SupplierRating: c.Slot.SupplierRating,
		Resources:      &res,
	}

	return entity.NewBidOrder(c.ID, c.BuyerID, c.Price, slot)
}
