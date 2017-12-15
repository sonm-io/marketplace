package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// NetworkTypeGreaterOrEqual specifies whether the network type is greater than or equal to the given value.
type NetworkTypeGreaterOrEqual struct {
	t ds.NetworkType
}

// NewNetworkTypeGreaterOrEqual creates a new instance of NetworkTypeGreaterOrEqual.
func NewNetworkTypeGreaterOrEqual(t ds.NetworkType) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &NetworkTypeGreaterOrEqual{t: t}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *NetworkTypeGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.NetworkType >= s.t
}
