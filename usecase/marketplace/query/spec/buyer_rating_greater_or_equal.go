package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// BuyerRatingGreaterOrEqual specifies whether the buyer's rating is greater than or equal to the given value.
type BuyerRatingGreaterOrEqual struct {
	intf.CompositeSpecification
	buyerRating int64
}

// NewBuyerRatingGreaterOrEqual creates a new instance of BuyerRatingGreaterOrEqual.
func NewBuyerRatingGreaterOrEqual(buyerRating int64) intf.CompositeSpecification {
	s := &BuyerRatingGreaterOrEqual{CompositeSpecification: &intf.BaseSpecification{}, buyerRating: buyerRating}
	s.Relate(s)
	return s
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *BuyerRatingGreaterOrEqual) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.Slot.BuyerRating >= s.buyerRating
}
