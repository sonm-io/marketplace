package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CPUCoresLessOrEqual specifies whether the cpu cores less than or equal to the given value.
type CPUCoresLessOrEqual struct {
	cpuCores uint64
}

// NewCPUCoresLessOrEqual creates a new instance of CPUCoresLessOrEqual.
func NewCPUCoresLessOrEqual(num uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &CPUCoresLessOrEqual{cpuCores: num}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *CPUCoresLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.CPUCores <= s.cpuCores
}
