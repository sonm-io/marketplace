package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// StorageLessOrEqual specifies whether the storage size is less than or equal to the given value.
type StorageLessOrEqual struct {
	size uint64
}

// NewStorageLessOrEqual creates a new instance of StorageLessOrEqual.
func NewStorageLessOrEqual(size uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &StorageLessOrEqual{size: size}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *StorageLessOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.Storage <= s.size
}
