package spec

import (
	"github.com/sonm-io/marketplace/report"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// IsBidOrderSpec specifies whether the given value is a bid order.
type IsBidOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsBidOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(report.Order)
	return order.OrderType == report.BID
}
