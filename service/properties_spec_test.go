package service

import (
	"github.com/stretchr/testify/assert"
	"testing"

	"github.com/sonm-io/marketplace/ds"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

func TestPropertiesSpecIsSatisfiedBy_RequestWithPropertiesOrderWithNoPropertiesGiven_FalseReturned(t *testing.T) {
	// assert
	req := makeGetOrdersReq(pb.OrderType_BID, map[string]float64{
		"foo": 777,
	})

	spec := NewPropertiesSpec(req)
	order := ds.Order{}

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_RequestMatchingBidOrderGiven_TrueReturned(t *testing.T) {
	// assert
	req := makeGetOrdersReq(pb.OrderType_BID, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(req)
	order := makeOrder(pb.OrderType_BID, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_RequestMatchingAskOrderGiven_TrueReturned(t *testing.T) {
	// assert
	q := makeGetOrdersReq(pb.OrderType_ASK, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(pb.OrderType_ASK, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_BidOrderLacksRequestedProperties_FalseReturned(t *testing.T) {
	// assert
	q := makeGetOrdersReq(pb.OrderType_BID, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(pb.OrderType_BID, map[string]float64{
		"foo": 777,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_AskOrderIsNotMatchingProperties_FalseReturned(t *testing.T) {
	// assert
	q := makeGetOrdersReq(pb.OrderType_ASK, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(pb.OrderType_ASK, map[string]float64{
		"foo": 776,
		"bar": 555,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func makeOrder(orderType pb.OrderType, properties map[string]float64) ds.Order {
	return ds.Order{
		Order: &pb.Order{
			OrderType: orderType,
			Slot: &pb.Slot{
				Resources: &pb.Resources{
					Properties: properties,
				},
			},
		},
	}
}

func makeGetOrdersReq(orderType pb.OrderType, properties map[string]float64) *pb.Order {
	return &pb.Order{
		OrderType: orderType,
		Slot: &pb.Slot{
			Resources: &pb.Resources{
				Properties: properties,
			},
		},
	}
}
