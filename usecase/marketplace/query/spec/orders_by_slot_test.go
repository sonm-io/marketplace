package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestOrdersBySlotIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		OrderType: ds.Bid,
		Slot: &ds.Slot{
			BuyerRating: 500,
			Resources: ds.Resources{
				CPUCores: 4,
				RAMBytes: 12000,
			},
		},
	}

	slot := &ds.Slot{
		BuyerRating: 300,
		Resources: ds.Resources{
			CPUCores: 2,
			RAMBytes: 12000,
		},
	}

	// act
	s, err := OrdersBySlot(ds.Bid, *slot)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.NoError(t, err)
	assert.True(t, obtained)
}

func TestOrdersBySlotIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			BuyerRating: 500,
		},
	}

	slot := ds.Slot{
		BuyerRating: 300,
		Resources: ds.Resources{
			CPUCores: 2,
			RAMBytes: 12000,
		},
	}

	// act
	s, err := OrdersBySlot(ds.Bid, slot)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.NoError(t, err)
	assert.False(t, obtained)
}
