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

	/*
		//s.compareCpuCoresBid(two) &&
		//s.compareRamBytesBid(two) &&
		s.compareGpuCountBid(two) &&
		s.compareStorageBid(two) &&
		s.compareNetTrafficInBid(two) &&
		s.compareNetTrafficOutBid(two) &&
		s.compareNetworkTypeBid(two)
	*/

}

func forBid(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsBidOrder().
		And(NewBuyerRatingGreaterOrEqual(slot.BuyerRating)).
		And(NewCPUCoresGreaterOrEqual(slot.Resources.CPUCores)).
		And(NewRAMBytesGreaterOrEqual(slot.Resources.RAMBytes))

	return s
}

func forAsk(slot ds.Slot) intf.CompositeSpecification {
	s := NewIsAskOrder().
		And(NewSupplierRatingGreaterOrEqual(slot.SupplierRating)).
		And((NewCPUCoresGreaterOrEqual(slot.Resources.CPUCores)).Not()).
		And((NewRAMBytesGreaterOrEqual(slot.Resources.RAMBytes)).Not())

	return s
}
