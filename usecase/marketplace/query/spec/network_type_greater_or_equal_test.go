package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkTypeGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetworkType: ds.Inbound,
			},
		},
	}

	// act
	s := NewNetworkTypeGreaterOrEqual(ds.Inbound)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestNetworkTypeGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetworkType: ds.NoNetwork,
			},
		},
	}

	// act
	s := NewNetworkTypeGreaterOrEqual(ds.Outbound)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
