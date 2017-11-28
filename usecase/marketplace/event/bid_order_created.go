package event

type BidOrderCreated struct {
	Order Order
}

func NewBidOrderCreated(order Order) BidOrderCreated {
	return BidOrderCreated{Order: order}
}

func (e BidOrderCreated) EventID() string {
	return e.Order.ID
}
