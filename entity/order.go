package entity

import (
	"errors"
	ds "github.com/sonm-io/marketplace/datastruct"
)

// Order represents an order.
type Order struct {
	ds.Order
}

var (
	errPriceIsZero        = errors.New("order price cannot be negative")
	errSupplierIsRequired = errors.New("supplier is required")
	errBuyerIsRequired    = errors.New("buyer is required")
)

func NewAskOrder(ID, supplierID string, price int64, slot ds.Slot) (*Order, error) {
	o := &Order{
		Order: ds.Order{
			ID:         ID,
			SupplierID: supplierID,
			Price:      price,
			OrderType:  ds.ANY,
			Slot:       &slot,
		},
	}

	if supplierID == "" {
		return nil, errSupplierIsRequired
	}

	if o.Price <= 0 {
		return nil, errPriceIsZero
	}

	return o, nil
}

func NewBidOrder(ID, buyerID string, price int64, slot ds.Slot) (*Order, error) {
	o := &Order{
		Order: ds.Order{
			ID:        ID,
			BuyerID:   buyerID,
			Price:     price,
			OrderType: ds.BID,
			Slot:      &slot,
		},
	}

	if o.Price <= 0 {
		return nil, errPriceIsZero
	}

	if buyerID == "" {
		return nil, errBuyerIsRequired
	}

	return o, nil
}
