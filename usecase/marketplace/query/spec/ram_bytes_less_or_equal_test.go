package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRamBytesLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				RAMBytes: 10000,
			},
		},
	}

	// act
	s := NewRAMBytesLessOrEqual(12000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestRamBytesLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				RAMBytes: 6000,
			},
		},
	}

	// act
	s := NewRAMBytesLessOrEqual(4000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
