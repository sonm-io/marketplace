package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsAskOrderSpec specifies whether the given value is an ask order.
type IsAskOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsAskOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.OrderType == ds.ASK
}
