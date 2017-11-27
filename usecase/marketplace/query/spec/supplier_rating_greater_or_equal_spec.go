package spec

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// SupplierRatingGreaterOrEqualSpec specifies whether the supplier's rating is greater than or equal to the given value.
type SupplierRatingGreaterOrEqualSpec struct {
	intf.AbstractSpecification

	supplierRating int64
}

// NewSupplierRatingGreaterOrEqualSpec creates a new instance of SupplierRatingGreaterOrEqualSpec.
func NewSupplierRatingGreaterOrEqualSpec(supplierRating int64) *SupplierRatingGreaterOrEqualSpec {
	return &SupplierRatingGreaterOrEqualSpec{supplierRating: supplierRating}
}

// IsSatisfiedBy implements Specification interface.
func (s *SupplierRatingGreaterOrEqualSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*entity.Order)
	return order.Slot.SupplierRating >= s.supplierRating
}
