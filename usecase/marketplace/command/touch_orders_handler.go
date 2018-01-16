package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/usecase/intf"
)

// OrderToucher updates orders' TTL.
type OrderToucher interface {
	Touch(IDs []string) error
}

// TouchOrdersHandler cancels orders.
type TouchOrdersHandler struct {
	t OrderToucher
}

// NewTouchOrdersHandler creates a new instance of TouchOrdersHandler.
func NewTouchOrdersHandler(t OrderToucher) TouchOrdersHandler {
	return TouchOrdersHandler{t: t}
}

// Handle handles the given command.
// Updates orders' TTL.
func (h TouchOrdersHandler) Handle(cmd intf.Command) error {
	c, ok := cmd.(TouchOrders)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	return h.t.Touch(c.IDs)
}
