package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuyerRatingGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			BuyerRating: 1500,
		},
	}

	// act
	s := NewBuyerRatingGreaterOrEqual(1000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestBuyerRatingGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			BuyerRating: 500,
		},
	}

	// act
	s := NewBuyerRatingGreaterOrEqual(1000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
