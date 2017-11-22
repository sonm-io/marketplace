package spec

import (
	"github.com/sonm-io/marketplace/entity"
	"github.com/sonm-io/marketplace/usecase/intf"
)

//
type IsAskOrderSpec struct {
	intf.AbstractSpecification
}

func (s IsAskOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(entity.Order)
	return order.OrderType == entity.ASK
}
//

//
type IsBidOrderSpec struct {
	intf.AbstractSpecification
}

func (s IsBidOrderSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(entity.Order)
	return order.OrderType == entity.BID
}
//

//
type SupplierRatingGreaterOrEqualSpec struct {
	intf.AbstractSpecification

	supplierRating int64
}

func NewSupplierRatingGreaterOrEqualSpec(supplierRating int64) *SupplierRatingGreaterOrEqualSpec{
	return &SupplierRatingGreaterOrEqualSpec{supplierRating:supplierRating}
}

func (s *SupplierRatingGreaterOrEqualSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(entity.Order)
	return order.Slot.SupplierRating >= s.supplierRating
}
//

//
type BuyerRatingGreaterOrEqualSpec struct {
	intf.AbstractSpecification

	buyerRating int64
}

func NewBuyerRatingGreaterOrEqualSpec(buyerRating int64) *BuyerRatingGreaterOrEqualSpec{
	return &BuyerRatingGreaterOrEqualSpec{buyerRating:buyerRating}
}

func (s *BuyerRatingGreaterOrEqualSpec) IsSatisfiedBy(object interface{}) bool {
	order := object.(entity.Order)
	return order.Slot.BuyerRating >= s.buyerRating
}
//