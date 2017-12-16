package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// StorageGreaterOrEqual specifies whether the storage size is greater than or equal to the given value.
type StorageGreaterOrEqual struct {
	size uint64
}

// NewStorageGreaterOrEqual creates a new instance of StorageGreaterOrEqual.
func NewStorageGreaterOrEqual(size uint64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &StorageGreaterOrEqual{size: size}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *StorageGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.Resources.Storage >= s.size
}
