package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

type CreateOrderStorage interface {
	Store(o *entity.Order) error
}

// CreateOrderHandler creates new orders.
type CreateOrderHandler struct {
	s CreateOrderStorage
}

// NewCreateOrderHandler creates a new instance of CreateOrderHandler.
func NewCreateOrderHandler(s CreateOrderStorage) CreateOrderHandler {
	return CreateOrderHandler{s: s}
}

// Handle handles the given command.
// Stores the given order.
func (h CreateOrderHandler) Handle(cmd intf.Command) error {

	c, ok := cmd.(CreateOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	order, err := h.createNewOrder(c)
	if err != nil {
		return err
	}

	return h.s.Store(order)
}

func (h CreateOrderHandler) createNewOrder(c CreateOrder) (*entity.Order, error) {
	var (
		order *entity.Order
		err   error
	)
	switch entity.OrderType(c.OrderType) {
	case entity.ASK:
		order, err = entity.NewAskOrder(c.ID, c.BuyerID, c.Price)
	case entity.BID:
		order, err = entity.NewBidOrder(c.ID, c.SupplierID, c.Price)
	default:
		err = fmt.Errorf("invalid order type given %v", c.OrderType)
	}
	return order, err
}
