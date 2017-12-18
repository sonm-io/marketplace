package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSupplierIDEqualsSatisfiedBy_SatisfyingOrderGiven_TrueReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		SupplierID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
	}

	// act
	s := NewSupplierIDEquals("0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db")
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.True(t, obtained)
}

func TestSupplierIDEqualsSatisfiedBy_UnsatisfyingOrderGiven_FalseReturned(t *testing.T) {
	// arrange
	order := &ds.Order{
		SupplierID: "0x9A8568CD389580B6737FF56b61BE4F4eE802E2Db",
	}

	// act
	s := NewSupplierIDEquals("0x8125721C2413d99a33E351e1F6Bb4e56b6b633FD")
	obtained := s.IsSatisfiedBy(order)

	// assert
	assert.False(t, obtained)
}
