package spec

import (
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// GetOrdersSpec is a factory method that creates a spec for GetOrder query.
func GetOrdersSpec(orderType int, slot report.Slot) intf.Specification {
	var s intf.Specification
	switch orderType {
	case report.ASK:
		s = NewSupplierRatingGreaterOrEqualSpec(slot.SupplierRating).
			And(IsAskOrderSpec{})
		s = IsAskOrderSpec{}
	case report.BID:
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
