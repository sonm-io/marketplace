package service

import (
	"fmt"

	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/proto"
)

// CreateBidOrder creates the given bid order.
func (ms *MarketService) CreateBidOrder(o pb.Order) error {
	order, err := ds.NewOrder(&o)
	if err != nil {
		return err
	}

	if order.ByuerID == "" {
		return fmt.Errorf("buyer is required")
	}

	return ms.createOrder(order)
}
