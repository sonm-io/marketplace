package service

import (
	"fmt"

	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/handler/proto"
)

// CreateAskOrder creates the given ask order.
func (ms *MarketService) CreateAskOrder(o pb.Order) error {
	order, err := ds.NewOrder(&o)
	if err != nil {
		return err
	}

	if order.SupplierID == "" {
		return fmt.Errorf("supplier is required")
	}

	return ms.createOrder(order)
}
