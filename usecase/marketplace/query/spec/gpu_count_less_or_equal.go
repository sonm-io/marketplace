package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// GPUCountLessOrEqual specifies whether the gpu count is less than or equal to the given value.
type GPUCountLessOrEqual struct {
	count ds.GPUCount
}

// NewGPUCountLessOrEqual creates a new instance of GPUCountLessOrEqual.
func NewGPUCountLessOrEqual(num ds.GPUCount) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &GPUCountLessOrEqual{count: num}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *GPUCountLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.GPUCount <= s.count
}
