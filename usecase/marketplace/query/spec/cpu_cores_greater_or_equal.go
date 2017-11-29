package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CPUCoresGreaterOrEqual specifies whether the cpu cores greater than or equal to the given value.
type CPUCoresGreaterOrEqual struct {
	intf.CompositeSpecification
	cpuCores uint64
}

// NewCPUCoresGreaterOrEqual creates a new instance of CPUCoresGreaterOrEqual.
func NewCPUCoresGreaterOrEqual(num uint64) intf.CompositeSpecification {
	s := &CPUCoresGreaterOrEqual{CompositeSpecification: &intf.BaseSpecification{}, cpuCores: num}
	s.Relate(s)

	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *CPUCoresGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.CpuCores >= s.cpuCores
}
