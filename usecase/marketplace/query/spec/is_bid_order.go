package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsBidOrder specifies whether the given value is a bid order.
type IsBidOrder struct {
	intf.CompositeSpecification
}

// NewIsBidOrder creates a new instance of IsBidOrder.
func NewIsBidOrder() intf.CompositeSpecification {
	s := &IsBidOrder{CompositeSpecification: &intf.BaseSpecification{}}
	s.Relate(s)

	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *IsBidOrder) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.OrderType == ds.BID
}
