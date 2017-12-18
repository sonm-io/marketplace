package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// NetworkOutboundLessOrEqual specifies whether the network outbound traffic is less than or equal to the given value.
type NetworkOutboundLessOrEqual struct {
	traffic uint64
}

// NewNetworkOutboundLessOrEqual creates a new instance of NetworkOutboundLessOrEqual.
func NewNetworkOutboundLessOrEqual(traffic uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &NetworkOutboundLessOrEqual{traffic: traffic}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *NetworkOutboundLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.NetTrafficOut <= s.traffic
}
