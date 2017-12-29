package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/usecase/intf"
)

// OrderCanceler cancels orders.
type OrderCanceler interface {
	Cancel(ID string) error
}

// CancelOrderHandler cancels orders.
type CancelOrderHandler struct {
	s OrderCanceler
}

// NewCancelOrderHandler creates a new instance of CancelOrderHandler.
func NewCancelOrderHandler(s OrderCanceler) CancelOrderHandler {
	return CancelOrderHandler{s: s}
}

// Handle handles the given command.
// Cancels the given order.
func (h CancelOrderHandler) Handle(cmd intf.Command) error {
	c, ok := cmd.(CancelOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	return h.s.Cancel(c.ID)
}
