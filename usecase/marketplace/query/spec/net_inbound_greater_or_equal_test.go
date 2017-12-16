package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkInboundGreaterOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficIn: 400,
			},
		},
	}

	// act
	s := NewNetworkInboundGreaterOrEqual(300)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestNetworkInboundGreaterOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficIn: 100,
			},
		},
	}

	// act
	s := NewNetworkInboundGreaterOrEqual(200)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
