package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// NetworkTypeLessOrEqual specifies whether the network type is less than or equal to the given value.
type NetworkTypeLessOrEqual struct {
	t ds.NetworkType
}

// NewNetworkTypeLessOrEqual creates a new instance of NetworkTypeLessOrEqual.
func NewNetworkTypeLessOrEqual(t ds.NetworkType) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &NetworkTypeLessOrEqual{t: t}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *NetworkTypeLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.NetworkType <= s.t
}
