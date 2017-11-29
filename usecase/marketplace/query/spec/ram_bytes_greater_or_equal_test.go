package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRamBytesGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				RamBytes: 12000,
			},
		},
	}

	// act
	s := NewRAMBytesGreaterOrEqual(10000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestRamBytesGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				RamBytes: 6000,
			},
		},
	}

	// act
	s := NewRAMBytesGreaterOrEqual(10000)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
