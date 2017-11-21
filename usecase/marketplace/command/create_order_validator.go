package command

import (
	"fmt"

	"github.com/sonm-io/marketplace/usecase/intf"
)

var (
	errPriceIsZero       = fmt.Errorf("order price cannot be negative or zero")
)

// CreateOrderValidator validates CreateOrder command.
// acts as a command handler decorator.
type CreateOrderValidator struct {
	h intf.CommandHandler
}

// NewCreateOrderValidator creates a new instance of CreateOrderValidator.
func NewCreateOrderValidator(h intf.CommandHandler) CreateOrderValidator {
	return CreateOrderValidator{h:h}
}
// Handle handles the given command.
// Stores the given order.
func (h CreateOrderValidator) Handle(cmd intf.Command) error {

	c, ok := cmd.(CreateOrder)
	if !ok {
		return fmt.Errorf("invalid command %v given", cmd)
	}

	if c.Price <= 0 {
		return errPriceIsZero
	}


	return h.Handle(cmd)
}

