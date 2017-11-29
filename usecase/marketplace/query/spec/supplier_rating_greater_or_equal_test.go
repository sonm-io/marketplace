package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplierRatingGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			SupplierRating: 1500,
		},
	}

	// act
	s := NewSupplierRatingGreaterOrEqual(1000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestSupplierRatingGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			SupplierRating: 500,
		},
	}

	// act
	s := NewSupplierRatingGreaterOrEqual(1000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
