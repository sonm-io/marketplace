package spec

import (
	"fmt"
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// MatchOrders is a factory method that creates a spec for GetOrders query.
func MatchOrders(order ds.Order) (intf.CompositeSpecification, error) {
	switch order.OrderType {
	case ds.Ask:
		return forAsk(order), nil
	case ds.Bid:
		return forBid(order), nil
	default:
		return nil, fmt.Errorf("searching by any type is not supported")
	}
}

func forBid(order ds.Order) intf.CompositeSpecification {
	s := NewIsBidOrder()
	if order.BuyerID != "" {
		s = s.And(NewBuyerIDEquals(order.BuyerID))
	}

	if order.SupplierID != "" {
		s = s.And(NewSupplierIDEquals(order.SupplierID))
	}

	if order.Slot == nil {
		return s
	}

	slot := order.Slot

	s.And(NewGPUCountLessOrEqual(slot.Resources.GPUCount)).
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

func forAsk(order ds.Order) intf.CompositeSpecification {
	s := NewIsAskOrder()
	if order.BuyerID != "" {
		s = s.And(NewBuyerIDEquals(order.BuyerID))
	}

	if order.SupplierID != "" {
		s = s.And(NewSupplierIDEquals(order.SupplierID))
	}

	if order.Slot == nil {
		return s
	}

	slot := order.Slot
	s.And(NewCPUCoresGreaterOrEqual(slot.Resources.CPUCores)).
		And(NewRAMBytesGreaterOrEqual(slot.Resources.RAMBytes)).
		And(NewGPUCountGreaterOrEqual(slot.Resources.GPUCount)).
		And(NewStorageGreaterOrEqual(slot.Resources.Storage)).
		And(NewNetworkTypeGreaterOrEqual(slot.Resources.NetworkType)).
		And(NewNetworkInboundGreaterOrEqual(slot.Resources.NetTrafficIn)).
		And(NewNetworkOutboundGreaterOrEqual(slot.Resources.NetTrafficOut))

	return s
}
