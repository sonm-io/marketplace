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

// NewAskOrder creates a new ask order.
func NewAskOrder(ID, supplierID, buyerID, price string, slot ds.Slot) (*Order, error) {
	o := &Order{
		Order: ds.Order{
			ID:         ID,
			SupplierID: supplierID,
			BuyerID:    buyerID,
			Price:      price,
			OrderType:  ds.Ask,
			Slot:       &slot,
		},
	}

	if supplierID == "" {
		return nil, errSupplierIsRequired
	}

	if o.Price == "" {
		return nil, errPriceIsZero
	}

	return o, nil
}

// NewBidOrder creates a new bid order.
func NewBidOrder(ID, buyerID, supplierID, price string, slot ds.Slot) (*Order, error) {
	o := &Order{
		Order: ds.Order{
			ID:         ID,
			BuyerID:    buyerID,
			SupplierID: supplierID,
			Price:      price,
			OrderType:  ds.Bid,
			Slot:       &slot,
		},
	}

	if o.Price == "" {
		return nil, errPriceIsZero
	}

	if buyerID == "" {
		return nil, errBuyerIsRequired
	}

	return o, nil
}
