package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// RamBytesGreaterOrEqual specifies whether the memory size in bytes greater than or equal to the given value.
type RamBytesGreaterOrEqual struct {
	intf.CompositeSpecification
	ramBytes uint64
}

// NewRamBytesGreaterOrEqual creates a new instance of RamBytesGreaterOrEqual.
func NewRamBytesGreaterOrEqual(ramBytes uint64) intf.CompositeSpecification {
	s := &RamBytesGreaterOrEqual{CompositeSpecification: &intf.BaseSpecification{}, ramBytes: ramBytes}
	s.Relate(s)
	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *RamBytesGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.RamBytes >= s.ramBytes
}
