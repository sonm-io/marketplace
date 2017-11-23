package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/usecase/intf"
)

// CancelOrderStorage order storage.
type CancelOrderStorage interface {
	Remove(ID string) error
}

// CancelOrderHandler cancels orders.
type CancelOrderHandler struct {
	s CancelOrderStorage
}

// NewCancelOrderHandler creates a new instance of CancelOrderHandler.
func NewCancelOrderHandler(s CancelOrderStorage) CancelOrderHandler {
	return CancelOrderHandler{s: s}
}

// Handle handles the given command.
// Cancels the given order.
func (h CancelOrderHandler) Handle(cmd intf.Command) error {
	c, ok := cmd.(CancelOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	return h.s.Remove(c.ID)
}
