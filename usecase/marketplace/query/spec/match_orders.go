package spec

import (
	"fmt"
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// MatchOrders is a factory method that creates a spec for GetOrders query.
func MatchOrders(order ds.Order) (intf.CompositeSpecification, error) {
	// if no matching criteria is given return spec which returns false
	if order.SupplierID == "" && order.BuyerID == "" && order.Slot == nil {
		return intf.BaseSpecification{}, nil
	}

	// spec which returns true by default
	spec := (intf.BaseSpecification{}).Not()
	if order.SupplierID != "" {
		spec.And(NewSupplierIDEquals(order.SupplierID))
	}

	if order.BuyerID != "" {
		spec.And(NewBuyerIDEquals(order.BuyerID))
	}

	// nothing to match against further
	if order.Slot == nil {
		return spec, nil
	}

	switch order.OrderType {
	case ds.Ask:
		return spec.And(forAsk(*order.Slot)), nil
	case ds.Bid:
		return spec.And(forBid(*order.Slot)), nil
	default:
		return nil, fmt.Errorf("searching by any type is not supported")
	}
}

func forBid(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsBidOrder().
		And(NewGPUCountLessOrEqual(slot.Resources.GPUCount)).
		And(NewNetworkTypeLessOrEqual(slot.Resources.NetworkType))

	if slot.Resources.CPUCores > 0 {
		s.And(NewCPUCoresLessOrEqual(slot.Resources.CPUCores))
	}

	if slot.Resources.RAMBytes > 0 {
		s.And(NewRAMBytesLessOrEqual(slot.Resources.RAMBytes))
	}

	if slot.Resources.Storage > 0 {
		s.And(NewStorageLessOrEqual(slot.Resources.Storage))
	}

	if slot.Resources.NetTrafficIn > 0 {
		s.And(NewNetworkInboundLessOrEqual(slot.Resources.NetTrafficIn))
	}

	if slot.Resources.NetTrafficOut > 0 {
		s.And(NewNetworkOutboundLessOrEqual(slot.Resources.NetTrafficOut))
	}

	return s
}

func forAsk(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsAskOrder().
		And(NewCPUCoresGreaterOrEqual(slot.Resources.CPUCores)).
		And(NewRAMBytesGreaterOrEqual(slot.Resources.RAMBytes)).
		And(NewGPUCountGreaterOrEqual(slot.Resources.GPUCount)).
		And(NewStorageGreaterOrEqual(slot.Resources.Storage)).
		And(NewNetworkTypeGreaterOrEqual(slot.Resources.NetworkType)).
		And(NewNetworkInboundGreaterOrEqual(slot.Resources.NetTrafficIn)).
		And(NewNetworkOutboundGreaterOrEqual(slot.Resources.NetTrafficOut))

	return s
}
