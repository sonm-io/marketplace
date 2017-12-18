package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// RAMBytesLessOrEqual specifies whether the memory size in bytes is less than or equal to the given value.
type RAMBytesLessOrEqual struct {
	ramBytes uint64
}

// NewRAMBytesLessOrEqual creates a new instance of RAMBytesLessOrEqual.
func NewRAMBytesLessOrEqual(ramBytes uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &RAMBytesLessOrEqual{ramBytes: ramBytes}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *RAMBytesLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.RAMBytes <= s.ramBytes
}
