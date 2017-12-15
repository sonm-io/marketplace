package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkOutboundGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficOut: 400,
			},
		},
	}

	// act
	s := NewNetworkOutboundGreaterOrEqual(300)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestNetworkOutboundGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficOut: 100,
			},
		},
	}

	// act
	s := NewNetworkOutboundGreaterOrEqual(200)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
