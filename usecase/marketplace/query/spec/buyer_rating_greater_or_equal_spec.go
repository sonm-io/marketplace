package spec

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// BuyerRatingGreaterOrEqualSpec specifies whether the buyer's rating is greater than or equal to the given value.
type BuyerRatingGreaterOrEqualSpec struct {
	intf.AbstractSpecification

	buyerRating int64
}

// NewBuyerRatingGreaterOrEqualSpec creates a new instance of BuyerRatingGreaterOrEqualSpec.
func NewBuyerRatingGreaterOrEqualSpec(buyerRating int64) *BuyerRatingGreaterOrEqualSpec {
	return &BuyerRatingGreaterOrEqualSpec{buyerRating: buyerRating}
}

// IsSatisfiedBy implements Specification interface.
func (s *BuyerRatingGreaterOrEqualSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(*entity.Order)
	return order.Slot.BuyerRating >= s.buyerRating
}
