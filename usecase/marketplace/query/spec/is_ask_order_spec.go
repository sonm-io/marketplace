package spec

import (
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// IsAskOrderSpec specifies whether the given value is an ask order.
type IsAskOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsAskOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*report.GetOrderReport)
	return order.OrderType == report.ASK
}
