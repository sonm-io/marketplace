package spec

import (
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// GetOrdersSpec is a factory method that creates a spec for GetOrder query.
func GetOrdersSpec(orderType report.OrderType, slot report.Slot) intf.Specification {
	var s intf.Specification
	switch orderType {
	case report.ASK:
		s = NewSupplierRatingGreaterOrEqualSpec(slot.SupplierRating).
			And(IsAskOrderSpec{})
	case report.BID:
		s = NewBuyerRatingGreaterOrEqualSpec(slot.BuyerRating).
			And(IsBidOrderSpec{})
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
