package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkOutboundLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficOut: 200,
			},
		},
	}

	// act
	s := NewNetworkOutboundLessOrEqual(300)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestNetworkOutboundLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficOut: 200,
			},
		},
	}

	// act
	s := NewNetworkOutboundLessOrEqual(100)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
