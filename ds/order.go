package ds

import (
	"errors"
	pb "github.com/sonm-io/marketplace/interface/grpc/proto"
)

// Order represents a safe order wrapper.
//
// This is used to decompose the validation out of the protocol. All
// methods must return the valid sub-structures.
type Order struct {
	*pb.Order
}

func (o *Order) Unwrap() *pb.Order {
	return o.Order
}

func NewOrder(o *pb.Order) (*Order, error) {
	if err := validateOrder(o); err != nil {
		return nil, err
	}

	return &Order{o}, nil
}

func validateOrder(o *pb.Order) error {
	if o == nil {
		return errors.New("order cannot be nil")
	}

	if o.PricePerSecond == nil {
		return errors.New("price cannot be nil")
	}

	if o.PricePerSecond.Unwrap().Sign() <= 0 {
		return errors.New("price/sec must be positive")
	}

	//if err := validateSlot(o.Slot); err != nil {
	//	return err
	//}

	return nil
}
