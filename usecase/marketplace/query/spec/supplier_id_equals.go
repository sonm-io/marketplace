package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// SupplierIDEquals specifies whether the Supplier's ID is equal to the given value.
type SupplierIDEquals struct {
	ID string
}

// NewSupplierIDEquals creates a new instance of SupplierIDEquals.
func NewSupplierIDEquals(ID string) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &SupplierIDEquals{ID: ID}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *SupplierIDEquals) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.SupplierID == s.ID
}
