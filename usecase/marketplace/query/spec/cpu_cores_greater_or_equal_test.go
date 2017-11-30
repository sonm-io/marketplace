package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCpuCoresGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				CPUCores: 4,
			},
		},
	}

	// act
	s := NewCPUCoresGreaterOrEqual(4)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestCpuCoresGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				CPUCores: 1,
			},
		},
	}

	// act
	s := NewCPUCoresGreaterOrEqual(2)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
