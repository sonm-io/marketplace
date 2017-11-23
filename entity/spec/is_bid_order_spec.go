package spec

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsBidOrderSpec specifies whether the given value is a bid order.
type IsBidOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsBidOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(entity.Order)
	return order.OrderType == entity.BID
}
