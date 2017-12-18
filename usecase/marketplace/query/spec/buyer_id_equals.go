package spec

import (
	ds "github.com/sonm-io/marketplace/datastruct"
	"github.com/sonm-io/marketplace/usecase/intf"
)

// BuyerIDEquals specifies whether the Buyer's ID is equal to the given value.
type BuyerIDEquals struct {
	ID string
}

// NewBuyerIDEquals creates a new instance of BuyerIDEquals.
func NewBuyerIDEquals(ID string) intf.CompositeSpecification {
	return intf.BaseSpecification{Specification: &BuyerIDEquals{ID: ID}}
}

// IsSatisfiedBy implements CompositeSpecification interface.
func (s *BuyerIDEquals) IsSatisfiedBy(object interface{}) bool {
	order := object.(*ds.Order)
	return order.BuyerID == s.ID
}
