package command

import (
	"fmt"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CreateAskOrderStorage adds an order to the storage.
type CreateAskOrderStorage interface {
	Add(o *ds.Order) error
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

	order, err := entity.NewAskOrder(c.ID, c.SupplierID, c.Price, c.Slot)
	if err != nil {
		return err
	}

	return h.s.Add(&order.Order)
}
