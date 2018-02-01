package sqllite

import (
	"testing"

	"github.com/stretchr/testify/assert"

	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/marketplace/query"
)

func TestPropertiesSpecIsSatisfiedBy_QueryWithPropertiesOrderWithNoPropertiesGiven_FalseReturned(t *testing.T) {
	// assert
	q := makeGetOrdersQuery(ds.Bid, map[string]float64{
		"foo": 777,
	})

	spec := NewPropertiesSpec(q)
	order := ds.Order{}

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_QueryMatchingBidOrderGiven_TrueReturned(t *testing.T) {
	// assert
	q := makeGetOrdersQuery(ds.Bid, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(ds.Bid, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_QueryMatchingAskOrderGiven_TrueReturned(t *testing.T) {
	// assert
	q := makeGetOrdersQuery(ds.Ask, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(ds.Ask, map[string]float64{
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
	q := makeGetOrdersQuery(ds.Bid, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(ds.Bid, map[string]float64{
		"foo": 777,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func TestPropertiesSpecIsSatisfiedBy_AskOrderIsNotMatchingProperties_FalseReturned(t *testing.T) {
	// assert
	q := makeGetOrdersQuery(ds.Ask, map[string]float64{
		"foo": 777,
		"bar": 555,
	})

	spec := NewPropertiesSpec(q)
	order := makeOrder(ds.Ask, map[string]float64{
		"foo": 776,
		"bar": 555,
	})

	// act
	obtained := spec.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func makeOrder(orderType ds.OrderType, properties map[string]float64) ds.Order {
	return ds.Order{
		OrderType: orderType,
		Slot: &ds.Slot{
			Resources: ds.Resources{
				Properties: properties,
			},
		},
	}
}

func makeGetOrdersQuery(orderType ds.OrderType, properties map[string]float64) query.GetOrders {
	return query.GetOrders{
		Order: ds.Order{
			OrderType: orderType,
			Slot: &ds.Slot{
				Resources: ds.Resources{
					Properties: properties,
				},
			},
		},
	}
}
