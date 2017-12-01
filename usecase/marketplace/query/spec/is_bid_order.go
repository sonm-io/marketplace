package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsBidOrder specifies whether the given value is a bid order.
type IsBidOrder struct{}

// NewIsBidOrder creates a new instance of IsBidOrder.
func NewIsBidOrder() intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &IsBidOrder{}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *IsBidOrder) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.OrderType == ds.BID
}
