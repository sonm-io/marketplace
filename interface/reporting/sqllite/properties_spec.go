package sqllite

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
)

// PropertiesSpec encapsulates the specification to match properties depending on OrderType.
type PropertiesSpec struct {
	q query.GetOrders
}

// NewPropertiesSpec returns a new instance of PropertiesSpec.
func NewPropertiesSpec(q query.GetOrders) PropertiesSpec {
	return PropertiesSpec{q: q}
}

// IsSatisfiedBy Checks if the given order's properties satisfy the matching criteria.
func (s PropertiesSpec) IsSatisfiedBy(order ds.Order) bool {
	if !s.queryHasProperties(s.q) {
		return true
	}

	if !s.orderHasProperties(order) {
		return false
	}

	var p1, p2 Properties
	p1 = order.Slot.Resources.Properties
	p2 = s.q.Order.Slot.Resources.Properties

	if s.q.Order.OrderType == ds.Bid && p1.LessOrEqual(p2) ||
		s.q.Order.OrderType == ds.Ask && p1.GreaterOrEqual(p2) {
		return true
	}

	return false
}

func (s PropertiesSpec) queryHasProperties(q query.GetOrders) bool {
	return q.Order.Slot != nil && len(q.Order.Slot.Resources.Properties) > 0
}

func (s PropertiesSpec) orderHasProperties(order ds.Order) bool {
	return order.Slot != nil && len(order.Slot.Resources.Properties) > 0
}
