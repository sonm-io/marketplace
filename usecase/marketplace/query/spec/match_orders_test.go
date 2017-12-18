package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestMatchOrdersIsSatisfiedBy_SatisfyingBidOrderGiven_TrueReturned(t *testing.T) {
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

	matchAgainst := ds.Order{
		OrderType: ds.Bid,
		Slot: &ds.Slot{
			BuyerRating: 300,
			Resources: ds.Resources{
				CPUCores: 2,
				RAMBytes: 12000,
			},
		},
	}

	// act
	s, err := MatchOrders(matchAgainst)
	require.NoError(t, err)

	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestMatchOrdersIsSatisfiedBy_UnsatisfyingBidOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			BuyerRating: 500,
		},
	}

	matchAgainst := ds.Order{
		Slot: &ds.Slot{
			BuyerRating: 300,
			Resources: ds.Resources{
				CPUCores: 2,
				RAMBytes: 12000,
			},
		},
	}

	// act
	s, err := MatchOrders(matchAgainst)
	require.NoError(t, err)

	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}

func TestMatchOrdersIsSatisfiedBy_SatisfyingAskOrderGiven_TrueReturned(t *testing.T) {

	// arrange
	order := &ds.Order{
		OrderType: ds.Ask,
		Slot: &ds.Slot{
			SupplierRating: 500,
			Resources: ds.Resources{
				CPUCores: 1,
				RAMBytes: 100000000,
			},
		},
	}

	matchAgainst := ds.Order{
		OrderType: ds.Ask,
		Slot: &ds.Slot{
			SupplierRating: 500,
			Resources: ds.Resources{
				CPUCores:    1,
				RAMBytes:    100000000,
				Storage:     1000000000,
				NetworkType: ds.Inbound,
				Properties: map[string]float64{
					"cycles": 42,
					"foo":    1101,
				},
			},
		},
	}

	// act
	s, err := MatchOrders(matchAgainst)
	require.NoError(t, err)

	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestMatchOrdersIsSatisfiedBy_OrderWithNoSlotWithSupplierGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		SupplierID: "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD",
		Slot: &ds.Slot{
			BuyerRating: 500,
		},
	}

	matchAgainst := ds.Order{
		SupplierID: "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD",
	}

	// act
	s, err := MatchOrders(matchAgainst)
	require.NoError(t, err)

	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestMatchOrdersIsSatisfiedBy_OrderWithNoSlotWithBuyerGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		BuyerID: "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD",
		Slot: &ds.Slot{
			BuyerRating: 500,
		},
	}

	matchAgainst := ds.Order{
		BuyerID: "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD",
	}

	// act
	s, err := MatchOrders(matchAgainst)
	require.NoError(t, err)

	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestMatchOrdersIsSatisfiedBy_OrderWithNoSlotWithNoOwnerGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		BuyerID: "0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD",
		Slot: &ds.Slot{
			BuyerRating: 500,
		},
	}

	matchAgainst := ds.Order{}

	// act
	s, err := MatchOrders(matchAgainst)
	require.NoError(t, err)

	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
