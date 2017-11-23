package spec

import (
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsAskOrderSpec specifies whether the given value is an ask order.
type IsAskOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsAskOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(report.Order)
	return order.OrderType == report.ASK
}
