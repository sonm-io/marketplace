package command

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CreateBidOrderStorage adds an order to the storage.
type CreateBidOrderStorage interface {
	Add(o *ds.Order) error
}

// CreateBidOrderHandler creates new bid orders.
type CreateBidOrderHandler struct {
	s CreateBidOrderStorage
}

// NewCreateBidOrderHandler creates a new instance of CreateBidOrderHandler.
func NewCreateBidOrderHandler(s CreateBidOrderStorage) *CreateBidOrderHandler {
	return &CreateBidOrderHandler{s: s}
}

// Handle handles the given command.
// Creates the given bid order.
func (h *CreateBidOrderHandler) Handle(cmd intf.Command) error {

	c, ok := cmd.(CreateBidOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	order, err := entity.NewBidOrder(c.ID, c.BuyerID, c.SupplierID, c.Price, c.Slot)
	if err != nil {
		return err
	}

	return h.s.Add(&order.Order)
}
