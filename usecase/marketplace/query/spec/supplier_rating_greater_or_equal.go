package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// SupplierRatingGreaterOrEqual specifies whether the supplier's rating is greater than or equal to the given value.
type SupplierRatingGreaterOrEqual struct {
	intf.CompositeSpecification
	supplierRating int64
}

// NewSupplierRatingGreaterOrEqual creates a new instance of SupplierRatingGreaterOrEqual.
func NewSupplierRatingGreaterOrEqual(supplierRating int64) intf.CompositeSpecification {
	s := &SupplierRatingGreaterOrEqual{CompositeSpecification: &intf.BaseSpecification{}, supplierRating: supplierRating}
	s.Relate(s)
	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *SupplierRatingGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.SupplierRating >= s.supplierRating
}
