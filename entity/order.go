package entity

import "errors"

// Order represents an order.
type Order struct {
	// Order ID, UUIDv4
	ID string
	// Buyer's EtherumID (ASK)
	BuyerID string
	// Supplier's is EtherumID (BID)
	SupplierID string
	// Order price
	Price int64
	// Order type (Bid or Ask)
	OrderType OrderType
	// Slot describes resource requirements
	Slot *Slot
}

type OrderType int32

const (
	ANY OrderType = 0
	BID OrderType = 1
	ASK OrderType = 2
)

var (
	errPriceIsZero = errors.New("order price cannot be negative")
)

func NewAskOrder(ID, supplierID string, price int64, slot Slot) (*Order, error) {
	o := &Order{
		ID:         ID,
		SupplierID: supplierID,
		Price:      price,
		OrderType:  ASK,
		Slot:       &slot,
	}

	if o.Price <= 0 {
		return nil, errPriceIsZero
	}

	return o, nil
}

func NewBidOrder(ID, buyerID string, price int64, slot Slot) (*Order, error) {
	o := &Order{
		ID:        ID,
		BuyerID:   buyerID,
		Price:     price,
		OrderType: BID,
		Slot:      &slot,
	}

	if o.Price <= 0 {
		return nil, errPriceIsZero
	}

	return o, nil
}
