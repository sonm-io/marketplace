package srv

import "github.com/sonm-io/marketplace/usecase/intf"

type Marketplace struct {
	commandBus   intf.CommandHandler
	orderByID    intf.QueryHandler
	ordersBySpec intf.QueryHandler
}

func NewMarketplace(c intf.CommandHandler, orderByID intf.QueryHandler, ordersBySpec intf.QueryHandler) *Marketplace {
	return &Marketplace{
		commandBus:   c,
		orderByID:    orderByID,
		ordersBySpec: ordersBySpec,
	}
}
