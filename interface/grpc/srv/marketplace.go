package srv

import "github.com/sonm-io/marketplace/usecase/intf"

// Marketplace a GRPC Server implementing Marketplace API.
type Marketplace struct {
	commandBus   intf.CommandHandler
	orderByID    intf.QueryHandler
	ordersBySpec intf.QueryHandler
}

// NewMarketplace creates a new instance of Marketplace.
func NewMarketplace(c intf.CommandHandler, orderByID intf.QueryHandler, ordersBySpec intf.QueryHandler) *Marketplace {
	return &Marketplace{
		commandBus:   c,
		orderByID:    orderByID,
		ordersBySpec: ordersBySpec,
	}
}
