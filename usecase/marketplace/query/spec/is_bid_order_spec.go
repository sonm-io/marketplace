package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsBidOrderSpec specifies whether the given value is a bid order.
type IsBidOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsBidOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.OrderType == ds.BID
}
