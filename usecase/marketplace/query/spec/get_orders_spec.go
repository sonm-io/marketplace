package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// GetOrdersSpec is a factory method that creates a spec for GetOrder query.
func GetOrdersSpec(orderType ds.OrderType, slot ds.Slot) intf.Specification {
	var s intf.Specification
	switch orderType {
	case ds.ASK:
		s = NewSupplierRatingGreaterOrEqualSpec(slot.SupplierRating).
			And(IsAskOrderSpec{})
		s = IsAskOrderSpec{}
	case ds.BID:
		s = NewBuyerRatingGreaterOrEqualSpec(slot.BuyerRating).
			And(IsBidOrderSpec{})
		s = IsBidOrderSpec{}
	default:
		panic("unknown order type given")
	}

	/*
		s.compareCpuCoresBid(two) &&
		s.compareRamBytesBid(two) &&
		s.compareGpuCountBid(two) &&
		s.compareStorageBid(two) &&
		s.compareNetTrafficInBid(two) &&
		s.compareNetTrafficOutBid(two) &&
		s.compareNetworkTypeBid(two)
	*/

	return s
}
