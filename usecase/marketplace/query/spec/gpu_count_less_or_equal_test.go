package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGPUCountLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				GPUCount: ds.SingleGPU,
			},
		},
	}

	// act
	s := NewGPUCountLessOrEqual(ds.SingleGPU)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestGPUCountLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				GPUCount: ds.SingleGPU,
			},
		},
	}

	// act
	s := NewGPUCountLessOrEqual(ds.NoGPU)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
