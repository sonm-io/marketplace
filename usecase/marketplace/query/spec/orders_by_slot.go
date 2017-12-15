package spec

import (
	"fmt"
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// OrdersBySlot is a factory method that creates a spec for GetOrders query.
func OrdersBySlot(orderType ds.OrderType, slot ds.Slot) (intf.CompositeSpecification, error) {
	switch orderType {
	case ds.Ask:
		return forAsk(slot), nil
	case ds.Bid:
		return forBid(slot), nil
	default:
		return nil, fmt.Errorf("invalid order type %v given", orderType)
	}
}

func forBid(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsBidOrder().
		//And(NewBuyerRatingGreaterOrEqual(slot.BuyerRating)).
		And(NewSupplierRatingGreaterOrEqual(slot.SupplierRating)).
		And(NewCPUCoresGreaterOrEqual(slot.Resources.CPUCores)).
		And(NewRAMBytesGreaterOrEqual(slot.Resources.RAMBytes)).
		And(NewGPUCountGreaterOrEqual(slot.Resources.GPUCount)).
		And(NewStorageGreaterOrEqual(slot.Resources.Storage)).
		And(NewNetworkTypeGreaterOrEqual(slot.Resources.NetworkType)).
		And(NewNetworkInboundGreaterOrEqual(slot.Resources.NetTrafficIn)).
		And(NewNetworkOutboundGreaterOrEqual(slot.Resources.NetTrafficOut))

	return s
}

func forAsk(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsAskOrder().
		And(NewSupplierRatingGreaterOrEqual(slot.SupplierRating)).
		And((NewCPUCoresGreaterOrEqual(slot.Resources.CPUCores)).Not()).
		And((NewRAMBytesGreaterOrEqual(slot.Resources.RAMBytes)).Not()).
		And((NewGPUCountGreaterOrEqual(slot.Resources.GPUCount)).Not()).
		And((NewStorageGreaterOrEqual(slot.Resources.Storage)).Not()).
		And((NewNetworkTypeGreaterOrEqual(slot.Resources.NetworkType)).Not()).
		And((NewNetworkInboundGreaterOrEqual(slot.Resources.NetTrafficIn)).Not()).
		And((NewNetworkOutboundGreaterOrEqual(slot.Resources.NetTrafficOut)).Not())

	return s
}
