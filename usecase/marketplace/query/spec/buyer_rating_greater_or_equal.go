package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// BuyerRatingGreaterOrEqual specifies whether the buyer's rating is greater than or equal to the given value.
type BuyerRatingGreaterOrEqual struct {
	buyerRating int64
}

// NewBuyerRatingGreaterOrEqual creates a new instance of BuyerRatingGreaterOrEqual.
func NewBuyerRatingGreaterOrEqual(buyerRating int64) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &BuyerRatingGreaterOrEqual{buyerRating: buyerRating}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *BuyerRatingGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.BuyerRating >= s.buyerRating
}
