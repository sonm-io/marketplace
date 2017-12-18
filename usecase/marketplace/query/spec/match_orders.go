package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// MatchOrders is a factory method that creates a spec for GetOrders query.
func MatchOrders(order ds.Order) (intf.CompositeSpecification, error) {
	// if Slot is not set, then match only by SupplierID or BuyerID.
	if order.Slot == nil {
		var spec intf.CompositeSpecification
		switch {
		case order.SupplierID != "":
			spec = NewSupplierIDEquals(order.SupplierID)
		case order.BuyerID != "":
			spec = NewBuyerIDEquals(order.BuyerID)
		default:
			spec = intf.BaseSpecification{}
		}
		return spec, nil
	}

	switch order.OrderType {
	case ds.Ask:
		return forAsk(*order.Slot), nil
	case ds.Bid:
		return forBid(*order.Slot), nil
	default:
		// return all orders
		return NewIsBidOrder().Or(NewIsAskOrder()), nil
	}
}

func forBid(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsBidOrder().
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
		And(NewCPUCoresLessOrEqual(slot.Resources.CPUCores)).
		And(NewRAMBytesLessOrEqual(slot.Resources.RAMBytes)).
		And(NewGPUCountLessOrEqual(slot.Resources.GPUCount)).
		And(NewStorageLessOrEqual(slot.Resources.Storage)).
		And(NewNetworkTypeLessOrEqual(slot.Resources.NetworkType)).
		And(NewNetworkInboundLessOrEqual(slot.Resources.NetTrafficIn)).
		And(NewNetworkOutboundLessOrEqual(slot.Resources.NetTrafficOut))

	return s
}
