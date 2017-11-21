package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/usecase/intf"
)

type cancelOrderStorage interface {
	Remove(ID string) error
}

// CancelOrderHandler stores orders.
type CancelOrderHandler struct {
	s cancelOrderStorage
}

// NewCancelOrderHandler creates a new instance of CancelOrderHandler.
func NewCancelOrderHandler(r cancelOrderStorage) CancelOrderHandler {
	return CancelOrderHandler{s: r}
}
// Handle handles the given command.
// Stores the given order.
func (h CancelOrderHandler) Handle(cmd intf.Command) error {
	c, ok := cmd.(CancelOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	return h.s.Remove(c.ID)
}
