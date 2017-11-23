package spec

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsAskOrderSpec specifies whether the given value is an ask order.
type IsAskOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsAskOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(entity.Order)
	return order.OrderType == entity.ASK
}
