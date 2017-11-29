package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// RAMBytesGreaterOrEqual specifies whether the memory size in bytes greater than or equal to the given value.
type RAMBytesGreaterOrEqual struct {
	intf.CompositeSpecification
	ramBytes uint64
}

// NewRAMBytesGreaterOrEqual creates a new instance of RAMBytesGreaterOrEqual.
func NewRAMBytesGreaterOrEqual(ramBytes uint64) intf.CompositeSpecification {
	s := &RAMBytesGreaterOrEqual{CompositeSpecification: &intf.BaseSpecification{}, ramBytes: ramBytes}
	s.Relate(s)
	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *RAMBytesGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.RamBytes >= s.ramBytes
}
