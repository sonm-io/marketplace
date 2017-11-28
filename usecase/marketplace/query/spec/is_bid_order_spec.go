package spec

import (
	"github.com/sonm-io/marketplace/usecase/intf"
	"github.com/sonm-io/marketplace/usecase/marketplace/query/report"
)

// IsBidOrderSpec specifies whether the given value is a bid order.
type IsBidOrderSpec struct {
	intf.AbstractSpecification
}

// IsSatisfiedBy implements Specification interface.
func (s IsBidOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*report.GetOrderReport)
	return order.OrderType == report.BID
}
