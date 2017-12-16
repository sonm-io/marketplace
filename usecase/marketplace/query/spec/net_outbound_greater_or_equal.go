package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// NetworkOutboundGreaterOrEqual specifies whether the network outbound traffic greater than or equal to the given value.
type NetworkOutboundGreaterOrEqual struct {
	traffic uint64
}

// NewNetworkOutboundGreaterOrEqual creates a new instance of NetworkOutboundGreaterOrEqual.
func NewNetworkOutboundGreaterOrEqual(traffic uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &NetworkOutboundGreaterOrEqual{traffic: traffic}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *NetworkOutboundGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.NetTrafficOut >= s.traffic
}
