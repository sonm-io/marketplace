package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsAskOrder specifies whether the given value is an ask order.
type IsAskOrder struct {
	intf.CompositeSpecification
}

// NewIsAskOrder creates a new instance of IsAskOrder.
func NewIsAskOrder() intf.CompositeSpecification {
	s := &IsAskOrder{CompositeSpecification: &intf.BaseSpecification{}}
	s.Relate(s)

	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *IsAskOrder) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.OrderType == ds.ASK
}
