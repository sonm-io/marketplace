package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// CpuCoresGreaterOrEqual specifies whether the cpu cores greater than or equal to the given value.
type CpuCoresGreaterOrEqual struct {
	intf.CompositeSpecification
	cpuCores uint64
}

// NewIsBidOrder NewCpuCoresGreaterOrEqual a new instance of CpuCoresGreaterOrEqual.
func NewCpuCoresGreaterOrEqual(num uint64) intf.CompositeSpecification {
	s := &CpuCoresGreaterOrEqual{CompositeSpecification: &intf.BaseSpecification{}, cpuCores: num}
	s.Relate(s)

	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *CpuCoresGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.CpuCores >= s.cpuCores
}
