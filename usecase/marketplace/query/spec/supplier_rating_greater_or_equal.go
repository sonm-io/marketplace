package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// SupplierRatingGreaterOrEqual specifies whether the supplier's rating is greater than or equal to the given value.
type SupplierRatingGreaterOrEqual struct {
	supplierRating int64
}

// NewSupplierRatingGreaterOrEqual creates a new instance of SupplierRatingGreaterOrEqual.
func NewSupplierRatingGreaterOrEqual(supplierRating int64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &SupplierRatingGreaterOrEqual{supplierRating: supplierRating}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *SupplierRatingGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.SupplierRating >= s.supplierRating
}
