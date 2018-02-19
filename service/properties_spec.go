package service

import (
	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/handler/proto"
)

// PropertiesSpec encapsulates the specification to match properties depending on OrderType.
type PropertiesSpec struct {
	order *pb.Order
}

// NewPropertiesSpec returns a new instance of PropertiesSpec.
func NewPropertiesSpec(order *pb.Order) PropertiesSpec {
	return PropertiesSpec{order: order}
}

// IsSatisfiedBy Checks if the given order's properties satisfy the matching criteria.
func (s PropertiesSpec) IsSatisfiedBy(order ds.Order) bool {
	if !s.queryHasProperties(s.order) {
		return true
	}

	if !s.orderHasProperties(order) {
		return false
	}

	var p1, p2 Properties
	p1 = order.Slot.Resources.Properties
	p2 = s.order.Slot.Resources.Properties

	if s.order.OrderType == pb.OrderType_BID && p1.LessOrEqual(p2) ||
		s.order.OrderType == pb.OrderType_ASK && p1.GreaterOrEqual(p2) {
		return true
	}

	return false
}

func (s PropertiesSpec) queryHasProperties(o *pb.Order) bool {
	return o.Slot != nil && len(o.Slot.Resources.Properties) > 0
}

func (s PropertiesSpec) orderHasProperties(order ds.Order) bool {
	return order.Order != nil && order.Order.Slot != nil && len(order.Order.Slot.Resources.Properties) > 0
}
