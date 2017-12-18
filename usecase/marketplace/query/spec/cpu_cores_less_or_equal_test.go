package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCpuCoresLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				CPUCores: 4,
			},
		},
	}

	// act
	s := NewCPUCoresLessOrEqual(5)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestCpuCoresLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				CPUCores: 3,
			},
		},
	}

	// act
	s := NewCPUCoresLessOrEqual(2)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
