package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsBidOrderIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		OrderType: ds.BID,
	}

	// act
	s := NewIsBidOrder()
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestIsBidOrderIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		OrderType: ds.ANY,
	}

	// act
	s := NewIsBidOrder()
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
