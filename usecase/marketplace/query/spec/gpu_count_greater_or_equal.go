package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// GPUCountGreaterOrEqual specifies whether the gpu count is greater than or equal to the given value.
type GPUCountGreaterOrEqual struct {
	count ds.GPUCount
}

// NewGPUCountGreaterOrEqual creates a new instance of GPUCountGreaterOrEqual.
func NewGPUCountGreaterOrEqual(num ds.GPUCount) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &GPUCountGreaterOrEqual{count: num}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *GPUCountGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.GPUCount >= s.count
}
