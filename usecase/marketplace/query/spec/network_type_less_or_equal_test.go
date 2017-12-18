package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkTypeLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetworkType: ds.NoNetwork,
			},
		},
	}

	// act
	s := NewNetworkTypeLessOrEqual(ds.Inbound)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestNetworkTypeLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetworkType: ds.Inbound,
			},
		},
	}

	// act
	s := NewNetworkTypeLessOrEqual(ds.NoNetwork)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
