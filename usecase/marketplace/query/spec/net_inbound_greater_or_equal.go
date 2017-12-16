package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// NetworkInboundGreaterOrEqual specifies whether the network inbound traffic greater than or equal to the given value.
type NetworkInboundGreaterOrEqual struct {
	traffic uint64
}

// NewNetworkInboundGreaterOrEqual creates a new instance of NetworkInboundGreaterOrEqual.
func NewNetworkInboundGreaterOrEqual(traffic uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &NetworkInboundGreaterOrEqual{traffic: traffic}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *NetworkInboundGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.NetTrafficIn >= s.traffic
}
