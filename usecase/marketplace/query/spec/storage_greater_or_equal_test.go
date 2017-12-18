package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStorageGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				Storage: 10000,
			},
		},
	}

	// act
	s := NewStorageGreaterOrEqual(8000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestStorageGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				Storage: 10000,
			},
		},
	}

	// act
	s := NewStorageGreaterOrEqual(12000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
