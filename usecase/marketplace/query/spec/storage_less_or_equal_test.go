package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				Storage: 8000,
			},
		},
	}

	// act
	s := NewStorageLessOrEqual(10000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestStorageLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				Storage: 12000,
			},
		},
	}

	// act
	s := NewStorageLessOrEqual(10000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
