package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNetworkInboundLessOrEqualIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficIn: 300,
			},
		},
	}

	// act
	s := NewNetworkInboundLessOrEqual(400)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestNetworkInboundLessOrEqualIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		Slot: &ds.Slot{
			Resources: ds.Resources{
				NetTrafficIn: 200,
			},
		},
	}

	// act
	s := NewNetworkInboundLessOrEqual(100)
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
