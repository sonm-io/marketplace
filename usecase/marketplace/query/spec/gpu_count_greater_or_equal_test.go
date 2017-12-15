package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGPUCountGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				GPUCount: ds.MultipleGPU,
			},
		},
	}

	// act
	s := NewGPUCountGreaterOrEqual(ds.SingleGPU)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestGPUCountGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				GPUCount: ds.NoGPU,
			},
		},
	}

	// act
	s := NewGPUCountGreaterOrEqual(2)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
