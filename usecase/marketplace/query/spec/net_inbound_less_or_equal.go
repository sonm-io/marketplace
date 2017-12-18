package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// NetworkInboundLessOrEqual specifies whether the network inbound traffic less than or equal to the given value.
type NetworkInboundLessOrEqual struct {
	traffic uint64
}

// NewNetworkInboundLessOrEqual creates a new instance of NetworkInboundLessOrEqual.
func NewNetworkInboundLessOrEqual(traffic uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &NetworkInboundLessOrEqual{traffic: traffic}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *NetworkInboundLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.NetTrafficIn <= s.traffic
}
