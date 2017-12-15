package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsAskOrderIsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		OrderType: ds.Ask,
	}

	// act
	s := NewIsAskOrder()
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestIsAskOrderIsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		OrderType: ds.Any,
	}

	// act
	s := NewIsAskOrder()
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
